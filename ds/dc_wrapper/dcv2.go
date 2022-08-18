package dc_wrapper

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type v2DC struct {
	dcV2   *dataCompressorv2.DataCompressorv2
	client core.ClientI
	addr   string
}

func NewV2DC(addr common.Address, client core.ClientI) *v2DC {
	dc, err := dataCompressorv2.NewDataCompressorv2(addr, client)
	log.CheckFatal(err)
	return &v2DC{
		dcV2:   dc,
		client: client,
		addr:   addr.Hex(),
	}
}

func (obj *v2DC) GetPoolData(opts *bind.CallOpts, _pool common.Address) (mainnet.DataTypesPoolData, error) {
	data, err := obj.dcV2.GetPoolData(opts, _pool)
	if err != nil {
		var blockNum int64
		if opts != nil {
			blockNum = opts.BlockNumber.Int64()
		}
		log.Fatal(err, blockNum, obj.addr)
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
		// CHECK FOR: edge case
		if obj.addr == "0x47DE3e0d505B6ed8f8FA3bbB9Ab9b303E2ebCe39" {
			if core.FetchVersion(creditManager.Hex(), opts.BlockNumber.Int64(), obj.client) == 2 {
				return obj.manualAccountCall(opts, creditManager, borrower)
			}
		}
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

// there was a smartcontractbug in 8 th dcv2 contract for july 2022 kovan deployment.
// dc address "0x47DE3e0d505B6ed8f8FA3bbB9Ab9b303E2ebCe39"
// in that address for ceditaccount opened with creditmanager v2, will fail till the next dc which is added to addressprovider at 32832988
func (obj *v2DC) manualAccountCall(opts *bind.CallOpts, creditManager common.Address, borrower common.Address) (mainnet.DataTypesCreditAccountDataExtended, error) {
	calls := []multicall.Multicall2Call{}
	blockNum := opts.BlockNumber.Int64()
	cmABI := core.GetAbi("CreditManagerv2")
	///////////////////////////////
	// Phase 1
	///////////////////////////////
	//
	accountData, err := cmABI.Pack("creditAccounts", borrower)
	log.CheckFatal(err)
	calls = append(calls, multicall.Multicall2Call{
		Target:   creditManager,
		CallData: accountData,
	})
	creditFacadeData, err := cmABI.Pack("creditFacade")
	log.CheckFatal(err)
	calls = append(calls, multicall.Multicall2Call{
		Target:   creditManager,
		CallData: creditFacadeData,
	})
	tokensCountData, err := cmABI.Pack("collateralTokensCount")
	log.CheckFatal(err)
	calls = append(calls, multicall.Multicall2Call{
		Target:   creditManager,
		CallData: tokensCountData,
	})
	result := core.MakeMultiCall(obj.client, blockNum, false, calls)
	account := getAddr(result[0])
	creditFacade := getAddr(result[1])
	tokensCount := getBigInt(result[2]).Int64()

	///////////////////////////////
	// Phase 2
	///////////////////////////////
	//
	calls = calls[:0]
	borrowedAmountData, err := cmABI.Pack("calcCreditAccountAccruedInterest", account)
	log.CheckFatal(err)
	calls = append(calls, multicall.Multicall2Call{
		Target:   creditManager,
		CallData: borrowedAmountData,
	})
	facadeABI := core.GetAbi("CreditFacade")
	hfData, err := facadeABI.Pack("calcCreditAccountHealthFactor", account)
	log.CheckFatal(err)
	calls = append(calls, multicall.Multicall2Call{
		Target:   creditFacade,
		CallData: hfData,
	})
	totalValueData, err := facadeABI.Pack("calcTotalValue", account)
	log.CheckFatal(err)
	calls = append(calls, multicall.Multicall2Call{
		Target:   creditFacade,
		CallData: totalValueData,
	})
	accountABI := core.GetAbi("CreditAccount")
	cumIndexData, err := accountABI.Pack("cumulativeIndexAtOpen")
	log.CheckFatal(err)
	calls = append(calls, multicall.Multicall2Call{
		Target:   account,
		CallData: cumIndexData,
	})
	for i := int64(0); i < tokensCount; i++ {
		data, err := cmABI.Pack("collateralTokens", big.NewInt(i))
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			CallData: data,
			Target:   creditManager,
		})
	}
	result = core.MakeMultiCall(obj.client, blockNum, false, calls)
	borrowedAmount, amountWithInterest := get2BigInt(result[0])
	hf := getBigInt(result[1])
	totalValue, _ := get2BigInt(result[2])
	cumIndex := getBigInt(result[3])

	tokens := make([]common.Address, 0, len(result[4:]))
	for _, entry := range result[4:] {
		tokens = append(tokens, getAddr(entry))
	}
	///////////////////////////////////////
	// Phase 3
	///////////////////////////////////////
	//
	balances := obj.getBalances(blockNum, tokens, account, creditFacade)
	return mainnet.DataTypesCreditAccountDataExtended{
		Addr:                       account,
		Borrower:                   borrower,
		InUse:                      false,
		CreditManager:              creditManager,
		BorrowedAmountPlusInterest: amountWithInterest,
		TotalValue:                 totalValue,
		HealthFactor:               hf,

		BorrowedAmount:        borrowedAmount,
		CumulativeIndexAtOpen: cumIndex,
		Balances:              balances,
		// BorrowRate:                 data.BorrowRate,
		// Underlying:,
		// RepayAmount:           data.RepayAmount,
		// LiquidationAmount:     data.LiquidationAmount,
		// CanBeClosed:           data.CanBeClosed,
		// Since:,
	}, nil
}

