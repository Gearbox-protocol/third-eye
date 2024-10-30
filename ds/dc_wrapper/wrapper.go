package dc_wrapper

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditAccountCompressor"
	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/marketCompressor"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type addressAndBlock struct {
	address common.Address
	block   int64
}
type DataCompressorWrapper struct {
	mu *sync.RWMutex
	// blockNumbers of dc in asc order
	DCBlockNum         []int64
	BlockNumToName     map[int64]string
	discoveredAtToAddr map[int64]common.Address
	// for v3 version to address
	versionToAddress map[core.VersionType][]addressAndBlock
	// for v1
	creditManagerToFilter map[common.Address]common.Address
	// for v310

	compressorByBlock map[CompressorType][]addressAndBlock

	//
	testing *DCTesting

	client core.ClientI
}

var DCV1 = "DCV1"
var DCV2 = "DCV2"
var TESTING = "TESTING"
var DCV3 = "DCV3"
var DCV310 = "DCV310"
var NODC = "NODC"
var NO_DC_FOUND_ERR = fmt.Errorf("No data compressor found")

func NewDataCompressorWrapper(client core.ClientI) *DataCompressorWrapper {
	return &DataCompressorWrapper{
		mu:                 &sync.RWMutex{},
		BlockNumToName:     make(map[int64]string),
		discoveredAtToAddr: make(map[int64]common.Address),
		client:             client,
		// for v1
		creditManagerToFilter: make(map[common.Address]common.Address),
		compressorByBlock:     make(map[CompressorType][]addressAndBlock),
		testing: &DCTesting{
			calls:  map[int64]*dc.DCCalls{},
			client: client,
		},
		versionToAddress: map[core.VersionType][]addressAndBlock{},
	}
}

// testing
func (dcw *DataCompressorWrapper) SetCalls(blockNum int64, calls *dc.DCCalls) {
	dcw.testing.calls[blockNum] = calls
}

func (dcw *DataCompressorWrapper) addDataCompressorv1v2(blockNum int64, addr string) {
	if len(dcw.DCBlockNum) > 0 && dcw.DCBlockNum[len(dcw.DCBlockNum)-1] >= blockNum {
		log.Fatalf("Current dc added at :%v, new dc:%s added at %d  ", dcw.DCBlockNum, addr, blockNum)
	}
	chainId, err := dcw.client.ChainID(context.TODO())
	log.CheckFatal(err)
	var key string
	if chainId.Int64() == 1 || chainId.Int64() == 7878 { //anvil fork and mainnet
		switch len(dcw.DCBlockNum) {
		case 0:
			key = DCV1
		case 1: // goerli deprecated
			// 2 is for goerli added datacompressor v2, as datacompressor added second time
			key = DCV2
		case 3: // ignore the dc 2.1 for v2.0, all data is already provided by dc2.0
			log.AMQPMsgf("New dataCompressor(%s) added", addr)
			return
		}
	} else if chainId.Int64() == 42 {
		// 	switch len(dcw.DCBlockNum) {
		// for old address provider 0xA526311C39523F60b184709227875b5f34793bD4
		// we had a datacompressor which was used while first gearbox test deployment for users, that happened in nov 2021
		// later around july 2022 in redeployed whole kovan setup where there was only 1 dc per gearbox 1
		// 	case 0:
		// 		key = OLDKOVAN
		// 	case 1:
		// 		key = DCV1
		// 	case 2:
		// 		key = DCV2
		// }
		switch length := len(dcw.DCBlockNum); {
		case length == 0:
			key = DCV1
		case length < 9:
			key = DCV2
		}
	} else {
		key = TESTING
	}
	dcw.BlockNumToName[blockNum] = key
	dcw.discoveredAtToAddr[blockNum] = common.HexToAddress(addr)
	dcw.DCBlockNum = append(dcw.DCBlockNum, blockNum)
}

// the data compressor are added in increasing order of blockNum for v1, v2
func (dcw *DataCompressorWrapper) AddDataCompressor(blockNum int64, addr string) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	dcw.addDataCompressorv1v2(blockNum, addr)
}

// for v300
func (dcw *DataCompressorWrapper) AddDataCompressorv300(version core.VersionType, addr string, blockNum int64) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	dcw.addDataCompressorv300(version, addr, blockNum)
}

