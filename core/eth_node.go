package core

import (
	"context"
	"github.com/Gearbox-protocol/third-eye/artifacts/multicall"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

type Node struct {
	Client  ethclient.ClientI
	ChainId int64
}

func (lf *Node) GetLogs(fromBlock, toBlock int64, addrs []common.Address, topics [][]common.Hash) ([]types.Log, error) {
	query := ethereum.FilterQuery{
		FromBlock: new(big.Int).SetInt64(fromBlock),
		ToBlock:   new(big.Int).SetInt64(toBlock),
		Addresses: addrs, //[]common.Address{common.HexToAddress(addr)},
		Topics:    topics,
	}
	var logs []types.Log
	var err error
	logs, err = lf.Client.FilterLogs(context.Background(), query)
	if err != nil {
		if err.Error() == QueryMoreThan10000Error ||
			strings.Contains(err.Error(), LogFilterLenError) ||
			err.Error() == LogFilterQueryTimeout {
			middle := (fromBlock + toBlock) / 2
			bottomHalfLogs, err := lf.GetLogs(fromBlock, middle-1, addrs, topics)
			if err != nil {
				return []types.Log{}, err
			}
			logs = append(logs, bottomHalfLogs...)

			topHalfLogs, err := lf.GetLogs(middle, toBlock, addrs, topics)
			if err != nil {
				return []types.Log{}, err
			}
			logs = append(logs, topHalfLogs...)
			return logs, nil
		}
	}
	return logs, err
}

func (lf *Node) GetLatestBlockNumber() int64 {
	latestBlockNum, err := lf.Client.BlockNumber(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	blockNumToReturn := int64(latestBlockNum)
	// skip 2 blocks ~30 sec latest block might reorder
	if lf.ChainId != 1337 {
		blockNumToReturn -= 2
	}
	log.Info("Lastest blocknumber", blockNumToReturn)
	return blockNumToReturn
}

func (lf *Node) GetHeader(blockNum int64) *types.Header {
	b, err := lf.Client.BlockByNumber(context.Background(), big.NewInt(blockNum))
	log.CheckFatal(err)
	return b.Header()
}

func (lf *Node) GasPrice(txHash common.Hash, baseFee *big.Int) *big.Int {
	tx, pending, err := lf.Client.TransactionByHash(context.TODO(), txHash)
	log.CheckFatal(err)
	if pending {
		log.Fatalf("Tx is pending, something not right %s", txHash.Hex())
	}
	if tx.Type() == 2 {
		return new(big.Int).Add(tx.GasTipCap(), baseFee)
	} else {
		return tx.GasPrice()
	}
}

func (lf *Node) EthUsed(txHash common.Hash, baseFee *big.Int) *big.Int {
	receipt := lf.GetReceipt(txHash)
	gasUsed := big.NewInt(int64(receipt.GasUsed))
	return new(big.Int).Mul(lf.GasPrice(txHash, baseFee), gasUsed)
}

func (lf *Node) GetReceipt(txHash common.Hash) *types.Receipt {
	receipt, err := lf.Client.TransactionReceipt(context.TODO(), txHash)
	log.CheckFatal(err)
	return receipt
}

func (lf *Node) GetLogsForTransfer(queryFrom, queryTill int64, hexAddrs []common.Address, treasuryAddrTopic []common.Hash) ([]types.Log, error) {
	topics := [][]common.Hash{
		{
			Topic("Transfer(address,address,uint256)"),
		},
	}
	otherAddrTopic := []common.Hash{}
	// from treasury to other address
	logs, err := lf.GetLogs(queryFrom, queryTill, hexAddrs, append(topics, treasuryAddrTopic, otherAddrTopic))
	if err != nil {
		return logs, err
	}

	// from other address to treasury
	newLogs, err := lf.GetLogs(queryFrom, queryTill, hexAddrs, append(topics, otherAddrTopic, treasuryAddrTopic))
	if err != nil {
		return logs, err
	}
	return append(newLogs, logs...), nil
}

func getMultiCallAddr() string {
	return "0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696"
}

func getMultiCallContract(client ethclient.ClientI) *multicall.Multicall {
	contract, err := multicall.NewMulticall(common.HexToAddress(getMultiCallAddr()), client)
	log.CheckFatal(err)
	return contract
}

func MakeMultiCall(client ethclient.ClientI, blockNum int64, successRequired bool, calls []multicall.Multicall2Call) []multicall.Multicall2Result {
	contract := getMultiCallContract(client)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	var result []multicall.Multicall2Result
	var tmpCalls []multicall.Multicall2Call
	callsInd := 0
	callsLen := len(calls)
	for callsInd < callsLen {
		for i := 0; i < 20 && callsInd < callsLen; i++ {
			tmpCalls = append(tmpCalls, calls[callsInd])
			callsInd++
		}
		tmpResult, err := contract.TryAggregate(opts, successRequired, tmpCalls)
		log.CheckFatal(err)
		result = append(result, tmpResult...)
		tmpCalls = []multicall.Multicall2Call{}
	}
	return result
}
