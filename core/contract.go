package core

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/aCL"
	"github.com/Gearbox-protocol/third-eye/artifacts/aCLTrait"
	"github.com/Gearbox-protocol/third-eye/artifacts/accountFactory"
	"github.com/Gearbox-protocol/third-eye/artifacts/addressProvider"
	"github.com/Gearbox-protocol/third-eye/artifacts/contractsRegister"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditAccount"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditFilter"
	"github.com/Gearbox-protocol/third-eye/artifacts/creditManager"
	"github.com/Gearbox-protocol/third-eye/artifacts/dieselToken"
	"github.com/Gearbox-protocol/third-eye/artifacts/gearToken"
	"github.com/Gearbox-protocol/third-eye/artifacts/linearInterestRateModel"
	"github.com/Gearbox-protocol/third-eye/artifacts/poolService"
	"github.com/Gearbox-protocol/third-eye/artifacts/priceOracle"
	"github.com/Gearbox-protocol/third-eye/artifacts/tokenMock"
	"github.com/Gearbox-protocol/third-eye/artifacts/uniswapv2Pool"
	"github.com/Gearbox-protocol/third-eye/artifacts/uniswapv2Router"
	"github.com/Gearbox-protocol/third-eye/artifacts/uniswapv3Pool"
	"github.com/Gearbox-protocol/third-eye/artifacts/wETHGateway"
	"github.com/Gearbox-protocol/third-eye/artifacts/yearnPriceFeed"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"context"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"math/big"
	"strings"
)

type Contract struct {
	DiscoveredAt int64             `gorm:"column:discovered_at" json:"discoveredAt"`
	FirstLogAt   int64             `gorm:"column:firstlog_at" json:"firstLogAt"`
	Address      string            `gorm:"primaryKey;column:address" json:"address"`
	Disabled     bool              `gorm:"column:disabled" json:"disabled"`
	ContractName string            `gorm:"column:type" json:"type"`
	Client       ethclient.ClientI `gorm:"-" json:"-"`
	ABI          *abi.ABI          `gorm:"-" json:"-"`
}

func (c *Contract) Disable() {
	c.Disabled = true
}

func NewContract(address, contractName string, discoveredAt int64, client ethclient.ClientI) *Contract {

	con := &Contract{
		ContractName: contractName,
		DiscoveredAt: discoveredAt,
		Address:      address,
		Client:       client,
	}
	con.FirstLogAt = con.DiscoverFirstLog()
	// for address provider discoveredAt is -1
	if discoveredAt == -1 {
		con.DiscoveredAt = con.FirstLogAt
	} else {
		con.DiscoveredAt = discoveredAt
	}
	return con
}

func GetAbi(contractName string) *abi.ABI {

	metadataMap := map[string]*bind.MetaData{

		// Configuration
		ACL:              aCL.ACLMetaData,
		AddressProvider:  addressProvider.AddressProviderMetaData,
		"ACLTrait":       aCLTrait.ACLTraitMetaData,
		ContractRegister: contractsRegister.ContractsRegisterMetaData,

		// Core
		AccountFactory:  accountFactory.AccountFactoryMetaData,
		"CreditAccount": creditAccount.CreditAccountMetaData,
		"WETHGateway":   wETHGateway.WETHGatewayMetaData,

		// Oracle
		PriceOracle:    priceOracle.PriceOracleMetaData,
		YearnPriceFeed: &bind.MetaData{ABI: yearnPriceFeed.YearnPriceFeedABI},

		// Pool
		CreditManager:             creditManager.CreditManagerMetaData,
		"LinearInterestRateModel": linearInterestRateModel.LinearInterestRateModelMetaData,
		CreditFilter:              &bind.MetaData{ABI: creditFilter.CreditFilterABI},
		Pool:                      poolService.PoolServiceMetaData,

		// GetUnderlyingToken
		"DieselToken":     dieselToken.DieselTokenMetaData,
		GearToken:         gearToken.GearTokenMetaData,
		"TokenMock":       tokenMock.TokenMockMetaData,
		"Uniswapv2Pool":   &bind.MetaData{ABI: uniswapv2Pool.Uniswapv2PoolABI},
		"Uniswapv3Pool":   &bind.MetaData{ABI: uniswapv3Pool.Uniswapv3PoolABI},
		"Uniswapv2Router": &bind.MetaData{ABI: uniswapv2Router.Uniswapv2RouterABI},
	}
	abiStr, ok := metadataMap[contractName]
	if !ok {
		log.Fatalf("ABI for %s doesn't exists", contractName)
	}

	abi, err := abiStr.GetAbi()
	if err != nil {
		log.Infof("Cant get ABI for %s", contractName)
		log.Fatal(err)
	}

	return abi
}

// setter
func (c *Contract) GetAbi() {
	c.ABI = GetAbi(c.ContractName)
}
func (c *Contract) SetAddress(addr string) {
	c.Address = addr
}

