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
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	dcv3 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressorv3"
	"github.com/Gearbox-protocol/sdk-go/artifacts/globalAccountCompressor"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/artifacts/poolCompressor"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

var ContractNameToCompressortype = map[string]CompressorType{
	// "POOL_COMPRESSOR": dc_wrapper.POOL_COMPRESSOR,
	"MARKET_COMPRESSOR":          MARKET_COMPRESSOR,
	"CREDIT_ACCOUNT_COMPRESSOR":  CREDIT_ACCOUNT_COMPRESSOR,
	"POOL_COMPRESSOR":            POOL_COMPRESSOR,
	"GLOBAL::ACCOUNT_COMPRESSOR": GLOBAL_ACCOUNT_COMPRESSOR,
	"GLOBAL::MARKET_COMPRESSOR":  GLOBAL_MARKET_COMPRESSOR,
}

func isCType(c CompressorType) bool {
	for _, v := range ContractNameToCompressortype {
		if c == v {
			return true
		}
	}
	return false
}
func compressors() []CompressorType {
	compressors := []CompressorType{}
	for _, v := range ContractNameToCompressortype {
		compressors = append(compressors, v)
	}
	return compressors
}
func (dcw *DataCompressorWrapper) AddCompressorType(addr common.Address, cType CompressorType, discoveredAt int64) {
	if !isCType(cType) {
		log.Fatal("ctype is wrong, ", cType)
	}
	if len(dcw.compressorByBlock[cType]) > 0 {
		last := dcw.compressorByBlock[cType][len(dcw.compressorByBlock[cType])-1]
		if last.block > discoveredAt {
			log.Fatalf("last %s has blocknum more than (%s) and new addr is %s", last, discoveredAt, addr)
		}
	}
	dcw.compressorByBlock[cType] = append(dcw.compressorByBlock[cType], addressAndBlock{
		address: addr,
		block:   discoveredAt,
	})
}

func (dcw *DataCompressorWrapper) LoadMultipleDC(multiDCs interface{}) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	dcMap, ok := (multiDCs).(map[string]interface{})
	if !ok {
		log.Fatalf("Converting address provider() details for dc to map failed %v", multiDCs)
	}

	blockNums := []int64{}
	{
		blockMap := map[int64]struct{}{}
		for k := range dcMap {
			splits := strings.Split(k, "_")
			blockNum, err := strconv.ParseInt(splits[0], 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			blockMap[blockNum] = struct{}{}
		}
		for blockNum := range blockMap {
			blockNums = append(blockNums, blockNum)
		}
		sort.Slice(blockNums, func(i, j int) bool { return blockNums[i] < blockNums[j] })
	}
	for _, blockNum := range blockNums {
		// TODO: NETWORK
		for _, suffix := range append(compressors(), "", "300") {
			key := fmt.Sprintf("%d", blockNum)
			if suffix != "" {
				key = fmt.Sprintf("%d_%s", blockNum, suffix)
			}
			if dcAddr, ok := dcMap[key].(string); ok {
				if suffix == "" {
					dcw.addDataCompressorv1v2(blockNum, dcAddr)
				} else if suffix == "300" {
					dcw.addDataCompressorv300(core.NewVersion(300), dcAddr, blockNum)
				} else {
					dcw.AddCompressorType(common.HexToAddress(dcAddr), suffix, blockNum)
				}
			}
		}
		//
	}
}

type CompressorType string

const (
	// TODO : NETWORK
	MARKET_COMPRESSOR         CompressorType = "MARKET"
	POOL_COMPRESSOR           CompressorType = "POOL"
	CREDIT_ACCOUNT_COMPRESSOR CompressorType = "ACCOUNT"
	GLOBAL_MARKET_COMPRESSOR  CompressorType = "GLOBALMARKET"
	GLOBAL_ACCOUNT_COMPRESSOR CompressorType = "GLOBALACCOUNT"
)

