package dc_wrapper

import (
	"context"
	"fmt"
	"sort"
	"strconv"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/test"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type DataCompressorWrapper struct {
	mu *sync.Mutex
	// blockNumbers of dc in asc order
	DCBlockNum         []int64
	BlockNumToName     map[int64]string
	discoveredAtToAddr map[int64]common.Address
	//
	v2DC    map[int64]*dataCompressorv2.DataCompressorv2
	v1DC    *MainnetDC
	testing *DCTesting

	client core.ClientI
}

var DCV1 = "DCV1"
var DCV2 = "DCV2"
var TESTING = "TESTING"

func NewDataCompressorWrapper(client core.ClientI) *DataCompressorWrapper {
	return &DataCompressorWrapper{
		mu:                 &sync.Mutex{},
		BlockNumToName:     make(map[int64]string),
		discoveredAtToAddr: make(map[int64]common.Address),
		client:             client,
		v1DC:               NewMainnetDC(client),
		v2DC:               map[int64]*dataCompressorv2.DataCompressorv2{},
		testing: &DCTesting{
			calls:  map[int64]*test.DCCalls{},
			client: client,
		},
	}
}

// testing
func (dcw *DataCompressorWrapper) SetCalls(blockNum int64, calls *test.DCCalls) {
	dcw.testing.calls[blockNum] = calls
}

func (dcw *DataCompressorWrapper) addDataCompressor(blockNum int64, addr string) {
	if len(dcw.DCBlockNum) > 0 && dcw.DCBlockNum[len(dcw.DCBlockNum)-1] >= blockNum {
		log.Fatal("Current dc added at :%v, new dc:%s added at %d  ", dcw.DCBlockNum, addr, blockNum)
	}
	chainId, err := dcw.client.ChainID(context.TODO())
	log.CheckFatal(err)
	var key string
	if chainId.Int64() == 1 || chainId.Int64() == 5 { // or for goerli
		switch len(dcw.DCBlockNum) {
		case 0:
			key = DCV1
		case 1, 2: // 2 is for goerli added datacompressor v2, as datacompressor added second time
			key = DCV2
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

// the data compressor are added in increasing order of blockNum
func (dcw *DataCompressorWrapper) AddDataCompressor(blockNum int64, addr string) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	dcw.addDataCompressor(blockNum, addr)
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
		dcAddr := dcMap[k]
		dcw.addDataCompressor(blockNum, dcAddr.(string))
	}
}

func (dcw *DataCompressorWrapper) GetCreditAccountData(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (dcv2.CreditAccountData, error) {
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key, discoveredAt := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case DCV2:
		dcw.setv2DC(discoveredAt)
		return dcw.v2DC[discoveredAt].GetCreditAccountData(opts, creditManager, borrower)
	case DCV1:
		dcw.setv1DC(discoveredAt)
		return dcw.v1DC.GetCreditAccountData(opts, creditManager, borrower)
	case TESTING:
		return dcw.testing.getAccountData(opts.BlockNumber.Int64(), fmt.Sprintf("%s_%s", creditManager, borrower))
	}
	panic(fmt.Sprintf("data compressor number %s not found for credit account data extended", key))
}

func (dcw *DataCompressorWrapper) GetCreditManagerData(opts *bind.CallOpts, _creditManager common.Address, borrower common.Address) (dcv2.CreditManagerData, error) {
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key, discoveredAt := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case DCV2:
		dcw.setv2DC(discoveredAt)
		return dcw.v2DC[discoveredAt].GetCreditManagerData(opts, _creditManager)
	case DCV1:
		dcw.setv1DC(discoveredAt)
		return dcw.v1DC.GetCreditManagerData(opts, _creditManager, borrower)
	case TESTING:
		return dcw.testing.getCMData(opts.BlockNumber.Int64(), _creditManager.Hex())
	}
	return dcv2.CreditManagerData{}, nil
}

func (dcw *DataCompressorWrapper) GetPoolData(opts *bind.CallOpts, _pool common.Address) (dcv2.PoolData, error) {
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key, discoveredAt := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case DCV2:
		dcw.setv2DC(discoveredAt)
		return dcw.v2DC[discoveredAt].GetPoolData(opts, _pool)
	case DCV1:
		dcw.setv1DC(discoveredAt)
		return dcw.v1DC.GetPoolData(opts, _pool)
	case TESTING:
		return dcw.testing.getPoolData(opts.BlockNumber.Int64(), _pool.Hex())
	}
	panic(fmt.Sprintf("data compressor number %s not found for pool data", key))
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

func (dcw *DataCompressorWrapper) setv1DC(discoveredAt int64) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	addr := dcw.discoveredAtToAddr[discoveredAt]
	dcw.v1DC.SetAddr(addr)
}

func (dcw *DataCompressorWrapper) AddCreditManagerToFilter(cmAddr, cfAddr string) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	dcw.v1DC.AddCreditManagerToFilter(cmAddr, cfAddr)
}

func (dcw *DataCompressorWrapper) setv2DC(discoveredAt int64) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	if dcw.v2DC[discoveredAt] == nil {
		addr := dcw.discoveredAtToAddr[discoveredAt]
		contractv2DC, err := dcv2.NewDataCompressorv2(addr, dcw.client)
		log.CheckFatal(err)
		dcw.v2DC[discoveredAt] = contractv2DC
	}
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
