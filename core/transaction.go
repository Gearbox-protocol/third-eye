package core

import (
	"math/big"
)

type (
	Transaction struct {
		Hash             string   `json:"hash"`
		Nonce            int64    `json:"nonce"`
		TransactionIndex int64    `json:"transaction_index"`
		From             string   `json:"from_addr"`
		To               string   `json:"to_addr"` // nil means contract creation
		Value            *big.Int `json:"value"`
		Gas              int64    `json:"gas_limit"`
		GasPrice         *big.Int `json:"gas_price"`
		Data             []byte   `json:"data"`
		BlockNum         int64    `json:"block_num"`
		Timestamp        int64    `json:"timestamp"`
	}
)
