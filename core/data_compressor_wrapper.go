package core

import (
	"context"
	"fmt"
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor"
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"sort"
	"strconv"
	"sync"
)

type DataCompressorWrapper struct {
	mu *sync.Mutex
	// blockNumbers of dc in asc order
	DCBlockNum     []int64
	BlockNumToName map[int64]string
	dcOldKovan     *dataCompressor.DataCompressor
	dcMainnet      *mainnet.DataCompressor
	NameToAddr     map[string]string
	client         ethclient.ClientI
	testing        *DCTesting
}

var OLDKOVAN = "OLDKOVAN"
var MAINNET = "MAINNET"
var TESTING = "TESTING"

func NewDataCompressorWrapper(client ethclient.ClientI) *DataCompressorWrapper {
	return &DataCompressorWrapper{
		mu:             &sync.Mutex{},
		BlockNumToName: make(map[int64]string),
		NameToAddr:     make(map[string]string),
		client:         client,
		testing: &DCTesting{
			calls: map[int64]*DCCalls{},
		},
	}
}

func (dcw *DataCompressorWrapper) SetCalls(blockNum int64, calls *DCCalls) {
	dcw.testing.calls[blockNum] = calls
}

func (dcw *DataCompressorWrapper) getDataCompressorIndex(blockNum int64) string {
	var name string
	for _, num := range dcw.DCBlockNum {
		// dc should be deployed before it is queried
		if num < blockNum {
			name = dcw.BlockNumToName[num]
		} else {
			break
		}
	}
	return name
}

// the data compressor are added in increasing order of blockNum
func (dcw *DataCompressorWrapper) AddDataCompressor(blockNum int64, addr string) {
	if len(dcw.DCBlockNum) > 0 && dcw.DCBlockNum[len(dcw.DCBlockNum)-1] >= blockNum {
		log.Fatal("Current dc added at :%v, new dc:%s added at %d  ", dcw.DCBlockNum, addr, blockNum)
	}
	chainId, err := dcw.client.ChainID(context.TODO())
	log.CheckFatal(err)
	var key string
	if chainId.Int64() == 1 {
		key = MAINNET
	} else if chainId.Int64() == 42 {
		switch len(dcw.DCBlockNum) {
		case 0:
			key = OLDKOVAN
		case 1:
			key = MAINNET
		}
	} else {
		key = TESTING
	}
	dcw.BlockNumToName[blockNum] = key
	dcw.NameToAddr[key] = addr
	dcw.DCBlockNum = append(dcw.DCBlockNum, blockNum)
}

func (dcw *DataCompressorWrapper) LoadMultipleDC(multiDCs interface{}) {
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
		dcw.AddDataCompressor(blockNum, dcAddr.(string))
	}
}

func (dcw *DataCompressorWrapper) GetCreditAccountDataExtended(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (mainnet.DataTypesCreditAccountDataExtended, error) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case OLDKOVAN:
		dcw.setOldKovan()
		data, err := dcw.dcOldKovan.GetCreditAccountDataExtended(opts, creditManager, borrower)
		if err != nil {
			log.Fatal(err)
		}
		latestFormat := mainnet.DataTypesCreditAccountDataExtended{
			Addr:                       data.Addr,
			Borrower:                   data.Borrower,
			InUse:                      data.InUse,
			CreditManager:              data.CreditManager,
			UnderlyingToken:            data.UnderlyingToken,
			BorrowedAmountPlusInterest: data.BorrowedAmountPlusInterest,
			TotalValue:                 data.TotalValue,
			HealthFactor:               data.HealthFactor,
			BorrowRate:                 data.BorrowRate,

			RepayAmount:           data.RepayAmount,
			LiquidationAmount:     data.LiquidationAmount,
			CanBeClosed:           data.CanBeClosed,
			BorrowedAmount:        data.BorrowedAmount,
			CumulativeIndexAtOpen: data.CumulativeIndexAtOpen,
			Since:                 data.Since,
		}
		for _, balance := range data.Balances {
			latestFormat.Balances = append(latestFormat.Balances, mainnet.DataTypesTokenBalance{
				Token:   balance.Token,
				Balance: balance.Balance,
			})
		}
		return latestFormat, err
	case MAINNET:
		dcw.setMainnet()
		return dcw.dcMainnet.GetCreditAccountDataExtended(opts, creditManager, borrower)
	case TESTING:
		return dcw.testing.getAccountData(opts.BlockNumber.Int64(), fmt.Sprintf("%s_%s", creditManager, borrower))
	}
	panic(fmt.Sprintf("data compressor number %s not found for credit account data", key))
}