func (dcw *DataCompressorWrapper) addDataCompressorv300(version core.VersionType, addr string, discoveredAt int64) {
	dcw.versionToAddress[version] = append(dcw.versionToAddress[version], addressAndBlock{
		address: common.HexToAddress(addr),
		block:   discoveredAt,
	})
	sort.Slice(dcw.versionToAddress[version], func(i, j int) bool {
		return dcw.versionToAddress[version][i].block < dcw.versionToAddress[version][j].block
	})
}

func (dcw *DataCompressorWrapper) AddCompressorType(addr common.Address, cType CompressorType, discoveredAt int64) {
	if cType != CREDIT_ACCOUNT_COMPRESSOR && cType != MARKET_COMPRESSOR {
		log.Fatal("ctype is wrong, ", cType)
	}
	if len(dcw.compressorByBlock[cType]) > 0 {
		last := dcw.compressorByBlock[cType][len(dcw.compressorByBlock[cType])-1]
		if last.block > discoveredAt {
			log.Fatalf("last %s has blocknum more than (%s) and new addr is %s", last, discoveredAt, addr)
		}
		dcw.compressorByBlock[cType] = append(dcw.compressorByBlock[cType], addressAndBlock{
			address: addr,
			block:   discoveredAt,
		})
	}
}

func (dcw *DataCompressorWrapper) LoadMultipleDC(multiDCs interface{}) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	dcMap, ok := (multiDCs).(map[string]interface{})
	if !ok {
		log.Fatalf("Converting address provider() details for dc to map failed %v", multiDCs)
	}

	var blockNums []int64
	for k := range dcMap {
		blockNum, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		blockNums = append(blockNums, blockNum)
	}
	sort.Slice(blockNums, func(i, j int) bool { return blockNums[i] < blockNums[j] })
	for _, blockNum := range blockNums {
		k := fmt.Sprintf("%d", blockNum)
		dcAddr := dcMap[k].(string)
		//
		if strings.Contains(dcAddr, "_") {
			addrs := strings.Split(dcAddr, "_")
			if addrs[1] == "300" {
				version, err := strconv.ParseInt(addrs[1], 10, 64)
				log.CheckFatal(err)
				dcw.addDataCompressorv300(core.NewVersion(int16(version)), addrs[0], blockNum)
			} else {
				dcw.AddCompressorType(common.HexToAddress(addrs[0]), CompressorType(addrs[1]), blockNum)
			}
		} else {
			dcw.addDataCompressorv1v2(blockNum, dcAddr)
		}
	}
}

type CompressorType string

const (
	MARKET_COMPRESSOR CompressorType = "MARKET"
	// CM_COMPRESSOR             CompressorType = "CM"
	CREDIT_ACCOUNT_COMPRESSOR CompressorType = "ACCOUNT"
)

// checks in versionToAddress for v300 and then checks in addrByBlock for v1,v2
func (dcw *DataCompressorWrapper) GetKeyAndAddress(version core.VersionType, blockNum int64, cType CompressorType) (string, common.Address) {
	// v300
	if version.MoreThanEq(core.NewVersion(300)) {
		// v310
		for i := len(dcw.compressorByBlock[cType]) - 1; i >= 0; i-- {
			compressorDetails := dcw.compressorByBlock[cType][i]
			if compressorDetails.block <= blockNum {
				return DCV310, compressorDetails.address
			}
		}
		// v300
		arr := dcw.versionToAddress[version]
		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i].block <= blockNum {
				return DCV3, arr[i].address
			}
		}
		return NODC, core.NULL_ADDR
	}
	// for v2, v1
	key, discoveredAt := dcw.getDataCompressorIndex(blockNum)
	return key, dcw.getDCAddr(discoveredAt)
}

// checks in versionToAddress for v300 and then checks in addrByBlock for v1,v2
func (dcw *DataCompressorWrapper) GetLatestv3DC() (common.Address, bool) {
	version := core.NewVersion(300)
	if len(dcw.versionToAddress[version]) == 0 {
		return core.NULL_ADDR, false
	}
	//
	dcs := dcw.versionToAddress[version]
	return dcs[len(dcs)-1].address, true
}

// blockNum can't be zero

