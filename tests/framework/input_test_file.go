package framework

import (
	"encoding/json"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
	"testing"
)

type TestMask struct {
	Mask    *core.BigInt `json:"mask"`
	Account string       `json:"account"`
}

type TestCall struct {
	Pools       []core.TestPoolCallData      `json:"pools"`
	CMs         []core.TestCMCallData        `json:"cms"`
	Accounts    []core.TestAccountCallData   `json:"accounts"`
	Masks       []TestMask                   `json:"masks"`
	ExecuteOnCM map[string][]*core.KnownCall `json:"executeOnCM"`
}

func (c *TestCall) Process() {
	return
}
func (c *TestEvent) Process(contractName string) types.Log {
	topic0 := core.Topic(c.Topics[0])
	c.Topics[0] = topic0.Hex()
	var topics []common.Hash
	for _, value := range c.Topics {
		splits := strings.Split(value, ":")
		var newTopic string
		if len(splits) == 1 {
			newTopic = value
		} else {
			switch splits[0] {
			case "bigint":
				arg, ok := new(big.Int).SetString(splits[1], 10)
				if !ok {
					log.Fatalf("bigint parsing failed for %s", value)
				}
				newTopic = fmt.Sprintf("%x", arg)
			}
		}
		topics = append(topics, common.HexToHash(newTopic))
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
	if len(c.Data) == 0 {
		return []byte{}, nil
	}
	abi := core.GetAbi(contractName)
	event, err := abi.EventByID(topic0)
	log.CheckFatal(err)
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
			case "addr":
				arg = common.HexToAddress(entry).Hex()
			case "bool":
				if splits[1] == "1" {
					arg = true
				} else {
					arg = false
				}
			}
		} else {
			arg = common.HexToAddress(entry).Hex()
		}
		args = append(args, arg)
	}
	return event.Inputs.NonIndexed().Pack(args...)
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
	States    TestState            `json:"states"`
}

func (testInput *TestInput) Get(file string, addressMap core.AddressMap, t *testing.T) *SyncAdapterMock {
	filePath := fmt.Sprintf("../inputs/%s", file)
	tmpObj := TestInput{}
	utils.ReadJsonAndSetInterface(filePath, &tmpObj)
	var syncAdapterObj *SyncAdapterMock

	for key, fileName := range tmpObj.MockFiles {
		mockFilePath := fmt.Sprintf("../inputs/%s", fileName)
		if key == "syncAdapters" {
			syncAdapterObj = &SyncAdapterMock{}
			addAddressSetJson(mockFilePath, syncAdapterObj, addressMap, t)
		}
	}
	addAddressSetJson(filePath, testInput, addressMap, t)
	return syncAdapterObj
}

func addAddressSetJson(filePath string, obj interface{}, addressMap core.AddressMap, t *testing.T) {
	var mock core.Json = utils.ReadJson(filePath)
	mock.ParseAddress(t, addressMap)
	// log.Info(utils.ToJson(mock))
	b, err := json.Marshal(mock)
	if err != nil {
		t.Error(err)
	}
	utils.SetJson(b, obj)
}
