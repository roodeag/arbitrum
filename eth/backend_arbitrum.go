package eth

import (
	"context"

	"github.com/roodeag/arbitrum/core"
	"github.com/roodeag/arbitrum/core/state"
	"github.com/roodeag/arbitrum/core/types"
	"github.com/roodeag/arbitrum/core/vm"
	"github.com/roodeag/arbitrum/eth/tracers"
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

func (eth *Ethereum) StateAtTransaction(ctx context.Context, block *types.Block, txIndex int, reexec uint64) (*core.Message, vm.BlockContext, *state.StateDB, tracers.StateReleaseFunc, error) {
	return eth.stateAtTransaction(ctx, block, txIndex, reexec)
}
