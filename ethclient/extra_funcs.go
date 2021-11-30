// /*
//  * Gearbox monitoring
//  * Copyright (c) 2021. Mikael Lazarev
//  *
//  */

package ethclient

import (
	// 	"context"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/log"
	// 	"github.com/Gearbox-protocol/third-eye/core"
	// 	"github.com/ethereum/go-ethereum"
	// 	"github.com/ethereum/go-ethereum/common"
	// 	"github.com/ethereum/go-ethereum/core/types"
	// 	// "github.com/ethereum/go-ethereum/ethclient"
	// 	"fmt"
	// 	"math/big"
)

func NewEthClient(config *config.Config) *Client {
	client, err := Dial(config.EthProvider)
	if err != nil {
		log.Fatal(err)
	}
	return client
}

// func (rc *Client) GetLastBlock() (int64, error) {
// 	result, err := rc.client.BlockNumber(context.Background())
// 	return int64(result), err
// }

// func (rc *Client) GetBlockByNumber(blockNum int64) (*core.Block, error) {

// 	block, err := rc.client.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	query := ethereum.FilterQuery{
// 		FromBlock: big.NewInt(blockNum),
// 		ToBlock:   big.NewInt(blockNum),
// 		Addresses: []common.Address{},
// 		Topics:    [][]common.Hash{},
// 	}

// 	rawLogs, err := rc.client.FilterLogs(context.Background(), query)
// 	if err != nil {
// 		return nil, fmt.Errorf("Cant get rawLogs for address provider discovery %s", err)
// 	}

// 	transactions := block.Transactions()
// 	txcount := len(block.Transactions())
// 	txs := make([]core.Transaction, 0)
// 	timestamp := block.Time()

// 	deployments := make(map[string]core.Transaction)

// 	for i := 0; i < txcount; i++ {
// 		rawTx := transactions[i]

// 		// Get sender address
// 		msg, err := rawTx.AsMessage(types.NewLondonSigner(rawTx.ChainId()), nil)
// 		if err != nil {
// 			log.Info("GetBlock", rawTx, err)
// 			continue
// 		}

// 		from := msg.From().Hex()

// 		to := "0x0"
// 		if rawTx.To() != nil {
// 			to = rawTx.To().Hex()
// 		}

// 		tx := core.Transaction{
// 			Hash:      rawTx.Hash().Hex(),
// 			Nonce:     int64(rawTx.Nonce()),
// 			From:      from,
// 			To:        to,
// 			Gas:       int64(rawTx.Gas()),
// 			GasPrice:  rawTx.GasPrice(),
// 			Value:     rawTx.Value(),
// 			Data:      rawTx.Data(),
// 			BlockNum:  blockNum,
// 			Timestamp: int64(timestamp),
// 		}

// 		if rawTx.To() != nil {
// 			txs = append(txs, tx)

// 		} else {
// 			receipt, err := rc.client.TransactionReceipt(context.Background(), rawTx.Hash())
// 			if err != nil {
// 				return nil, fmt.Errorf("Cant get tx receipt", err)
// 			}

// 			deployments[receipt.ContractAddress.Hex()] = tx
// 		}

// 	}

// 	return &core.Block{
// 		BlockNumber:       blockNum,
// 		Timestamp:         int64(block.Time()),
// 		Hash:              block.Hash().Hex(),
// 		Deployments:       deployments,
// 		Logs:              rawLogs,
// 		//  PoolStat:          make([]core.PoolStat, 0),
// 		//  CreditManagerStat: make([]core.CreditManagerStat, 0),
// 		//  Operations:        make([]core.Operation, 0),
// 		//  PriceFeed:         make([]core.PriceItem, 0),
// 		//  CreditOperation:   make([]core.CreditOperation, 0),
// 		//  CSSnapshots:       make([]core.CreditSessionSnapshot, 0),
// 	}, nil
// }

//  func (rc *Client) GetLogsByNumber(blockNum int64) (*core.Block, error) {

// 	block, err := rc.client.BlockByNumber(context.Background(), big.NewInt(blockNum))
// 	if err != nil {
// 		return nil, err
// 	}

// 	query := ethereum.FilterQuery{
// 		FromBlock: big.NewInt(blockNum),
// 		ToBlock:   big.NewInt(blockNum),
// 		Addresses: []common.Address{},
// 		Topics:    [][]common.Hash{},
// 	}

// 	rawLogs, err := rc.client.FilterLogs(context.Background(), query)
// 	if err != nil {
// 		return nil, fmt.Errorf("Cant get rawLogs for address provider discovery", err)
// 	}

// 	return &core.Block{
// 		BlockNumber:       blockNum,
// 		Timestamp:         int64(block.Time()),
// 		Hash:              block.Hash().Hex(),
// 		Logs:              rawLogs,
// 		//  PoolStat:          make([]core.PoolStat, 0),
// 		//  CreditManagerStat: make([]core.CreditManagerStat, 0),
// 		//  Operations:        make([]core.Operation, 0),
// 		//  PriceFeed:         make([]core.PriceItem, 0),
// 		//  CreditOperation:   make([]core.CreditOperation, 0),
// 		//  CSSnapshots:       make([]core.CreditSessionSnapshot, 0),
// 	}, nil
// }

//  func (rc *Client) GetGasPaid(txHash string) *big.Int {
// 	tx, _, err := rc.client.TransactionByHash(context.Background(), common.HexToHash(txHash))
// 	if err != nil {
// 		log.Fatal("Cant get transaction by hash: ", tx)
// 	}

// 	receipt, err := rc.client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
// 	if err != nil {
// 		log.Fatal("Cant get receipts by hash: ", tx)
// 	}

// 	gasUsed := new(big.Int).Mul(tx.GasPrice(), big.NewInt(int64(receipt.GasUsed)))
// 	return gasUsed

// }
