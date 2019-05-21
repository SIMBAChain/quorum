// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package eth

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus/ethash"
	"github.com/ethereum/go-ethereum/consensus/istanbul"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/eth/downloader"
	"github.com/ethereum/go-ethereum/eth/gasprice"
)

var _ = (*configMarshaling)(nil)

// MarshalTOML marshals as TOML.
func (c Config) MarshalTOML() (interface{}, error) {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               uint64
		SyncMode                downloader.SyncMode
		NoPruning               bool
		LightServ               int  `toml:",omitempty"`
		LightPeers              int  `toml:",omitempty"`
		SkipBcVersionCheck      bool `toml:"-"`
		DatabaseHandles         int  `toml:"-"`
		DatabaseCache           int
		TrieCleanCache          int
		TrieDirtyCache          int
		TrieTimeout             time.Duration
		Etherbase               common.Address `toml:",omitempty"`
		MinerNotify             []string       `toml:",omitempty"`
		MinerExtraData          hexutil.Bytes  `toml:",omitempty"`
		MinerGasFloor           uint64
		MinerGasCeil            uint64
		MinerGasPrice           *big.Int
		MinerRecommit           time.Duration
		MinerNoverify           bool
		Ethash                  ethash.Config
		TxPool                  core.TxPoolConfig
		GPO                     gasprice.Config
		EnablePreimageRecording bool
		Istanbul                istanbul.Config
		DocRoot                 string `toml:"-"`
		EWASMInterpreter        string
		EVMInterpreter          string
	}
	var enc Config
	enc.Genesis = c.Genesis
	enc.NetworkId = c.NetworkId
	enc.SyncMode = c.SyncMode
	enc.NoPruning = c.NoPruning
	enc.LightServ = c.LightServ
	enc.LightPeers = c.LightPeers
	enc.SkipBcVersionCheck = c.SkipBcVersionCheck
	enc.DatabaseHandles = c.DatabaseHandles
	enc.DatabaseCache = c.DatabaseCache
	enc.TrieCleanCache = c.TrieCleanCache
	enc.TrieDirtyCache = c.TrieDirtyCache
	enc.TrieTimeout = c.TrieTimeout
	enc.Etherbase = c.Etherbase
	enc.MinerNotify = c.MinerNotify
	enc.MinerExtraData = c.MinerExtraData
	enc.MinerGasFloor = c.MinerGasFloor
	enc.MinerGasCeil = c.MinerGasCeil
	enc.MinerGasPrice = c.MinerGasPrice
	enc.MinerRecommit = c.MinerRecommit
	enc.MinerNoverify = c.MinerNoverify
	enc.Ethash = c.Ethash
	enc.TxPool = c.TxPool
	enc.GPO = c.GPO
	enc.EnablePreimageRecording = c.EnablePreimageRecording
	enc.Istanbul = c.Istanbul
	enc.DocRoot = c.DocRoot
	enc.EWASMInterpreter = c.EWASMInterpreter
	enc.EVMInterpreter = c.EVMInterpreter
	return &enc, nil
}

// UnmarshalTOML unmarshals from TOML.
func (c *Config) UnmarshalTOML(unmarshal func(interface{}) error) error {
	type Config struct {
		Genesis                 *core.Genesis `toml:",omitempty"`
		NetworkId               *uint64
		SyncMode                *downloader.SyncMode
		NoPruning               *bool
		LightServ               *int  `toml:",omitempty"`
		LightPeers              *int  `toml:",omitempty"`
		SkipBcVersionCheck      *bool `toml:"-"`
		DatabaseHandles         *int  `toml:"-"`
		DatabaseCache           *int
		TrieCleanCache          *int
		TrieDirtyCache          *int
		TrieTimeout             *time.Duration
		Etherbase               *common.Address `toml:",omitempty"`
		MinerNotify             []string        `toml:",omitempty"`
		MinerExtraData          *hexutil.Bytes  `toml:",omitempty"`
		MinerGasFloor           *uint64
		MinerGasCeil            *uint64
		MinerGasPrice           *big.Int
		MinerRecommit           *time.Duration
		MinerNoverify           *bool
		Ethash                  *ethash.Config
		TxPool                  *core.TxPoolConfig
		GPO                     *gasprice.Config
		EnablePreimageRecording *bool
		Istanbul                *istanbul.Config
		DocRoot                 *string `toml:"-"`
		EWASMInterpreter        *string
		EVMInterpreter          *string
	}
	var dec Config
	if err := unmarshal(&dec); err != nil {
		return err
	}
	if dec.Genesis != nil {
		c.Genesis = dec.Genesis
	}
	if dec.NetworkId != nil {
		c.NetworkId = *dec.NetworkId
	}
	if dec.SyncMode != nil {
		c.SyncMode = *dec.SyncMode
	}
	if dec.NoPruning != nil {
		c.NoPruning = *dec.NoPruning
	}
	if dec.LightServ != nil {
		c.LightServ = *dec.LightServ
	}
	if dec.LightPeers != nil {
		c.LightPeers = *dec.LightPeers
	}
	if dec.SkipBcVersionCheck != nil {
		c.SkipBcVersionCheck = *dec.SkipBcVersionCheck
	}
	if dec.DatabaseHandles != nil {
		c.DatabaseHandles = *dec.DatabaseHandles
	}
	if dec.DatabaseCache != nil {
		c.DatabaseCache = *dec.DatabaseCache
	}
	if dec.TrieCleanCache != nil {
		c.TrieCleanCache = *dec.TrieCleanCache
	}
	if dec.TrieDirtyCache != nil {
		c.TrieDirtyCache = *dec.TrieDirtyCache
	}
	if dec.TrieTimeout != nil {
		c.TrieTimeout = *dec.TrieTimeout
	}
	if dec.Etherbase != nil {
		c.Etherbase = *dec.Etherbase
	}
	if dec.MinerNotify != nil {
		c.MinerNotify = dec.MinerNotify
	}
	if dec.MinerExtraData != nil {
		c.MinerExtraData = *dec.MinerExtraData
	}
	if dec.MinerGasFloor != nil {
		c.MinerGasFloor = *dec.MinerGasFloor
	}
	if dec.MinerGasCeil != nil {
		c.MinerGasCeil = *dec.MinerGasCeil
	}
	if dec.MinerGasPrice != nil {
		c.MinerGasPrice = dec.MinerGasPrice
	}
	if dec.MinerRecommit != nil {
		c.MinerRecommit = *dec.MinerRecommit
	}
	if dec.MinerNoverify != nil {
		c.MinerNoverify = *dec.MinerNoverify
	}
	if dec.Ethash != nil {
		c.Ethash = *dec.Ethash
	}
	if dec.TxPool != nil {
		c.TxPool = *dec.TxPool
	}
	if dec.GPO != nil {
		c.GPO = *dec.GPO
	}
	if dec.EnablePreimageRecording != nil {
		c.EnablePreimageRecording = *dec.EnablePreimageRecording
	}
	if dec.Istanbul != nil {
		c.Istanbul = *dec.Istanbul
	}
	if dec.DocRoot != nil {
		c.DocRoot = *dec.DocRoot
	}
	if dec.EWASMInterpreter != nil {
		c.EWASMInterpreter = *dec.EWASMInterpreter
	}
	if dec.EVMInterpreter != nil {
		c.EVMInterpreter = *dec.EVMInterpreter
	}
	return nil
}
