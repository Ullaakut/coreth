// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
	"strings"

	"github.com/ava-labs/coreth"

	"github.com/ava-labs/coreth/core/types"
	"github.com/ava-labs/gecko/api"
	"github.com/ava-labs/gecko/utils/constants"
	"github.com/ava-labs/go-ethereum/common"
	"github.com/ava-labs/go-ethereum/common/hexutil"
	"github.com/ava-labs/go-ethereum/crypto"
)

const (
	version = "Athereum 1.0"
)

// test constants
const (
	GenesisTestAddr = "0x751a0b96e1042bee789452ecb20253fba40dbe85"
	GenesisTestKey  = "0xabd71b35d559563fea757f0f5edbde286fb8c043105b15abb7cd57189306d7d1"
)

// DebugAPI introduces helper functions for debuging
type DebugAPI struct{ vm *VM }

// SnowmanAPI introduces snowman specific functionality to the evm
type SnowmanAPI struct{ vm *VM }

// NetAPI offers network related API methods
type NetAPI struct{ vm *VM }

type AvaAPI struct{ vm *VM }

// NewNetAPI creates a new net API instance.
func NewNetAPI(vm *VM) *NetAPI { return &NetAPI{vm} }

// Listening returns an indication if the node is listening for network connections.
func (s *NetAPI) Listening() bool { return true } // always listening

// PeerCount returns the number of connected peers
func (s *NetAPI) PeerCount() hexutil.Uint { return hexutil.Uint(0) } // TODO: report number of connected peers

// Version returns the current ethereum protocol version.
func (s *NetAPI) Version() string { return fmt.Sprintf("%d", s.vm.networkID) }

// Web3API offers helper API methods
type Web3API struct{}

// ClientVersion returns the version of the vm running
func (s *Web3API) ClientVersion() string { return version }

// Sha3 returns the bytes returned by hashing [input] with Keccak256
func (s *Web3API) Sha3(input hexutil.Bytes) hexutil.Bytes { return crypto.Keccak256(input) }

// GetAcceptedFrontReply defines the reply that will be sent from the
// GetAcceptedFront API call
type GetAcceptedFrontReply struct {
	Hash   common.Hash `json:"hash"`
	Number *big.Int    `json:"number"`
}

// GetAcceptedFront returns the last accepted block's hash and height
func (api *SnowmanAPI) GetAcceptedFront(ctx context.Context) (*GetAcceptedFrontReply, error) {
	blk := api.vm.getLastAccepted().ethBlock
	return &GetAcceptedFrontReply{
		Hash:   blk.Hash(),
		Number: blk.Number(),
	}, nil
}

// GetGenesisBalance returns the current funds in the genesis
func (api *DebugAPI) GetGenesisBalance(ctx context.Context) (*hexutil.Big, error) {
	lastAccepted := api.vm.getLastAccepted()
	api.vm.ctx.Log.Verbo("Currently accepted block front: %s", lastAccepted.ethBlock.Hash().Hex())
	state, err := api.vm.chain.BlockState(lastAccepted.ethBlock)
	if err != nil {
		return nil, err
	}
	return (*hexutil.Big)(state.GetBalance(common.HexToAddress(GenesisTestAddr))), nil
}

// SpendGenesis funds
func (api *DebugAPI) SpendGenesis(ctx context.Context, nonce uint64) error {
	api.vm.ctx.Log.Info("Spending the genesis")

	value := big.NewInt(1000000000000)
	gasLimit := 21000
	gasPrice := big.NewInt(1000000000)

	genPrivateKey, err := crypto.HexToECDSA(GenesisTestKey[2:])
	if err != nil {
		return err
	}
	bob, err := coreth.NewKey(rand.Reader)
	if err != nil {
		return err
	}

	tx := types.NewTransaction(nonce, bob.Address, value, uint64(gasLimit), gasPrice, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(api.vm.chainID), genPrivateKey)
	if err != nil {
		return err
	}

	if err := api.vm.issueRemoteTxs([]*types.Transaction{signedTx}); err != nil {
		return err
	}

	return nil
}

