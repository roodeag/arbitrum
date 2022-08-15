package arbitrum

import (
	"context"

	"github.com/roodeag/arbitrum/core"
	"github.com/roodeag/arbitrum/core/types"
)

type ArbInterface interface {
	PublishTransaction(ctx context.Context, tx *types.Transaction) error
	BlockChain() *core.BlockChain
	ArbNode() interface{}
}
