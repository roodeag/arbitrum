package vm

import "github.com/roodeag/arbitrum/common"

var (
	PrecompiledContractsArbitrum = make(map[common.Address]PrecompiledContract)
	PrecompiledAddressesArbitrum []common.Address
)
