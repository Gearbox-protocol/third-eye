package dc_wrapper

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type oldKovanDC struct {
	dcOldKovan *dataCompressor.DataCompressor
}

func NewOldKovanDC(addr common.Address, client core.ClientI) *oldKovanDC {
	dc, err := dataCompressor.NewDataCompressor(addr, client)
	log.CheckFatal(err)
	return &oldKovanDC{
		dcOldKovan: dc,
	}
}

func (obj *oldKovanDC) GetPoolData(opts *bind.CallOpts, _pool common.Address) (mainnet.DataTypesPoolData, error) {
	data, err := obj.dcOldKovan.GetPoolData(opts, _pool)
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
}

func (obj *oldKovanDC) GetCreditAccountDataExtended(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (mainnet.DataTypesCreditAccountDataExtended, error) {
	data, err := obj.dcOldKovan.GetCreditAccountDataExtended(opts, creditManager, borrower)
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
}

func (obj *oldKovanDC) GetCreditManagerData(opts *bind.CallOpts, _creditManager common.Address, borrower common.Address) (mainnet.DataTypesCreditManagerData, error) {
	data, err := obj.dcOldKovan.GetCreditManagerData(opts, _creditManager, borrower)
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
}