// IssueBlock to the chain
func (api *DebugAPI) IssueBlock(ctx context.Context) error {
	api.vm.ctx.Log.Info("Issuing a new block")

	return api.vm.tryBlockGen()
}

// ExportKeyArgs are arguments for ExportKey
type ExportKeyArgs struct {
	api.UserPass
	Address string `json:"address"`
}

// ExportKeyReply is the response for ExportKey
type ExportKeyReply struct {
	// The decrypted PrivateKey for the Address provided in the arguments
	PrivateKey string `json:"privateKey"`
}

// ExportKey returns a private key from the provided user
func (service *AvaAPI) ExportKey(r *http.Request, args *ExportKeyArgs, reply *ExportKeyReply) error {
	service.vm.ctx.Log.Info("Platform: ExportKey called")
	db, err := service.vm.ctx.Keystore.GetDatabase(args.Username, args.Password)
	if err != nil {
		return fmt.Errorf("problem retrieving user '%s': %w", args.Username, err)
	}
	user := user{db: db}
	if address, err := service.vm.ParseAddress(args.Address); err != nil {
		return fmt.Errorf("couldn't parse %s to address: %s", args.Address, err)
	} else if sk, err := user.getKey(address); err != nil {
		return fmt.Errorf("problem retrieving private key: %w", err)
	} else {
		reply.PrivateKey = common.ToHex(crypto.FromECDSA(sk))
		//constants.SecretKeyPrefix + formatting.CB58{Bytes: sk.Bytes()}.String()
		return nil
	}
}

// ImportKeyArgs are arguments for ImportKey
type ImportKeyArgs struct {
	api.UserPass
	PrivateKey string `json:"privateKey"`
}

// ImportKey adds a private key to the provided user
func (service *AvaAPI) ImportKey(r *http.Request, args *ImportKeyArgs, reply *api.JsonAddress) error {
	service.vm.ctx.Log.Info("Platform: ImportKey called for user '%s'", args.Username)
	db, err := service.vm.ctx.Keystore.GetDatabase(args.Username, args.Password)
	if err != nil {
		return fmt.Errorf("problem retrieving data: %w", err)
	}

	user := user{db: db}

	if !strings.HasPrefix(args.PrivateKey, constants.SecretKeyPrefix) {
		return fmt.Errorf("private key missing %s prefix", constants.SecretKeyPrefix)
	}
	sk, err := crypto.ToECDSA(common.FromHex(args.PrivateKey))
	if err != nil {
		return fmt.Errorf("invalid private key")
	}
	if err = user.putAddress(sk); err != nil {
		return fmt.Errorf("problem saving key %w", err)
	}

	reply.Address, err = service.vm.FormatAddress(crypto.PubkeyToAddress(sk.PublicKey))
	if err != nil {
		return fmt.Errorf("problem formatting address: %w", err)
	}
	return nil
}

// ImportAVAArgs are the arguments to ImportAVA
type ImportAVAArgs struct {
	api.UserPass
	// The address that will receive the imported funds
	To string `json:"to"`
}

// ImportAVA returns an unsigned transaction to import AVA from the X-Chain.
// The AVA must have already been exported from the X-Chain.
func (service *AvaAPI) ImportAVA(_ *http.Request, args *ImportAVAArgs, response *api.JsonTxID) error {
	service.vm.ctx.Log.Info("Platform: ImportAVA called")

	// Get the user's info
	db, err := service.vm.ctx.Keystore.GetDatabase(args.Username, args.Password)
	if err != nil {
		return fmt.Errorf("couldn't get user '%s': %w", args.Username, err)
	}
	user := user{db: db}

	to, err := service.vm.ParseAddress(args.To)
	if err != nil { // Parse address
		return fmt.Errorf("couldn't parse argument 'to' to an address: %w", err)
	}

	privKeys, err := user.getKeys()
	if err != nil { // Get keys
		return fmt.Errorf("couldn't get keys controlled by the user: %w", err)
	}

	tx, err := service.vm.newImportTx(to, privKeys)
	if err != nil {
		return err
	}

	response.TxID = tx.ID()
	return service.vm.issueTx(tx)
}
