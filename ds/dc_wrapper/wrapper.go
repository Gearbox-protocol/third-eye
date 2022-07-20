package dc_wrapper

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"strconv"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditAccount"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type DataCompressorWrapper struct {
	mu *sync.Mutex
	// blockNumbers of dc in asc order
	DCBlockNum     []int64
	BlockNumToName map[int64]string
	oldKovanDC     *oldKovanDC
	v2DC           map[int64]*v2DC
	dcMainnet      *mainnet.DataCompressor

	discoveredAtToAddr map[int64]string
	client             core.ClientI
	testing            *DCTesting
}

var OLDKOVAN = "OLDKOVAN"
var MAINNET = "MAINNET"
var TESTING = "TESTING"
var DCV2 = "DCV2"

func NewDataCompressorWrapper(client core.ClientI) *DataCompressorWrapper {
	return &DataCompressorWrapper{
		mu:                 &sync.Mutex{},
		BlockNumToName:     make(map[int64]string),
		discoveredAtToAddr: make(map[int64]string),
		client:             client,
		v2DC:               make(map[int64]*v2DC),
		testing: &DCTesting{
			calls: map[int64]*DCCalls{},
		},
	}
}

// testing
func (dcw *DataCompressorWrapper) SetCalls(blockNum int64, calls *DCCalls) {
	dcw.testing.calls[blockNum] = calls
}

