package tests

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
)

type TestCall struct {
	Pools    []mainnet.DataTypesPoolData                  `json:"pools"`
	CMs      []mainnet.DataTypesCreditManagerData         `json:"cms"`
	Accounts []mainnet.DataTypesCreditAccountDataExtended `json:"accounts"`
}

func (c *TestCall) Process() {
	return
}
func (c *TestEvent) Process(contractName string) types.Log {
	topic0 := core.Topic(c.Topics[0])
	c.Topics[0] = topic0.Hex()
	var topics []common.Hash
	for _, value := range c.Topics {
		topics = append(topics, common.HexToHash(value))
	}
	data, err := c.ParseData(contractName, topic0)
	log.CheckFatal(err)
	return types.Log{
		Data:    data,
		Topics:  topics,
		Address: common.HexToAddress(c.Address),
		TxHash:  common.HexToHash(c.TxHash),
	}
}

func (c *TestEvent) ParseData(contractName string, topic0 common.Hash) ([]byte, error) {
	abi := core.GetAbi(contractName)
	event, err := abi.EventByID(topic0)
	if err != nil {
		log.CheckFatal(err)
	}
	var args []interface{}
	for _, entry := range c.Data {
		var arg interface{}
		splits := strings.Split(entry, ":")
		if len(splits) == 2 {
			var ok bool
			switch splits[0] {
			case "bigint":
				arg, ok = new(big.Int).SetString(splits[1], 10)
				if !ok {
					log.Fatalf("bigint parsing failed for %s", entry)
				}
			case "string":
				arg = common.HexToAddress(entry).Hex()
			}
		} else {
			arg = common.HexToAddress(entry).Hex()
		}
		args = append(args, arg)
	}
	return event.Inputs.Pack(args...)
}

type TestEvent struct {
	Address string   `json:"address"`
	Data    []string `json:"data"`
	Topics  []string `json:"topics"`
	TxHash  string   `json:"txHash"`
}
type BlockInput struct {
	Events []TestEvent `json:"events"`
	Calls  TestCall    `json:"calls"`
}
type TestInput struct {
	Blocks    map[int64]BlockInput `json:"blocks"`
	MockFiles map[string]string    `json:"mocks"`
}