func (obj *v2DC) getBalances(blockNum int64, tokens []common.Address, account, creditFacade common.Address) (balances []mainnet.DataTypesTokenBalance) {
	tokenABI := core.GetAbi("Token")
	facadeABI := core.GetAbi("CreditFacade")
	balanceData, err := tokenABI.Pack("balanceOf", account)
	calls := make([]multicall.Multicall2Call, 0, len(tokens)*2)
	log.CheckFatal(err)
	for _, token := range tokens {
		allowedTokenData, err := facadeABI.Pack("isTokenAllowed", token)
		log.CheckFatal(err)
		calls = append(calls, multicall.Multicall2Call{
			Target:   token,
			CallData: balanceData,
		}, multicall.Multicall2Call{
			Target:   creditFacade,
			CallData: allowedTokenData,
		})
	}
	result := core.MakeMultiCall(obj.client, blockNum, false, calls)
	var balance *big.Int
	for i, entry := range result {
		if entry.Success {
			if i%2 == 0 {
				amount, err := tokenABI.Unpack("balanceOf", entry.ReturnData)
				log.CheckFatal(err)
				balance = amount[0].(*big.Int)
			} else if i%2 == 1 {
				values, err := facadeABI.Unpack("isTokenAllowed", entry.ReturnData)
				log.CheckFatal(err)
				isAllowed := values[0].(bool)
				balances = append(balances, mainnet.DataTypesTokenBalance{
					Token:     tokens[i/2],
					Balance:   balance,
					IsAllowed: isAllowed,
				})
			}
		} else {
			log.Fatal(i, tokens[i/2])
		}
	}
	return
}

func getAddr(result multicall.Multicall2Result) common.Address {
	cmABI := core.GetAbi("CreditManagerv2")
	if result.Success {
		values, err := cmABI.Unpack("pool", result.ReturnData)
		log.CheckFatal(err)
		return values[0].(common.Address)
	} else {
		panic("")
	}
}

func getBigInt(result multicall.Multicall2Result) *big.Int {
	facadeABI := core.GetAbi("CreditManagerv2")
	if result.Success {
		values, err := facadeABI.Unpack("collateralTokensCount", result.ReturnData)
		log.CheckFatal(err)
		return values[0].(*big.Int)
	} else {
		panic("")
	}
}

func get2BigInt(result multicall.Multicall2Result) (*big.Int, *big.Int) {
	facadeABI := core.GetAbi("CreditManagerv2")
	if result.Success {
		values, err := facadeABI.Unpack("calcCreditAccountAccruedInterest", result.ReturnData)
		log.CheckFatal(err)
		return values[0].(*big.Int), values[1].(*big.Int)
	} else {
		panic("")
	}
}