func (dcw *DataCompressorWrapper) addDataCompressor(blockNum int64, addr string) {
	if len(dcw.DCBlockNum) > 0 && dcw.DCBlockNum[len(dcw.DCBlockNum)-1] >= blockNum {
		log.Fatal("Current dc added at :%v, new dc:%s added at %d  ", dcw.DCBlockNum, addr, blockNum)
	}
	chainId, err := dcw.client.ChainID(context.TODO())
	log.CheckFatal(err)
	var key string
	if chainId.Int64() == 1 {
		switch len(dcw.DCBlockNum) {
		case 0:
			key = MAINNET
		case 1:
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
		// 		key = MAINNET
		// 	case 2:
		// 		key = DCV2
		// }
		switch length := len(dcw.DCBlockNum); {
		case length == 0:
			key = MAINNET
		case length < 9:
			key = DCV2
		}
	} else {
		key = TESTING
	}
	dcw.BlockNumToName[blockNum] = key
	dcw.discoveredAtToAddr[blockNum] = addr
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

func (dcw *DataCompressorWrapper) GetCreditAccountDataExtended(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (mainnet.DataTypesCreditAccountDataExtended, error) {
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key, discoveredAt := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case OLDKOVAN:
		dcw.setOldKovan(discoveredAt)
		return dcw.oldKovanDC.GetCreditAccountDataExtended(opts, creditManager, borrower)
	case DCV2:
		dcw.setV2(discoveredAt)
		return dcw.v2DC[discoveredAt].GetCreditAccountData(opts, creditManager, borrower)
	case MAINNET:
		dcw.setMainnet(discoveredAt)
		return dcw.dcMainnet.GetCreditAccountDataExtended(opts, creditManager, borrower)
	case TESTING:
		return dcw.testing.getAccountData(opts.BlockNumber.Int64(), fmt.Sprintf("%s_%s", creditManager, borrower))
	}
	panic(fmt.Sprintf("data compressor number %s not found for credit account data extended", key))
}

func (dcw *DataCompressorWrapper) GetCreditAccountDataExtendedForHack(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (*mainnet.DataTypesCreditAccountDataExtended, error) {
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key, discoveredAt := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case MAINNET:
		dcw.setMainnet(discoveredAt)
		data, err := dcw.dcMainnet.GetCreditAccountData(opts, creditManager, borrower)
		if err != nil {
			return nil, err
		}
		account, err := creditAccount.NewCreditAccount(data.Addr, dcw.client)
		if err != nil {
			return nil, err
		}
		cumIndex, err := account.CumulativeIndexAtOpen(opts)
		if err != nil {
			return nil, err
		}
		borrowedAmount, err := account.BorrowedAmount(opts)
		if err != nil {
			return nil, err
		}
		return &mainnet.DataTypesCreditAccountDataExtended{
			Addr:                       data.Addr,
			Borrower:                   data.Borrower,
			InUse:                      data.InUse,
			CreditManager:              data.CreditManager,
			UnderlyingToken:            data.UnderlyingToken,
			BorrowedAmountPlusInterest: data.BorrowedAmountPlusInterest,
			TotalValue:                 data.TotalValue,
			HealthFactor:               data.HealthFactor,
			BorrowRate:                 data.BorrowRate,

			RepayAmount:           borrowedAmount,
			LiquidationAmount:     borrowedAmount,
			CanBeClosed:           false,
			BorrowedAmount:        borrowedAmount,
			CumulativeIndexAtOpen: cumIndex,
			Since:                 new(big.Int),
			Balances:              data.Balances,
		}, nil
	}
	panic(fmt.Sprintf("data compressor number %s not found for credit account data", key))
}

func (dcw *DataCompressorWrapper) GetCreditManagerData(opts *bind.CallOpts, _creditManager common.Address, borrower common.Address) (mainnet.DataTypesCreditManagerData, error) {
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key, discoveredAt := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case OLDKOVAN:
		dcw.setOldKovan(discoveredAt)
		return dcw.oldKovanDC.GetCreditManagerData(opts, _creditManager, borrower)
	case DCV2:
		dcw.setV2(discoveredAt)
		return dcw.v2DC[discoveredAt].GetCreditManagerData(opts, _creditManager)
	case MAINNET:
		dcw.setMainnet(discoveredAt)
		return dcw.dcMainnet.GetCreditManagerData(opts, _creditManager, borrower)
	case TESTING:
		return dcw.testing.getCMData(opts.BlockNumber.Int64(), _creditManager.Hex())
	}
	panic(fmt.Sprintf("data compressor number %s not found for credit manager data", key))
}

func (dcw *DataCompressorWrapper) GetPoolData(opts *bind.CallOpts, _pool common.Address) (mainnet.DataTypesPoolData, error) {
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key, discoveredAt := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case OLDKOVAN:
		dcw.setOldKovan(discoveredAt)
		return dcw.oldKovanDC.GetPoolData(opts, _pool)
	case DCV2:
		dcw.setV2(discoveredAt)
		return dcw.v2DC[discoveredAt].GetPoolData(opts, _pool)
	case MAINNET:
		dcw.setMainnet(discoveredAt)
		return dcw.dcMainnet.GetPoolData(opts, _pool)
	case TESTING:
		return dcw.testing.getPoolData(opts.BlockNumber.Int64(), _pool.Hex())
	}
	panic(fmt.Sprintf("data compressor number %s not found for pool data", key))
}

//
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

func (dcw *DataCompressorWrapper) setOldKovan(discoveredAt int64) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	if dcw.oldKovanDC == nil {
		addr := dcw.discoveredAtToAddr[discoveredAt]
		dcw.oldKovanDC = NewOldKovanDC(common.HexToAddress(addr), dcw.client)
	}
}

func (dcw *DataCompressorWrapper) setMainnet(discoveredAt int64) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	if dcw.dcMainnet == nil {
		addr := dcw.discoveredAtToAddr[discoveredAt]
		var err error
		dcw.dcMainnet, err = mainnet.NewDataCompressor(common.HexToAddress(addr), dcw.client)
		log.CheckFatal(err)
	}
}

func (dcw *DataCompressorWrapper) setV2(discoveredAt int64) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	if dcw.v2DC[discoveredAt] == nil {
		addr := dcw.discoveredAtToAddr[discoveredAt]
		dcw.v2DC[discoveredAt] = NewV2DC(common.HexToAddress(addr), dcw.client)
	}
}

func (dcw *DataCompressorWrapper) ToJson() string {
	return utils.ToJson(dcw)
}
