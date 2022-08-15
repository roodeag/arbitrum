package eth

import (
	"github.com/roodeag/arbitrum/core"
	"github.com/roodeag/arbitrum/core/state"
	"github.com/roodeag/arbitrum/core/types"
	"github.com/roodeag/arbitrum/core/vm"
	"github.com/roodeag/arbitrum/ethdb"
)

func NewArbEthereum(
	blockchain *core.BlockChain,
	chainDb ethdb.Database,
) *Ethereum {
	return &Ethereum{
		blockchain: blockchain,
		chainDb:    chainDb,
	}
}

func (eth *Ethereum) StateAtTransaction(block *types.Block, txIndex int, reexec uint64) (core.Message, vm.BlockContext, *state.StateDB, error) {
	return eth.stateAtTransaction(block, txIndex, reexec)
}