// checks in versionToAddress for v300 and then checks in addrByBlock for v1,v2
func (dcw *DataCompressorWrapper) GetKeyAndAddress(version core.VersionType, blockNum int64, cTypes []CompressorType) (string, common.Address, CompressorType) {
	if blockNum == 0 {
		log.Fatal("blockNum can't be zero")
	}
	// v300
	if version.MoreThanEq(core.NewVersion(300)) {
		// v310
		for _, cType := range cTypes {
			for i := len(dcw.compressorByBlock[cType]) - 1; i >= 0; i-- {
				compressorDetails := dcw.compressorByBlock[cType][i]
				if compressorDetails.block <= blockNum {
					return DCV310, compressorDetails.address, cType
				}
			}
		}
		// v300
		arr := dcw.versionToAddress[version]
		for i := len(arr) - 1; i >= 0; i-- {
			if arr[i].block <= blockNum {
				return DCV3, arr[i].address, ""
			}
		}
		return NODC, core.NULL_ADDR, ""
	}
	// for v2, v1
	key, discoveredAt := dcw.getDataCompressorIndex(blockNum)
	return key, dcw.getDCAddr(discoveredAt), ""
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

func (dcw *DataCompressorWrapper) Retry(blockNum int64, account common.Address, v3Pods []dataCompressorv3.PriceOnDemand, v3PodCalls []multicall.Multicall2Call) (dc.CreditAccountCallData, error) {
	key, dcAddr, compressorType := dcw.GetKeyAndAddress(core.NewVersion(300), blockNum, []CompressorType{CREDIT_ACCOUNT_COMPRESSOR, GLOBAL_ACCOUNT_COMPRESSOR})
	opts := &bind.CallOpts{BlockNumber: big.NewInt(blockNum)}
	switch key {
	case NODC:
		return dc.CreditAccountCallData{}, NO_DC_FOUND_ERR
	case DCV3:
		con, err := dataCompressorv3.NewDataCompressorv3(dcAddr, dcw.client)
		log.CheckFatal(err)
		data, err := con.GetCreditAccountData(opts, account, v3Pods)
		if err != nil || !data.IsSuccessful {
			log.Warn("after retry, getCreditAccoutn data is still not successful", blockNum, account)
			return dc.CreditAccountCallData{}, err
		}
		return dc.GetAccountDataFromDCCall(dcw.client, core.NULL_ADDR, blockNum, data)
	case DCV310:
		abiType := func() string {
			if compressorType == CREDIT_ACCOUNT_COMPRESSOR {
				return "CreditAccountCompressor"
			} else if compressorType == GLOBAL_ACCOUNT_COMPRESSOR {
				return "GlobalAccountCompressor"
			}
			log.Fatal("compressor type not found")
			return ""
		}()
		callData, err := core.GetAbi(abiType).Pack("getCreditAccountData", account)
		log.CheckFatal(err)
		v3PodCalls = append(v3PodCalls, multicall.Multicall2Call{
			Target:   dcAddr,
			CallData: callData,
		})
		ans := core.MakeMultiCall(dcw.client, blockNum, false, v3PodCalls)
		data := ans[len(ans)-1]
		if !data.Success {
			log.Warn("after retry, getCreditAccoutn data is still not successful", blockNum, account)
			return dc.CreditAccountCallData{}, err
		}
		out, err := core.GetAbi(abiType).Unpack("getCreditAccountData", data.ReturnData)
		if err != nil {
			return dc.CreditAccountCallData{}, err
		}
		var accountData creditAccountCompressor.CreditAccountData
		switch compressorType {
		case CREDIT_ACCOUNT_COMPRESSOR:
			accountData = *abi.ConvertType(out[0], new(creditAccountCompressor.CreditAccountData)).(*creditAccountCompressor.CreditAccountData)
		case GLOBAL_ACCOUNT_COMPRESSOR:
			x := abi.ConvertType(out[0], new(globalAccountCompressor.CreditAccountData)).(*globalAccountCompressor.CreditAccountData)
			accountData = dc.Convert(x)
		}
		if !accountData.Success {
			log.Warn("Not success v3", blockNum, account)
			// return dc.CreditAccountCallData{}, err
		}
		return AddFieldsToAccountv310(dcw.client, blockNum, accountData)
	}
	return dc.CreditAccountCallData{}, fmt.Errorf("data compressor number %s not found for credit account data", key)
}

// blockNum can't be zero

func (dcw *DataCompressorWrapper) GetCreditAccountData(version core.VersionType, blockNum int64, creditManager common.Address, borrower common.Address, account common.Address) (
	call multicall.Multicall2Call,
	resultFn func([]byte) (dc.CreditAccountCallData, error),
	errReturn error) {
	//
	key, dcAddr, compressorType := dcw.GetKeyAndAddress(version, blockNum, []CompressorType{CREDIT_ACCOUNT_COMPRESSOR, GLOBAL_ACCOUNT_COMPRESSOR})
	switch key {
	case NODC:
		errReturn = NO_DC_FOUND_ERR
	case DCV310:
		switch compressorType {
		case CREDIT_ACCOUNT_COMPRESSOR:
			data, err := core.GetAbi("CreditAccountCompressor").Pack("getCreditAccountData", account)
			call, errReturn = multicall.Multicall2Call{
				Target:   dcAddr,
				CallData: data,
			}, err
		case GLOBAL_ACCOUNT_COMPRESSOR:
			data, err := core.GetAbi("GlobalAccountCompressor").Pack("getCreditAccountData", account)
			call, errReturn = multicall.Multicall2Call{
				Target:   dcAddr,
				CallData: data,
			}, err
		}
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
			var out []interface{}
			var err error
			switch compressorType {
			case CREDIT_ACCOUNT_COMPRESSOR:
				out, err = core.GetAbi("CreditAccountCompressor").Unpack("getCreditAccountData", bytes)
			case GLOBAL_ACCOUNT_COMPRESSOR:
				out, err = core.GetAbi("GlobalAccountCompressor").Unpack("getCreditAccountData", bytes)
			}
			if err != nil {
				return dc.CreditAccountCallData{}, err
			}
			var accountData creditAccountCompressor.CreditAccountData
			switch compressorType {
			case CREDIT_ACCOUNT_COMPRESSOR:
				accountData = *abi.ConvertType(out[0], new(creditAccountCompressor.CreditAccountData)).(*creditAccountCompressor.CreditAccountData)
			case GLOBAL_ACCOUNT_COMPRESSOR:
				x := abi.ConvertType(out[0], new(globalAccountCompressor.CreditAccountData)).(*globalAccountCompressor.CreditAccountData)
				accountData = dc.Convert(x)
			}
			return AddFieldsToAccountv310(dcw.client, blockNum, accountData)
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
	key, dcAddr, _ := dcw.GetKeyAndAddress(version, blockNum, []CompressorType{MARKET_COMPRESSOR, GLOBAL_MARKET_COMPRESSOR})
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
			debt := struct {
				MinDebt *big.Int
				MaxDebt *big.Int
			}{}
			debt.MinDebt = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
			debt.MaxDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
			// cmData := *abi.ConvertType(out[0], new(debts)).(*debts)
			return dc.CMCallData{
				Addr:    _creditManager,
				MinDebt: (*core.BigInt)(debt.MinDebt),
				MaxDebt: (*core.BigInt)(debt.MaxDebt),
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

type PoolZapperInfo struct {
	Zappers     []dcv3.ZapperInfo
	Addr        common.Address
	Underlying  common.Address
	DieselToken common.Address
}

// zapper info from v300 onwards
func (dcw *DataCompressorWrapper) getZapperInfov3(blockNum int64, poolAddrs ...common.Address) ([]PoolZapperInfo, error) {
	compressor, found := dcw.GetLatestv3DC()
	if !found {
		return nil, nil
	}
	opts := &bind.CallOpts{BlockNumber: big.NewInt(blockNum)}
	con, err := dcv3.NewDataCompressorv3(compressor, dcw.client)
	log.CheckFatal(err)
	poolList := []dcv3.PoolData{}
	zapperInfo := []PoolZapperInfo{}
	if len(poolAddrs) != 0 {
		for _, addr := range poolAddrs {
			data, err := con.GetPoolData(opts, addr)
			if err != nil {
				return nil, fmt.Errorf("pool ZapperInfo not found in dc %d: %v", blockNum, poolAddrs)
			}
			poolList = append(poolList, data)
		}
	} else {
		var err error
		poolList, err = con.GetPoolsV3List(nil)
		log.CheckFatal(err)
	}
	for _, pool := range poolList {
		obj := &PoolZapperInfo{
			Addr:        pool.Addr,
			Underlying:  pool.Underlying,
			DieselToken: pool.DieselToken,
			Zappers:     pool.Zappers,
		}
		zapperInfo = append(zapperInfo, *obj)
	}
	return zapperInfo, nil
}

func (dcw *DataCompressorWrapper) GetZapperInfo(blockNum int64, poolAddrs ...common.Address) (ans []PoolZapperInfo) {
	key, _, _ := dcw.GetKeyAndAddress(core.NewVersion(300), blockNum, []CompressorType{MARKET_COMPRESSOR})
	switch key {
	case DCV310:
		data, err := dcw.getZapperInfov3(blockNum, poolAddrs...)
		if err != nil {
			// log.Warn(err) // of v3.1 pools have zapper info
		}
		return data
		// marketConfigs := GetMarketConfigurators()
		// con, err := marketCompressor.NewMarketCompressor(compressor, dcw.client)
		// log.CheckFatal(err)
		// poolAddrs = []common.Address{}
		// markets, err := con.GetMarkets(opts, marketCompressor.MarketFilter{Pools: poolAddrs})
		// log.CheckFatal(err)
		// for _, market := range markets {
		// 	log.Info(market.Configurator, market.Pool.BaseParams.Addr)
		// 	if !utils.Contains(marketConfigs, market.Configurator) {
		// 		continue
		// 	}
		// 	obj := &PoolZapperInfo{
		// 		Addr:        market.Pool.BaseParams.Addr,
		// 		Underlying:  market.Pool.Underlying,
		// 		DieselToken: market.Pool.BaseParams.Addr,
		// 	}
		// 	for _, zapper := range market.Zappers {
		// 		obj.Zappers = append(obj.Zappers, dcv3.ZapperInfo{
		// 			Zapper:   zapper.BaseParams.Addr,
		// 			TokenIn:  zapper.TokenIn.Addr,
		// 			TokenOut: zapper.TokenOut.Addr,
		// 		})
		// 	}
		// 	ans = append(ans, *obj)
		// }
		// log.Infof("%s: zapper info, v310: %d pooladdr:%v", log.DetectFuncAtStackN(2), len(data), poolAddrs)
		// return ans
	case DCV3:
		data, notv3 := dcw.getZapperInfov3(blockNum, poolAddrs...)
		log.CheckFatal(notv3)
		log.Infof("%s: zapper info, v3:%d pooladdr:%v", log.DetectFuncAtStackN(2), len(data), poolAddrs)
		return data
	default:
		// log.Fatal("No data compressor found for zapper info")
		return nil
	}
}

// blockNum can't be zero
func (dcw *DataCompressorWrapper) GetPoolData(version core.VersionType, blockNum int64, _pool common.Address) (
	call multicall.Multicall2Call,
	resultFn func([]byte) (dc.PoolCallData, error),
	errReturn error) {
	//
	// TODO : NETWORK
	key, dcAddr, compressorType := dcw.GetKeyAndAddress(version, blockNum, []CompressorType{POOL_COMPRESSOR, GLOBAL_MARKET_COMPRESSOR})
	switch key {
	case NODC:
		errReturn = NO_DC_FOUND_ERR
	case DCV310:
		switch compressorType {
		case POOL_COMPRESSOR:
			data, err := core.GetAbi("PoolCompressor").Pack("getPoolState", _pool)
			call, errReturn = multicall.Multicall2Call{
				Target:   dcAddr,
				CallData: data,
			}, err
		case GLOBAL_MARKET_COMPRESSOR:
			data, err := core.GetAbi("GlobalMarketCompressor").Pack("getPoolState", _pool)
			call, errReturn = multicall.Multicall2Call{
				Target:   dcAddr,
				CallData: data,
			}, err
		}
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
			var out []interface{}
			var err error
			switch compressorType {
			case POOL_COMPRESSOR:
				out, err = core.GetAbi("PoolCompressor").Unpack("getPoolState", bytes)
			case GLOBAL_MARKET_COMPRESSOR:
				out, err = core.GetAbi("GlobalMarketCompressor").Unpack("getPoolState", bytes)
			}
			if err != nil {
				return dc.PoolCallData{}, err
			}
			poolData := *abi.ConvertType(out[0], new(poolCompressor.PoolState)).(*poolCompressor.PoolState)
			return dc.GetPoolDataFromDCCall(poolData)
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