func (dcw *DataCompressorWrapper) GetCreditAccountData(version core.VersionType, blockNum int64, creditManager common.Address, borrower common.Address, account common.Address) (
	call multicall.Multicall2Call,
	resultFn func([]byte) (dc.CreditAccountCallData, error),
	errReturn error) {
	//
	key, dcAddr := dcw.GetKeyAndAddress(version, blockNum, CREDIT_ACCOUNT_COMPRESSOR)
	switch key {
	case NODC:
		errReturn = NO_DC_FOUND_ERR
	case DCV310:
		data, err := core.GetAbi("CreditAccountCompressor").Pack("getCreditAccountData", account)
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case DCV3:
		data, err := core.GetAbi("DataCompressorv3").Pack("getCreditAccountData", account, []dcv3.PriceOnDemand{})
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case DCV2:
		data, err := core.GetAbi("DataCompressorV2").Pack("getCreditAccountData", creditManager, borrower)
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case DCV1:
		data, err := core.GetAbi("DataCompressorMainnet").Pack("getCreditAccountDataExtended", creditManager, borrower)
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case TESTING:
		data, err := core.GetAbi("DataCompressorMainnet").Pack("getCreditAccountDataExtended", creditManager, borrower)
		call, errReturn = multicall.Multicall2Call{
			Target:   common.HexToAddress("0x0000000000000000000000000000000000000001"),
			CallData: data,
		}, err
	default:
		panic(fmt.Sprintf("data compressor number %s not found for credit account data extended", key))
	}
	resultFn = func(bytes []byte) (dc.CreditAccountCallData, error) {
		switch key {
		case NODC:
			log.Fatal("No data compressor found for credit account data")
		case DCV310:
			out, err := core.GetAbi("CreditAccountCompressor").Unpack("getCreditAccountData", bytes)
			if err != nil {
				return dc.CreditAccountCallData{}, err
			}
			accountData := *abi.ConvertType(out[0], new(creditAccountCompressor.CreditAccountData)).(*creditAccountCompressor.CreditAccountData)
			return dcw.addFieldsToAccountv310(blockNum, accountData)
		case DCV3:
			out, err := core.GetAbi("DataCompressorv3").Unpack("getCreditAccountData", bytes)
			if err != nil {
				return dc.CreditAccountCallData{}, err
			}
			accountData := *abi.ConvertType(out[0], new(dcv3.CreditAccountData)).(*dcv3.CreditAccountData)
			return dc.GetAccountDataFromDCCall(dcw.client, core.NULL_ADDR, blockNum, accountData)
		case DCV2:
			out, err := core.GetAbi("DataCompressorV2").Unpack("getCreditAccountData", bytes)
			if err != nil {
				return dc.CreditAccountCallData{}, err
			}
			accountData := *abi.ConvertType(out[0], new(dcv2.CreditAccountData)).(*dcv2.CreditAccountData)
			return dc.GetAccountDataFromDCCall(dcw.client, core.NULL_ADDR, blockNum, accountData)
		case DCV1:
			out, err := core.GetAbi("DataCompressorMainnet").Unpack("getCreditAccountDataExtended", bytes)
			if err != nil {
				return dc.CreditAccountCallData{}, err
			}
			accountData := *abi.ConvertType(out[0], new(mainnet.DataTypesCreditAccountDataExtended)).(*mainnet.DataTypesCreditAccountDataExtended)
			return dc.GetAccountDataFromDCCall(dcw.client, dcw.creditManagerToFilter[creditManager], blockNum, accountData)
		case TESTING:
			return dcw.testing.getAccountData(blockNum, fmt.Sprintf("%s_%s", creditManager, borrower))
		}
		panic(fmt.Sprintf("data compressor number %s not found for pool data", key))
	}
	return
}

// blockNum can't be zero

