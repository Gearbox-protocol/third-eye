package core

import (
	"fmt"
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor"
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"sort"
)

type DataCompressorWrapper struct {
	// blockNumbers of dc in asc order
	dcBlockNum []int64
	dcAddr     map[int64]string
	dc0        *dataCompressor.DataCompressor
	dc1        *mainnet.DataCompressor
	client     *ethclient.Client
}

func NewDataCompressorWrapper(client *ethclient.Client) *DataCompressorWrapper {
	return &DataCompressorWrapper{
		dcAddr: make(map[int64]string),
		client: client,
	}
}

func (dcw *DataCompressorWrapper) getDataCompressorIndex(blockNum int64) int {
	latestSmallDCIndex := -1
	for index, num := range dcw.dcBlockNum {
		// dc should be deployed before it is queried
		if num < blockNum {
			latestSmallDCIndex = index
		} else {
			break
		}
	}
	return latestSmallDCIndex
}

func (dcw *DataCompressorWrapper) AddDataCompressor(blockNum int64, addr string) {
	dcw.dcAddr[blockNum] = addr
	dcw.dcBlockNum = append(dcw.dcBlockNum, blockNum)
	arr := dcw.dcBlockNum
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	dcw.dcBlockNum = arr
}

func (dcw *DataCompressorWrapper) GetCreditAccountDataExtended(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (mainnet.DataTypesCreditAccountDataExtended, error) {
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case 0:
		dcw.setDC0()
		data, err := dcw.dc0.GetCreditAccountDataExtended(opts, creditManager, borrower)
		log.CheckFatal(err)
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
	case 1:
		dcw.setDC1()
		return dcw.dc1.GetCreditAccountDataExtended(opts, creditManager, borrower)
	}
	panic(fmt.Sprintf("data compressor number %d not found for credit account data", key))
}

func (dcw *DataCompressorWrapper) GetCreditManagerData(opts *bind.CallOpts, _creditManager common.Address, borrower common.Address) (mainnet.DataTypesCreditManagerData, error) {
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case 0:
		dcw.setDC0()
		data, err := dcw.dc0.GetCreditManagerData(opts, _creditManager, borrower)
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
	case 1:
		dcw.setDC1()
		return dcw.dc1.GetCreditManagerData(opts, _creditManager, borrower)
	}
	panic(fmt.Sprintf("data compressor number %d not found for credit manager data", key))
}

func (dcw *DataCompressorWrapper) GetPoolData(opts *bind.CallOpts, _pool common.Address) (mainnet.DataTypesPoolData, error) {
	if opts == nil || opts.BlockNumber == nil {
		panic("opts or blockNumber is nil")
	}
	key := dcw.getDataCompressorIndex(opts.BlockNumber.Int64())
	switch key {
	case 0:
		dcw.setDC0()
		data, err := dcw.dc0.GetPoolData(opts, _pool)
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
	case 1:
		dcw.setDC1()
		return dcw.dc1.GetPoolData(opts, _pool)
	}
	panic(fmt.Sprintf("data compressor number %d not found for pool data", key))
}

func (dcw *DataCompressorWrapper) setDC0() {
	if dcw.dc0 == nil {
		addr := dcw.dcAddr[dcw.dcBlockNum[0]]
		var err error
		dcw.dc0, err = dataCompressor.NewDataCompressor(common.HexToAddress(addr), dcw.client)
		log.CheckFatal(err)
	}
}

func (dcw *DataCompressorWrapper) setDC1() {
	if dcw.dc1 == nil {
		addr := dcw.dcAddr[dcw.dcBlockNum[1]]
		var err error
		dcw.dc1, err = mainnet.NewDataCompressor(common.HexToAddress(addr), dcw.client)
		log.CheckFatal(err)
	}
}
