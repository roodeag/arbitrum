package main

import (
	"context"
	"fmt"
	"github.com/roodeag/arbitrum/common"
	"github.com/roodeag/arbitrum/ethclient"
)

func main() {
	cli, err := ethclient.Dial("https://arb-mainnet.g.alchemy.com/v2/VTZXVSMPR2Wi_TR5xxAosfFkz4nbsGQr")
	if err != nil {
		panic(err)
	}
	tx, _, err := cli.TransactionByHash(context.Background(), common.HexToHash("0x394dff525e96d63e9958fa58d0243a9404b476ce4142e61775ab61b2c505784d"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("price=%s\n", tx.GasPrice().String())
	fmt.Printf("GasFeeCap=%s\n", tx.GasFeeCap().String())
	fmt.Printf("GasTipCap=%s\n", tx.GasTipCap().String())
	receipt, err := cli.TransactionReceipt(context.Background(), common.HexToHash("0x394dff525e96d63e9958fa58d0243a9404b476ce4142e61775ab61b2c505784d"))
	if err != nil {
		panic(err)
	}
	fmt.Printf("CumulativeGasUsed=%d\n", receipt.CumulativeGasUsed)
	fmt.Printf("GasUsedForL1=%d\n", receipt.GasUsedForL1)
	fmt.Printf("GasUsed=%d\n", receipt.GasUsed)
	fmt.Printf("EffectiveGasPrice=%s\n", receipt.EffectiveGasPrice.String())
}
