// (c) 2019-2020, Ava Labs, Inc.
//
// This file is a derived work, based on the go-ethereum library whose original
// notices appear below.
//
// It is distributed under a license compatible with the licensing terms of the
// original code from which it is derived.
//
// Much love to the original authors for their work.
// **********
// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

// Avalanche ChainIDs
var (
	// FlareChainID ...
	FlareChainID = big.NewInt(14)
	// SongbirdChainID ...
	SongbirdChainID = big.NewInt(19)
	// CostonChainID ...
	CostonChainID = big.NewInt(16)
	// LocalChainID ...
	// Maximum possible chainID, see: https://github.com/ethereum/EIPs/issues/2294
	LocalChainID = big.NewInt(9_223_372_036_854_775_771)

	errNonGenesisForkByHeight = errors.New("coreth only supports forking by height at the genesis block")
)

var (
	// FlareChainConfig is the configuration for Avalanche Main Network
	FlareChainConfig = &ChainConfig{
		ChainID:                     FlareChainID,
		HomesteadBlock:              big.NewInt(0),
		DAOForkBlock:                big.NewInt(0),
		DAOForkSupport:              true,
		EIP150Block:                 big.NewInt(0),
		EIP150Hash:                  common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0"),
		EIP155Block:                 big.NewInt(0),
		EIP158Block:                 big.NewInt(0),
		ByzantiumBlock:              big.NewInt(0),
		ConstantinopleBlock:         big.NewInt(0),
		PetersburgBlock:             big.NewInt(0),
		IstanbulBlock:               big.NewInt(0),
		MuirGlacierBlock:            big.NewInt(0),
		ApricotPhase1BlockTimestamp: big.NewInt(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
		ApricotPhase2BlockTimestamp: big.NewInt(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
		ApricotPhase3BlockTimestamp: big.NewInt(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
		ApricotPhase4BlockTimestamp: big.NewInt(time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
	}

	// SongbirdChainConfig is the configuration for the Fuji Test Network
	SongbirdChainConfig = &ChainConfig{
		ChainID:                     SongbirdChainID,
		HomesteadBlock:              big.NewInt(0),
		DAOForkBlock:                big.NewInt(0),
		DAOForkSupport:              true,
		EIP150Block:                 big.NewInt(0),
		EIP150Hash:                  common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0"),
		EIP155Block:                 big.NewInt(0),
		EIP158Block:                 big.NewInt(0),
		ByzantiumBlock:              big.NewInt(0),
		ConstantinopleBlock:         big.NewInt(0),
		PetersburgBlock:             big.NewInt(0),
		IstanbulBlock:               big.NewInt(0),
		MuirGlacierBlock:            big.NewInt(0),
		ApricotPhase1BlockTimestamp: big.NewInt(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
		ApricotPhase2BlockTimestamp: big.NewInt(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
		ApricotPhase3BlockTimestamp: big.NewInt(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
		ApricotPhase4BlockTimestamp: big.NewInt(time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
	}

	// CostonChainConfig is the configuration for the Fuji Test Network
	CostonChainConfig = &ChainConfig{
		ChainID:                     CostonChainID,
		HomesteadBlock:              big.NewInt(0),
		DAOForkBlock:                big.NewInt(0),
		DAOForkSupport:              true,
		EIP150Block:                 big.NewInt(0),
		EIP150Hash:                  common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0"),
		EIP155Block:                 big.NewInt(0),
		EIP158Block:                 big.NewInt(0),
		ByzantiumBlock:              big.NewInt(0),
		ConstantinopleBlock:         big.NewInt(0),
		PetersburgBlock:             big.NewInt(0),
		IstanbulBlock:               big.NewInt(0),
		MuirGlacierBlock:            big.NewInt(0),
		ApricotPhase1BlockTimestamp: big.NewInt(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
		ApricotPhase2BlockTimestamp: big.NewInt(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
		ApricotPhase3BlockTimestamp: big.NewInt(time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
		ApricotPhase4BlockTimestamp: big.NewInt(time.Date(2100, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()),
	}

	// FlareLocalChainConfig is the configuration for the Avalanche Local Network
	FlareLocalChainConfig = &ChainConfig{
		ChainID:                     LocalChainID,
		HomesteadBlock:              big.NewInt(0),
		DAOForkBlock:                big.NewInt(0),
		DAOForkSupport:              true,
		EIP150Block:                 big.NewInt(0),
		EIP150Hash:                  common.HexToHash("0x2086799aeebeae135c246c65021c82b4e15a2c451340993aacfd2751886514f0"),
		EIP155Block:                 big.NewInt(0),
		EIP158Block:                 big.NewInt(0),
		ByzantiumBlock:              big.NewInt(0),
		ConstantinopleBlock:         big.NewInt(0),
		PetersburgBlock:             big.NewInt(0),
		IstanbulBlock:               big.NewInt(0),
		MuirGlacierBlock:            big.NewInt(0),
		ApricotPhase1BlockTimestamp: big.NewInt(0),
		ApricotPhase2BlockTimestamp: big.NewInt(0),
		ApricotPhase3BlockTimestamp: big.NewInt(0),
		ApricotPhase4BlockTimestamp: big.NewInt(0),
	}

	TestChainConfig         = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	TestLaunchConfig        = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil, nil}
	TestApricotPhase1Config = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil, nil}
	TestApricotPhase2Config = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil, nil}
	TestApricotPhase3Config = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), nil}
	TestApricotPhase4Config = &ChainConfig{big.NewInt(1), big.NewInt(0), nil, false, big.NewInt(0), common.Hash{}, big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0), big.NewInt(0)}
	TestRules               = TestChainConfig.AvalancheRules(new(big.Int), new(big.Int))
)

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock      *big.Int `json:"byzantiumBlock,omitempty"`      // Byzantium switch block (nil = no fork, 0 = already on byzantium)
	ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	PetersburgBlock     *big.Int `json:"petersburgBlock,omitempty"`     // Petersburg switch block (nil = same as Constantinople)
	IstanbulBlock       *big.Int `json:"istanbulBlock,omitempty"`       // Istanbul switch block (nil = no fork, 0 = already on istanbul)
	MuirGlacierBlock    *big.Int `json:"muirGlacierBlock,omitempty"`    // Eip-2384 (bomb delay) switch block (nil = no fork, 0 = already activated)

	// Avalanche Network Upgrades
	ApricotPhase1BlockTimestamp *big.Int `json:"apricotPhase1BlockTimestamp,omitempty"` // Apricot Phase 1 Block Timestamp (nil = no fork, 0 = already activated)
	// Apricot Phase 2 Block Timestamp (nil = no fork, 0 = already activated)
	// Apricot Phase 2 includes a modified version of the Berlin Hard Fork from Ethereum
	ApricotPhase2BlockTimestamp *big.Int `json:"apricotPhase2BlockTimestamp,omitempty"`
	// Apricot Phase 3 introduces dynamic fees and a modified version of the London Hard Fork from Ethereum (nil = no fork, 0 = already activated)
	ApricotPhase3BlockTimestamp *big.Int `json:"apricotPhase3BlockTimestamp,omitempty"`
	// Apricot Phase 4 introduces the notion of a block fee to the dynamic fee algorithm (nil = no fork, 0 = already activated)
	ApricotPhase4BlockTimestamp *big.Int `json:"apricotPhase4BlockTimestamp,omitempty"`
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Constantinople: %v Petersburg: %v Istanbul: %v, Muir Glacier: %v, Apricot Phase 1: %v, Apricot Phase 2: %v, Apricot Phase 3: %v, Apricot Phase 4: %v, Engine: Dummy Consensus Engine}",
		c.ChainID,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		c.ConstantinopleBlock,
		c.PetersburgBlock,
		c.IstanbulBlock,
		c.MuirGlacierBlock,
		c.ApricotPhase1BlockTimestamp,
		c.ApricotPhase2BlockTimestamp,
		c.ApricotPhase3BlockTimestamp,
		c.ApricotPhase4BlockTimestamp,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAOFork returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

// IsEIP150 returns whether num is either equal to the EIP150 fork block or greater.
func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

// IsEIP155 returns whether num is either equal to the EIP155 fork block or greater.
func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

// IsEIP158 returns whether num is either equal to the EIP158 fork block or greater.
func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

// IsByzantium returns whether num is either equal to the Byzantium fork block or greater.
func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	return isForked(c.ByzantiumBlock, num)
}

// IsConstantinople returns whether num is either equal to the Constantinople fork block or greater.
func (c *ChainConfig) IsConstantinople(num *big.Int) bool {
	return isForked(c.ConstantinopleBlock, num)
}

// IsMuirGlacier returns whether num is either equal to the Muir Glacier (EIP-2384) fork block or greater.
func (c *ChainConfig) IsMuirGlacier(num *big.Int) bool {
	return isForked(c.MuirGlacierBlock, num)
}

// IsPetersburg returns whether num is either
// - equal to or greater than the PetersburgBlock fork block,
// - OR is nil, and Constantinople is active
func (c *ChainConfig) IsPetersburg(num *big.Int) bool {
	return isForked(c.PetersburgBlock, num) || c.PetersburgBlock == nil && isForked(c.ConstantinopleBlock, num)
}

// IsIstanbul returns whether num is either equal to the Istanbul fork block or greater.
func (c *ChainConfig) IsIstanbul(num *big.Int) bool {
	return isForked(c.IstanbulBlock, num)
}

// Avalanche Upgrades:

// IsApricotPhase1 returns whether [blockTimestamp] represents a block
// with a timestamp after the Apricot Phase 1 upgrade time.
func (c *ChainConfig) IsApricotPhase1(blockTimestamp *big.Int) bool {
	return isForked(c.ApricotPhase1BlockTimestamp, blockTimestamp)
}

// IsApricotPhase2 returns whether [blockTimestamp] represents a block
// with a timestamp after the Apricot Phase 2 upgrade time.
func (c *ChainConfig) IsApricotPhase2(blockTimestamp *big.Int) bool {
	return isForked(c.ApricotPhase2BlockTimestamp, blockTimestamp)
}

// IsApricotPhase3 returns whether [blockTimestamp] represents a block
// with a timestamp after the Apricot Phase 3 upgrade time.
func (c *ChainConfig) IsApricotPhase3(blockTimestamp *big.Int) bool {
	return isForked(c.ApricotPhase3BlockTimestamp, blockTimestamp)
}

// IsApricotPhase4 returns whether [blockTimestamp] represents a block
// with a timestamp after the Apricot Phase 4 upgrade time.
func (c *ChainConfig) IsApricotPhase4(blockTimestamp *big.Int) bool {
	return isForked(c.ApricotPhase4BlockTimestamp, blockTimestamp)
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

// CheckConfigForkOrder checks that we don't "skip" any forks, geth isn't pluggable enough
// to guarantee that forks can be implemented in a different order than on official networks
func (c *ChainConfig) CheckConfigForkOrder() error {
	type fork struct {
		name     string
		block    *big.Int
		optional bool // if true, the fork may be nil and next fork is still allowed
	}
	var lastFork fork
	for _, cur := range []fork{
		{name: "homesteadBlock", block: c.HomesteadBlock},
		{name: "daoForkBlock", block: c.DAOForkBlock, optional: true},
		{name: "eip150Block", block: c.EIP150Block},
		{name: "eip155Block", block: c.EIP155Block},
		{name: "eip158Block", block: c.EIP158Block},
		{name: "byzantiumBlock", block: c.ByzantiumBlock},
		{name: "constantinopleBlock", block: c.ConstantinopleBlock},
		{name: "petersburgBlock", block: c.PetersburgBlock},
		{name: "istanbulBlock", block: c.IstanbulBlock},
		{name: "muirGlacierBlock", block: c.MuirGlacierBlock, optional: true},
	} {
		if cur.block != nil && common.Big0.Cmp(cur.block) != 0 {
			return errNonGenesisForkByHeight
		}
		if lastFork.name != "" {
			// Next one must be higher number
			if lastFork.block == nil && cur.block != nil {
				return fmt.Errorf("unsupported fork ordering: %v not enabled, but %v enabled at %v",
					lastFork.name, cur.name, cur.block)
			}
			if lastFork.block != nil && cur.block != nil {
				if lastFork.block.Cmp(cur.block) > 0 {
					return fmt.Errorf("unsupported fork ordering: %v enabled at %v, but %v enabled at %v",
						lastFork.name, lastFork.block, cur.name, cur.block)
				}
			}
		}
		// If it was optional and not set, then ignore it
		if !cur.optional || cur.block != nil {
			lastFork = cur
		}
	}

	// Note: ApricotPhase1 and ApricotPhase2 override the rules set by block number
	// hard forks. In Avalanche, hard forks must take place via block timestamps instead
	// of block numbers since blocks are produced asynchronously. Therefore, we do not
	// check that the block timestamps for Apricot Phase1 and Phase2 in the same way as for
	// the block number forks since it would not be a meaningful comparison.
	// Instead, we check only that Apricot Phases are enabled in order.
	lastFork = fork{}
	for _, cur := range []fork{
		{name: "apricotPhase1BlockTimestamp", block: c.ApricotPhase1BlockTimestamp},
		{name: "apricotPhase2BlockTimestamp", block: c.ApricotPhase2BlockTimestamp},
		{name: "apricotPhase3BlockTimestamp", block: c.ApricotPhase3BlockTimestamp},
		{name: "apricotPhase4BlockTimestamp", block: c.ApricotPhase4BlockTimestamp},
	} {
		if lastFork.name != "" {
			// Next one must be higher number
			if lastFork.block == nil && cur.block != nil {
				return fmt.Errorf("unsupported fork ordering: %v not enabled, but %v enabled at %v",
					lastFork.name, cur.name, cur.block)
			}
			if lastFork.block != nil && cur.block != nil {
				if lastFork.block.Cmp(cur.block) > 0 {
					return fmt.Errorf("unsupported fork ordering: %v enabled at %v, but %v enabled at %v",
						lastFork.name, lastFork.block, cur.name, cur.block)
				}
			}
		}
		// If it was optional and not set, then ignore it
		if !cur.optional || cur.block != nil {
			lastFork = cur
		}
	}
	// TODO(aaronbuchwald) check that avalanche block timestamps are at least possible with the other rule set changes
	// additional change: require that block number hard forks are either 0 or nil since they should not
	// be enabled at a specific block number.

	return nil
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainID, newcfg.ChainID) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	if isForkIncompatible(c.ConstantinopleBlock, newcfg.ConstantinopleBlock, head) {
		return newCompatError("Constantinople fork block", c.ConstantinopleBlock, newcfg.ConstantinopleBlock)
	}
	if isForkIncompatible(c.PetersburgBlock, newcfg.PetersburgBlock, head) {
		// the only case where we allow Petersburg to be set in the past is if it is equal to Constantinople
		// mainly to satisfy fork ordering requirements which state that Petersburg fork be set if Constantinople fork is set
		if isForkIncompatible(c.ConstantinopleBlock, newcfg.PetersburgBlock, head) {
			return newCompatError("Petersburg fork block", c.PetersburgBlock, newcfg.PetersburgBlock)
		}
	}
	if isForkIncompatible(c.IstanbulBlock, newcfg.IstanbulBlock, head) {
		return newCompatError("Istanbul fork block", c.IstanbulBlock, newcfg.IstanbulBlock)
	}
	if isForkIncompatible(c.MuirGlacierBlock, newcfg.MuirGlacierBlock, head) {
		return newCompatError("Muir Glacier fork block", c.MuirGlacierBlock, newcfg.MuirGlacierBlock)
	}
	// TODO(aaronbuchwald) ensure that Avalanche Blocktimestamps are not modified
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntactic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID                                                 *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158               bool
	IsByzantium, IsConstantinople, IsPetersburg, IsIstanbul bool

	// Rules for Avalanche releases
	IsApricotPhase1 bool
	IsApricotPhase2 bool
	IsApricotPhase3 bool
	IsApricotPhase4 bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID:          new(big.Int).Set(chainID),
		IsHomestead:      c.IsHomestead(num),
		IsEIP150:         c.IsEIP150(num),
		IsEIP155:         c.IsEIP155(num),
		IsEIP158:         c.IsEIP158(num),
		IsByzantium:      c.IsByzantium(num),
		IsConstantinople: c.IsConstantinople(num),
		IsPetersburg:     c.IsPetersburg(num),
		IsIstanbul:       c.IsIstanbul(num),
	}
}

// AvalancheRules returns the Avalanche modified rules to support Avalanche
// network upgrades
func (c *ChainConfig) AvalancheRules(blockNum, blockTimestamp *big.Int) Rules {
	rules := c.rules(blockNum)

	rules.IsApricotPhase1 = c.IsApricotPhase1(blockTimestamp)
	rules.IsApricotPhase2 = c.IsApricotPhase2(blockTimestamp)
	rules.IsApricotPhase3 = c.IsApricotPhase3(blockTimestamp)
	rules.IsApricotPhase4 = c.IsApricotPhase4(blockTimestamp)
	return rules
}
