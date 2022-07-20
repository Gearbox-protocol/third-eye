package dc_wrapper

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorV2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type v2DC struct {
	dcV2 *dataCompressorV2.DataCompressorV2
}

func NewV2DC(addr common.Address, client core.ClientI) *v2DC {
	dc, err := dataCompressorV2.NewDataCompressorV2(addr, client)
	log.CheckFatal(err)
	return &v2DC{
		dcV2: dc,
	}
}

func (obj *v2DC) GetPoolData(opts *bind.CallOpts, _pool common.Address) (mainnet.DataTypesPoolData, error) {
	data, err := obj.dcV2.GetPoolData(opts, _pool)
	if err != nil {
		var blockNum int64
		if opts != nil {
			blockNum = opts.BlockNumber.Int64()
		}
		log.Fatal(err, blockNum)
	}
	latestFormat := mainnet.DataTypesPoolData{
		Addr:                   data.Addr,
		IsWETH:                 data.IsWETH,
		UnderlyingToken:        data.Underlying,
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
}

func (obj *v2DC) GetCreditAccountData(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (mainnet.DataTypesCreditAccountDataExtended, error) {
	data, err := obj.dcV2.GetCreditAccountData(opts, creditManager, borrower)
	if err != nil {
		log.Fatal(err)
	}
	latestFormat := mainnet.DataTypesCreditAccountDataExtended{
		Addr:                       data.Addr,
		Borrower:                   data.Borrower,
		InUse:                      data.InUse,
		CreditManager:              data.CreditManager,
		UnderlyingToken:            data.Underlying,
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
			Token:     balance.Token,
			Balance:   balance.Balance,
			IsAllowed: balance.IsAllowed,
		})
	}
	return latestFormat, err
}

func (obj *v2DC) GetCreditManagerData(opts *bind.CallOpts, _creditManager common.Address) (mainnet.DataTypesCreditManagerData, error) {
	data, err := obj.dcV2.GetCreditManagerData(opts, _creditManager)
	log.CheckFatal(err)
	latestFormat := mainnet.DataTypesCreditManagerData{
		Addr: data.Addr,
		//
		HasAccount: true,
		//
		UnderlyingToken:    data.Underlying,
		IsWETH:             data.IsWETH,
		CanBorrow:          data.CanBorrow,
		BorrowRate:         data.BorrowRate,
		MinAmount:          data.MinAmount,
		MaxAmount:          data.MaxAmount,
		MaxLeverageFactor:  data.MaxLeverageFactor,
		AvailableLiquidity: data.AvailableLiquidity,
		AllowedTokens:      data.CollateralTokens,
	}
	for _, adapter := range data.Adapters {
		latestFormat.Adapters = append(latestFormat.Adapters, mainnet.DataTypesContractAdapter{
			Adapter:         adapter.Adapter,
			AllowedContract: adapter.AllowedContract,
		})
	}
	return latestFormat, err
}
