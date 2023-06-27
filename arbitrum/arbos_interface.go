package arbitrum

import (
	"context"

	"github.com/roodeag/arbitrum/arbitrum_types"
	"github.com/roodeag/arbitrum/core"
	"github.com/roodeag/arbitrum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction, options *arbitrum_types.ConditionalOptions) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
