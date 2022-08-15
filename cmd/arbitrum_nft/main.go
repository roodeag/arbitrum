package main

import (
	"bytes"
	"context"
	"fmt"
	ethereum "github.com/roodeag/arbitrum"
	"github.com/roodeag/arbitrum/accounts/abi"
	erc20 "github.com/roodeag/arbitrum/cmd/arbitrum_nft/abi"
	"github.com/roodeag/arbitrum/common"
	"github.com/roodeag/arbitrum/core/types"
	"github.com/roodeag/arbitrum/ethclient"
	"math/big"
	"strings"
	"time"
)

var TransferEventHash = common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

const TransferEventName = "Transfer"

func main() {
	cli, err := ethclient.Dial("https://rinkeby.arbitrum.io/rpc")
	if err != nil {
		panic(err)
	}
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	height, err := cli.BlockNumber(ctxWithTimeout)
	if err != nil {
		panic(err)
	}

	fmt.Printf("height=%d\n", height)

	blockHash := common.HexToHash("0x197ddee1d74eaf7edbd4d46e72aa362c894c01fcf33e976fc4ba7eadf0c01085")
	topics := [][]common.Hash{{TransferEventHash}}
	logs, err := cli.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []common.Address{common.HexToAddress("0xe12bdce43b37e644e72c66b96364327facd9cca3")},
	})

	if err != nil {
		panic(err)
	}

	models := make([]*TxLog, 0, len(logs))
	contractAbi, err := abi.JSON(strings.NewReader(erc20.AbiABI))
	if err != nil {
		panic(err)
	}
	for idx, tLog := range logs {
		event, err := ParseEvent(&contractAbi, &tLog)
		if err != nil {
			_ = fmt.Errorf("parse event log error, er=%s", err.Error())
			continue
		}
		if event == nil {
			continue
		}

		txLog := event.ToTxLog()
		txLog.Chain = "NOVA"
		txLog.Coin = "ETH"
		txLog.OutputIndex = int64(idx)
		txLog.Timestamp = int64(0)
		txLog.Height = int64(tLog.BlockNumber)
		txLog.BlockHash = tLog.BlockHash.Hex()
		txLog.TxHash = tLog.TxHash.Hex()
		fmt.Printf("tx log=%+v\n", txLog)
		models = append(models, txLog)
	}

}

type ContractEvent interface {
	ToTxLog() *TxLog
}

type TransferEvent struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

type TxLog struct {
	Id           int64  `gorm:"primary_key,type:bigint(20)"`
	Chain        string `gorm:"index:idx_chain_coin_height,type:varchar(32)"`
	Coin         string `gorm:"index:idx_chain_coin_height,type:varchar(32)"`
	TxHash       string `gorm:"uniqueIndex:idx_c_c_t_h_o_i_s_a,index:idx_tx_hash,type:varchar(128)"`
	OutputIndex  int64  `gorm:"uniqueIndex:idx_c_c_t_h_o_i_s_a,type:bigint(20);default:-1"`
	SenderAddr   string `gorm:"uniqueIndex:idx_c_c_t_h_o_i_s_a,type:varchar(128)"`
	ReceiverAddr string `gorm:"type:varchar(128)"`
	Memo         string `gorm:"type:varchar(64)"`
	Amount       string `gorm:"type:varchar(64)"`
	Fee          string `gorm:"type:varchar(64)"`
	Timestamp    int64  `gorm:"type:bigint(20)"`
	BlockHash    string `gorm:"type:varchar(128)"`
	Height       int64  `gorm:"index:idx_chain_coin_height,type:bigint(20)"`
	Sig          string `gorm:"type:varchar(128)"`
	PubKey       string `gorm:"type:varchar(64)"`
	ConfirmedNum int64  `gorm:"index:confirmed_num,type:bigint(20)"`
	CreateTime   int64  `gorm:"index:create_time,type:bigint(20)"`
	UpdateTime   int64  `gorm:"index:update_time,type:bigint(20)"`
	Status       int64  `gorm:"type:tinyint(1);default:1"`
}

func ParseTransferEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	var ev TransferEvent

	if len(log.Data) > 0 {
		err := abi.UnpackIntoInterface(&ev, TransferEventName, log.Data)
		if err != nil {
			return nil, err
		}
	} else {
		if len(log.Topics) > 3 {
			ev.Value = log.Topics[3].Big()
		}
	}

	ev.From = common.BytesToAddress(log.Topics[1].Bytes())
	ev.To = common.BytesToAddress(log.Topics[2].Bytes())
	return ev, nil
}

func (ev TransferEvent) ToTxLog() *TxLog {
	return &TxLog{
		SenderAddr:   ev.From.String(),
		ReceiverAddr: ev.To.String(),
		Amount:       ev.Value.String(),
	}
}

func ParseEvent(abi *abi.ABI, log *types.Log) (ContractEvent, error) {
	if bytes.Equal(log.Topics[0][:], TransferEventHash[:]) {
		return ParseTransferEvent(abi, log)
	}
	return nil, nil
}