func (dcw *DataCompressorWrapper) GetCreditManagerData(version core.VersionType, blockNum int64, _creditManager common.Address, cf string) (
	call multicall.Multicall2Call,
	resultFn func([]byte) (dc.CMCallData, error),
	errReturn error) {
	//
	key, dcAddr := dcw.GetKeyAndAddress(version, blockNum, MARKET_COMPRESSOR)
	switch key {
	case NODC:
		errReturn = NO_DC_FOUND_ERR
	case DCV310:
		// data, err := core.GetAbi("PoolCompressor").Pack("getCreditManagerData", _creditManager)
		// call, errReturn = multicall.Multicall2Call{
		// 	Target:   dcAddr,
		// 	CallData: data,
		// }, err
		data, err := hex.DecodeString("166bf9d9")
		log.CheckFatal(err)
		call = multicall.Multicall2Call{
			Target:   common.HexToAddress(cf),
			CallData: data,
		}
		errReturn = nil
	case DCV3:
		data, err := core.GetAbi("DataCompressorv3").Pack("getCreditManagerData", _creditManager)
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case DCV2:
		data, err := core.GetAbi("DataCompressorV2").Pack("getCreditManagerData", _creditManager)
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case DCV1:
		data, err := core.GetAbi("DataCompressorMainnet").Pack("getCreditManagerData", _creditManager, _creditManager)
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case TESTING:
		data, err := core.GetAbi("DataCompressorMainnet").Pack("getCreditManagerData", _creditManager, _creditManager)
		call, errReturn = multicall.Multicall2Call{
			Target:   common.HexToAddress("0x0000000000000000000000000000000000000001"),
			CallData: data,
		}, err
	}
	//
	resultFn = func(bytes []byte) (dc.CMCallData, error) {
		switch key {
		case NODC:
			log.Fatal("No data compressor found for credit manager data")
		case DCV310:
			out, err := core.GetAbi("CreditFacadev3").Unpack("debtLimits", bytes)
			if err != nil {
				return dc.CMCallData{}, err
			}
			type debts struct {
				MinDebt *big.Int
				MaxDebt *big.Int
			}
			cmData := *abi.ConvertType(out[0], new(debts)).(*debts)
			return dc.CMCallData{
				Addr:    _creditManager,
				MinDebt: (*core.BigInt)(cmData.MinDebt),
				MaxDebt: (*core.BigInt)(cmData.MaxDebt),
			}, nil
		case DCV3:
			out, err := core.GetAbi("DataCompressorv3").Unpack("getCreditManagerData", bytes)
			if err != nil {
				return dc.CMCallData{}, err
			}
			cmData := *abi.ConvertType(out[0], new(dcv3.CreditManagerData)).(*dcv3.CreditManagerData)
			return dc.GetCMDataFromDCCall(cmData)
		case DCV2:
			out, err := core.GetAbi("DataCompressorV2").Unpack("getCreditManagerData", bytes)
			if err != nil {
				return dc.CMCallData{}, err
			}
			cmData := *abi.ConvertType(out[0], new(dcv2.CreditManagerData)).(*dcv2.CreditManagerData)
			return dc.GetCMDataFromDCCall(cmData)
		case DCV1:
			out, err := core.GetAbi("DataCompressorMainnet").Unpack("getCreditManagerData", bytes)
			if err != nil {
				return dc.CMCallData{}, err
			}
			cmData := *abi.ConvertType(out[0], new(mainnet.DataTypesCreditManagerData)).(*mainnet.DataTypesCreditManagerData)
			return dc.GetCMDataFromDCCall(cmData)
		case TESTING:
			return dcw.testing.getCMData(blockNum, _creditManager.Hex())
		}
		panic(fmt.Sprintf("data compressor number %s not found for pool data", key))
	}
	return
}

func (dcw *DataCompressorWrapper) GetPoolListv3() ([]dcv3.PoolData, bool) {
	dcAddr, found := dcw.GetLatestv3DC()
	if !found {
		return nil, false
	}
	con, err := dcv3.NewDataCompressorv3(dcAddr, dcw.client)
	log.CheckFatal(err)
	poolList, err := con.GetPoolsV3List(nil)
	log.CheckFatal(err)
	return poolList, true
}

// blockNum can't be zero
func (dcw *DataCompressorWrapper) GetPoolData(version core.VersionType, blockNum int64, _pool common.Address) (
	call multicall.Multicall2Call,
	resultFn func([]byte) (dc.PoolCallData, error),
	errReturn error) {
	//
	key, dcAddr := dcw.GetKeyAndAddress(version, blockNum, MARKET_COMPRESSOR)
	switch key {
	case NODC:
		errReturn = NO_DC_FOUND_ERR
	case DCV310:
		data, err := core.GetAbi("MarketCompressor").Pack("getMarketData0", _pool)
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case DCV3:
		data, err := core.GetAbi("DataCompressorv3").Pack("getPoolData", _pool)
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case DCV2:
		data, err := core.GetAbi("DataCompressorV2").Pack("getPoolData", _pool)
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case DCV1:
		data, err := core.GetAbi("DataCompressorMainnet").Pack("getPoolData", _pool)
		call, errReturn = multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: data,
		}, err
	case TESTING:
		data, err := core.GetAbi("DataCompressorMainnet").Pack("getPoolData", _pool)
		call, errReturn = multicall.Multicall2Call{
			Target:   common.HexToAddress("0x0000000000000000000000000000000000000001"),
			CallData: data,
		}, err
	default:
		panic(fmt.Sprintf("data compressor number %s not found for pool data", key))
	}
	//
	resultFn = func(bytes []byte) (dc.PoolCallData, error) {
		switch key {
		case NODC:
			log.Fatal("No data compressor found for pool data")
		case DCV310:
			out, err := core.GetAbi("MarketCompressor").Unpack("getMarketData0", bytes)
			if err != nil {
				return dc.PoolCallData{}, err
			}
			poolData := *abi.ConvertType(out[0], new(marketCompressor.MarketData)).(*marketCompressor.MarketData)
			return dc.GetPoolDataFromDCCall(poolData.Pool)
		case DCV3:
			out, err := core.GetAbi("DataCompressorv3").Unpack("getPoolData", bytes)
			if err != nil {
				return dc.PoolCallData{}, err
			}
			poolData := *abi.ConvertType(out[0], new(dcv3.PoolData)).(*dcv3.PoolData)
			return dc.GetPoolDataFromDCCall(poolData)
		case DCV2:
			out, err := core.GetAbi("DataCompressorV2").Unpack("getPoolData", bytes)
			if err != nil {
				return dc.PoolCallData{}, err
			}
			poolData := *abi.ConvertType(out[0], new(dcv2.PoolData)).(*dcv2.PoolData)
			return dc.GetPoolDataFromDCCall(poolData)
		case DCV1:
			out, err := core.GetAbi("DataCompressorMainnet").Unpack("getPoolData", bytes)
			if err != nil {
				return dc.PoolCallData{}, err
			}
			poolData := *abi.ConvertType(out[0], new(mainnet.DataTypesPoolData)).(*mainnet.DataTypesPoolData)
			return dc.GetPoolDataFromDCCall(poolData)
		case TESTING:
			return dcw.testing.getPoolData(blockNum, _pool.Hex())
		}
		panic(fmt.Sprintf("data compressor number %s not found for pool data", key))
	}
	return
}