// Getter

func (c *Contract) GetAddress() string {
	return c.Address
}

func (c *Contract) GetName() string {
	return c.ContractName
}
func (c *Contract) IsDisabled() bool {
	return c.Disabled
}
func (c *Contract) GetFirstLog() int64 {
	return c.FirstLogAt
}
func (c *Contract) GetDiscoveredAt() int64 {
	return c.DiscoveredAt
}

// Extras

func (c *Contract) DiscoverFirstLog() int64 {

	// log.Debugf("Discovering first log of: %s\n", s.Address)
	lastBlock, err := c.Client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal("Cant get last block at discovery " + err.Error())
	}

	FirstLogAt, err := c.findFirstLogBound(1, int64(lastBlock))
	if err != nil {
		log.Fatal(c.Address, err.Error())
	}

	return FirstLogAt
}

func (c *Contract) findFirstLogBound(fromBlock, toBlock int64) (int64, error) {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			common.HexToAddress(c.Address),
		},
		Topics: [][]common.Hash{},
	}

	logs, err := c.Client.FilterLogs(context.Background(), query)
	if err != nil {
		if err.Error() == QueryMoreThan10000Error ||
			strings.Contains(err.Error(), LogFilterLenError) {
			middle := (fromBlock + toBlock) / 2

			log.Verbosef("FirstLog %d %d %d", fromBlock, middle-1, toBlock)
			foundLow, err := c.findFirstLogBound(fromBlock, middle-1)
			if err != nil && err.Error() != "no events found" {
				return 0, err
			}

			foundHigh, err := c.findFirstLogBound(middle, toBlock)
			if err != nil && err.Error() != "no events found" && err.Error() != "Cant find any events" {
				return 0, err
			}

			if foundLow == 0 && foundHigh == 0 {
				return 0, fmt.Errorf("No events was found for the contract")
			}

			if foundLow == 0 {
				return foundHigh, nil
			}

			return foundLow, nil

		}
		return 0, err
	}

	FirstLogAt := int64(0)

	for _, vLog := range logs {
		block := int64(vLog.BlockNumber)
		if block < FirstLogAt || FirstLogAt == 0 {
			FirstLogAt = block
		}
	}

	if FirstLogAt == MaxUint {
		return 0, fmt.Errorf("no events found")
	}

	return FirstLogAt, nil
}

func (c *Contract) FindLastLogBound(fromBlock, toBlock int64, topics []common.Hash) (int64, error) {

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{
			common.HexToAddress(c.Address),
		},
		Topics: [][]common.Hash{
			topics,
		},
	}
	logs, err := c.Client.FilterLogs(context.Background(), query)
	if err != nil {
		if err.Error() == QueryMoreThan10000Error ||
			strings.Contains(err.Error(), LogFilterLenError) {
			middle := (fromBlock + toBlock) / 2
			foundHigh, err := c.FindLastLogBound(middle, toBlock, topics)
			if err != nil {
				return 0, err
			}
			if foundHigh != 0 {
				return foundHigh, nil
			}
			foundLow, err := c.FindLastLogBound(fromBlock, middle-1, topics)
			if err != nil {
				return 0, err
			}
			if foundLow != 0 {
				return foundLow, nil
			}
		}
		return 0, err
	}

	logLen := len(logs)
	if logLen > 0 {
		return int64(logs[logLen-1].BlockNumber), nil
	} else {
		return 0, nil
	}
}

func (c *Contract) UnpackLogIntoMap(out map[string]interface{}, event string, txLog types.Log) error {
	if txLog.Topics[0] != c.ABI.Events[event].ID {
		return fmt.Errorf("event signature mismatch")
	}
	if len(txLog.Data) > 0 {
		if err := c.ABI.UnpackIntoMap(out, event, txLog.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range c.ABI.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	return abi.ParseTopicsIntoMap(out, indexed, txLog.Topics[1:])
}

func (c *Contract) ParseEvent(eventName string, txLog *types.Log) (string, *Json) {
	data := map[string]interface{}{}
	if eventName == "TransferAccount" && len(txLog.Data) > 0 {
		data = map[string]interface{}{
			"oldOwner": common.BytesToAddress(txLog.Data[:32]).Hex(),
			"newOwner": common.BytesToAddress(txLog.Data[32:]).Hex(),
		}
	} else {
		if err := c.UnpackLogIntoMap(data, eventName, *txLog); err != nil {
			log.Fatal(err)
		}
	}
	// add order
	var argNames []interface{}
	for _, input := range c.ABI.Events[eventName].Inputs {
		argNames = append(argNames, input.Name)
	}
	data["_order"] = argNames
	jsonData := Json(data)
	jsonData.CheckSumAddress()
	return c.ABI.Events[eventName].Sig, &jsonData
}