func (dcw *DataCompressorWrapper) GetCreditManagerData(opts *bind.CallOpts, _creditManager common.Address, borrower common.Address) (mainnet.DataTypesCreditManagerData, error) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case OLDKOVAN:
		dcw.setOldKovan()
		data, err := dcw.dcOldKovan.GetCreditManagerData(opts, _creditManager, borrower)
		log.CheckFatal(err)
		latestFormat := mainnet.DataTypesCreditManagerData{
			Addr:               data.Addr,
			HasAccount:         data.HasAccount,
			UnderlyingToken:    data.UnderlyingToken,
			IsWETH:             data.IsWETH,
			CanBorrow:          data.CanBorrow,
			BorrowRate:         data.BorrowRate,
			MinAmount:          data.MinAmount,
			MaxAmount:          data.MaxAmount,
			MaxLeverageFactor:  data.MaxLeverageFactor,
			AvailableLiquidity: data.AvailableLiquidity,
			AllowedTokens:      data.AllowedTokens,
		}
		for _, adapter := range data.Adapters {
			latestFormat.Adapters = append(latestFormat.Adapters, mainnet.DataTypesContractAdapter{
				Adapter:         adapter.Adapter,
				AllowedContract: adapter.AllowedContract,
			})
		}
		return latestFormat, err
	case MAINNET:
		dcw.setMainnet()
		return dcw.dcMainnet.GetCreditManagerData(opts, _creditManager, borrower)
	case TESTING:
		return dcw.testing.getCMData(opts.BlockNumber.Int64(), _creditManager.Hex())
	}
	panic(fmt.Sprintf("data compressor number %s not found for credit manager data", key))
}

func (dcw *DataCompressorWrapper) GetPoolData(opts *bind.CallOpts, _pool common.Address) (mainnet.DataTypesPoolData, error) {
	dcw.mu.Lock()
	defer dcw.mu.Unlock()
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case OLDKOVAN:
		dcw.setOldKovan()
		data, err := dcw.dcOldKovan.GetPoolData(opts, _pool)
		log.CheckFatal(err)
		latestFormat := mainnet.DataTypesPoolData{
			Addr:                   data.Addr,
			IsWETH:                 data.IsWETH,
			UnderlyingToken:        data.UnderlyingToken,
			DieselToken:            data.DieselToken,
			LinearCumulativeIndex:  data.LinearCumulativeIndex,
			AvailableLiquidity:     data.AvailableLiquidity,
			ExpectedLiquidity:      data.ExpectedLiquidity,
			ExpectedLiquidityLimit: data.ExpectedLiquidityLimit,
			TotalBorrowed:          data.TotalBorrowed,
			DepositAPYRAY:          data.DepositAPYRAY,
			BorrowAPYRAY:           data.BorrowAPYRAY,
			DieselRateRAY:          data.DieselRateRAY,
			WithdrawFee:            data.WithdrawFee,
			CumulativeIndexRAY:     data.CumulativeIndexRAY,
			TimestampLU:            data.TimestampLU,
		}
		return latestFormat, err
	case MAINNET:
		dcw.setMainnet()
		return dcw.dcMainnet.GetPoolData(opts, _pool)
	case TESTING:
		return dcw.testing.getPoolData(opts.BlockNumber.Int64(), _pool.Hex())
	}
	panic(fmt.Sprintf("data compressor number %s not found for pool data", key))
}

func (dcw *DataCompressorWrapper) setOldKovan() {
	if dcw.dcOldKovan == nil {
		addr := dcw.NameToAddr[OLDKOVAN]
		var err error
		dcw.dcOldKovan, err = dataCompressor.NewDataCompressor(common.HexToAddress(addr), dcw.client)
		log.CheckFatal(err)
	}
}

func (dcw *DataCompressorWrapper) setMainnet() {
	if dcw.dcMainnet == nil {
		addr := dcw.NameToAddr[MAINNET]
		var err error
		dcw.dcMainnet, err = mainnet.NewDataCompressor(common.HexToAddress(addr), dcw.client)
		log.CheckFatal(err)
	}
}

func (dcw *DataCompressorWrapper) ToJson() string {
	return utils.ToJson(dcw)
}