// get the last datacompressor added before blockNum
// blockNum to name
func (dcw *DataCompressorWrapper) getDataCompressorIndex(blockNum int64) (name string, discoveredAt int64) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	for _, num := range dcw.DCBlockNum {
		// dc should be deployed before it is queried
		if num < blockNum {
			name = dcw.BlockNumToName[num]
			discoveredAt = num
		} else {
			break
		}
	}
	return
}

func (dcw *DataCompressorWrapper) AddCreditManagerToFilter(cmAddr, cfAddr string) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	dcw.creditManagerToFilter[common.HexToAddress(cmAddr)] = common.HexToAddress(cfAddr)
}

func (dcw *DataCompressorWrapper) getDCAddr(discoveredAt int64) common.Address {
	dcw.mu.RLock()
	defer dcw.mu.RUnlock()
	return dcw.discoveredAtToAddr[discoveredAt]
}

func (dcw *DataCompressorWrapper) ToJson() string {
	return utils.ToJson(dcw)
}

// func (dcw *DataCompressorWrapper) GetCreditAccountDataForHack(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (*dcv2.CreditAccountData, error) {
// 	if opts == nil || opts.BlockNumber == nil {
// 		panic("opts or blockNumber is nil")
// 	}
// 	key, discoveredAt := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
// 	switch key {
// 	case MAINNET:
// 		dcw.setMainnet(discoveredAt)
// 		data, err := dcw.dcMainnet.GetCreditAccountData(opts, creditManager, borrower)
// 		if err != nil {
// 			return nil, err
// 		}
// 		account, err := creditAccount.NewCreditAccount(data.Addr, dcw.client)
// 		if err != nil {
// 			return nil, err
// 		}
// 		cumIndex, err := account.CumulativeIndexAtOpen(opts)
// 		if err != nil {
// 			return nil, err
// 		}
// 		borrowedAmount, err := account.BorrowedAmount(opts)
// 		if err != nil {
// 			return nil, err
// 		}
// 		return &dcv2.CreditAccountData{
// 			Addr:                       data.Addr,
// 			Borrower:                   data.Borrower,
// 			InUse:                      data.InUse,
// 			CreditManager:              data.CreditManager,
// 			Underlying:                 data.Underlying,
// 			BorrowedAmountPlusInterest: data.BorrowedAmountPlusInterest,
// 			TotalValue:                 data.TotalValue,
// 			HealthFactor:               data.HealthFactor,
// 			BorrowRate:                 data.BorrowRate,

// 			RepayAmount:           borrowedAmount,
// 			LiquidationAmount:     borrowedAmount,
// 			CanBeClosed:           false,
// 			BorrowedAmount:        borrowedAmount,
// 			CumulativeIndexAtOpen: cumIndex,
// 			Since:                 new(big.Int),
// 			Balances:              data.Balances,
// 		}, nil
// 	}
// 	panic(fmt.Sprintf("data compressor number %s not found for credit account data", key))
// }
