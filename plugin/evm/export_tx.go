// (c) 2019-2020, Ava Labs, Inc. All rights reserved.
// See the file LICENSE for licensing terms.

package evm

import (
	"fmt"
	"math/big"

	"github.com/flare-foundation/coreth/core/state"
	"github.com/flare-foundation/coreth/params"

	"github.com/flare-foundation/flare/database"
	"github.com/flare-foundation/flare/ids"
	"github.com/flare-foundation/flare/snow"
	"github.com/flare-foundation/flare/utils/crypto"
	"github.com/flare-foundation/flare/vms/components/avax"
)

// UnsignedExportTx is an unsigned ExportTx
type UnsignedExportTx struct {
	avax.Metadata
	// ID of the network on which this tx was issued
	NetworkID uint32 `serialize:"true" json:"networkID"`
	// ID of this blockchain.
	BlockchainID ids.ID `serialize:"true" json:"blockchainID"`
	// Which chain to send the funds to
	DestinationChain ids.ID `serialize:"true" json:"destinationChain"`
	// Inputs
	Ins []EVMInput `serialize:"true" json:"inputs"`
	// Outputs that are exported to the chain
	ExportedOutputs []*avax.TransferableOutput `serialize:"true" json:"exportedOutputs"`
}

// InputUTXOs returns a set of all the hash(address:nonce) exporting funds.
func (tx *UnsignedExportTx) InputUTXOs() ids.Set {
	return ids.Set{}
}

// Verify this transaction is well-formed
func (tx *UnsignedExportTx) Verify(
	xChainID ids.ID,
	ctx *snow.Context,
	rules params.Rules,
) error {
	return errWrongChainID
}

func (tx *UnsignedExportTx) GasUsed() (uint64, error) {
	return 0, fmt.Errorf("exportTx transactions disabled")
}

// Amount of [assetID] burned by this transaction
func (tx *UnsignedExportTx) Burned(assetID ids.ID) (uint64, error) {
	return 0, fmt.Errorf("exportTx transactions disabled")
}

// SemanticVerify this transaction is valid.
func (tx *UnsignedExportTx) SemanticVerify(
	vm *VM,
	stx *Tx,
	_ *Block,
	baseFee *big.Int,
	rules params.Rules,
) error {
	return fmt.Errorf("exportTx transactions disabled")
}

// Accept this transaction.
func (tx *UnsignedExportTx) Accept(ctx *snow.Context, batch database.Batch) error {
	return fmt.Errorf("exportTx transactions disabled")
}

// newExportTx returns a new ExportTx
func (vm *VM) newExportTx(
	assetID ids.ID, // AssetID of the tokens to export
	amount uint64, // Amount of tokens to export
	chainID ids.ID, // Chain to send the UTXOs to
	to ids.ShortID, // Address of chain recipient
	baseFee *big.Int, // fee to use post-AP3
	keys []*crypto.PrivateKeySECP256K1R, // Pay the fee and provide the tokens
) (*Tx, error) {
	return nil, errWrongChainID
}

// EVMStateTransfer executes the state update from the atomic export transaction
func (tx *UnsignedExportTx) EVMStateTransfer(ctx *snow.Context, state *state.StateDB) error {
	return errInsufficientFunds
}
