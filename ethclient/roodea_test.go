package ethclient

import (
	"context"
	"math/big"
	"testing"
)

func TestGetBlock(t *testing.T) {
	cli, err := Dial("https://goerli-rollup.arbitrum.io/rpc")
	if err != nil {
		t.Fatal(err)
	}
	blk, err := cli.BlockByNumber(context.Background(), big.NewInt(2014718))
	if err != nil {
		t.Fatal(err)
	}
	for i, transaction := range blk.Transactions() {
		txBuf, err := transaction.MarshalJSON()
		if err != nil {
			t.Fatal(err)
		}
		t.Logf("%d=%s\n", i, txBuf)
	}
}
