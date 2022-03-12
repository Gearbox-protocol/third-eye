package core

import (
	"math/big"
)

const LogFilterLenError = "Log response size exceeded. You can make eth_getLogs requests with up to a 2K block range and no limit on the response size, or you can request any block range with a cap of 10K logs in the response."
const QueryMoreThan10000Error = "query returned more than 10000 results"
const LogFilterQueryTimeout = "Query timeout exceeded. Consider reducing your block range."
const NoOfBlocksPerMin int64 = 5
const NoOfBlocksPerHr int64 = NoOfBlocksPerMin * 60

var WETHPrice, USDCPrice *big.Int

func init() {
	WETHPrice, _ = new(big.Int).SetString("1000000000000000000", 10)
	USDCPrice, _ = new(big.Int).SetString("100000000", 10)
}
