package dc_wrapper

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/mainnet"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type MainnetDC struct {
	dc                    *mainnet.DataCompressor
	creditManagerToFilter map[common.Address]common.Address
	client                core.ClientI
}

func NewMainnetDC(client core.ClientI) *MainnetDC {

	return &MainnetDC{
		client:                client,
		creditManagerToFilter: make(map[common.Address]common.Address),
	}
}

func (m *MainnetDC) SetAddr(addr common.Address) {
	dc, err := mainnet.NewDataCompressor(addr, m.client)
	log.CheckFatal(err)
	m.dc = dc
}

func (repo *MainnetDC) AddCreditManagerToFilter(cmAddr, cfAddr string) {
	repo.creditManagerToFilter[common.HexToAddress(cmAddr)] = common.HexToAddress(cfAddr)
}

func getMask(client core.ClientI, blockNum int64, cfAddr, accountAddr common.Address) *big.Int {
	// abi := core.GetAbi("CreditFilter")
	// data, err := abi.Pack("enabledTokens", accountAddr)
	// log.CheckFatal(err)
	// calls := []multicall.Multicall2Call{{
	// 	Target:   cfAddr,
	// 	CallData: data,
	// }}
	// results := core.MakeMultiCall(client, blockNum, false, calls)
	// if !results[0].Success {
	// 	log.Fatalf("Getting mask for addr(%s) on cm(%s) failed", accountAddr, cfAddr)
	// }
	// return new(big.Int).SetBytes(results[0].ReturnData)
	data := make([]byte, 4+32)
	sig, err := hex.DecodeString("b451cecc") // enabledTokens
	log.CheckFatal(err)
	copy(data, sig)
	for i, b := range accountAddr.Bytes() {
		data[4+12+i] = b
	}
	//
	data = append(data)
	msg := ethereum.CallMsg{
		To:   &cfAddr,
		Data: data,
	}
	returnData, err := client.CallContract(context.TODO(), msg, big.NewInt(blockNum))
	log.CheckFatal(err)
	return new(big.Int).SetBytes(returnData)

	//
	// case 2:
	// cm, err := creditManagerv2.NewCreditManagerv2(common.HexToAddress(cmAddr), repo.client)
	// log.CheckFatal(err)
	// mask, err := cm.EnabledTokensMap(opts, common.HexToAddress(accountAddr))
	// log.CheckFatal(err)
	// return mask
}

func (dc *MainnetDC) GetPoolData(opts *bind.CallOpts, _pool common.Address) (dcv2.PoolData, error) {
	data, err := dc.dc.GetPoolData(opts, _pool)
	log.CheckFatal(err)
	latestFormat := dcv2.PoolData{
		Addr:                   data.Addr,
		IsWETH:                 data.IsWETH,
		Underlying:             data.UnderlyingToken,
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
		Version:                1,
	}
	return latestFormat, nil
}

func (mainnetDC *MainnetDC) GetCreditManagerData(opts *bind.CallOpts,
	_creditManager common.Address, borrower common.Address) (dcv2.CreditManagerData, error) {
	data, err := mainnetDC.dc.GetCreditManagerData(opts, _creditManager, borrower)
	log.CheckFatal(err)
	latestFormat := dcv2.CreditManagerData{
		Addr:               data.Addr,
		Underlying:         data.UnderlyingToken,
		IsWETH:             data.IsWETH,
		CanBorrow:          data.CanBorrow,
		BorrowRate:         data.BorrowRate,
		MinAmount:          data.MinAmount,
		MaxAmount:          data.MaxAmount,
		MaxLeverageFactor:  data.MaxLeverageFactor,
		AvailableLiquidity: data.AvailableLiquidity,
		CollateralTokens:   data.AllowedTokens,
	}
	// for _, adapter := range data.Adapters {
	// 	latestFormat.Adapters = append(latestFormat.Adapters, dcv2.ContractAdapter{
	// 		Adapter:         adapter.Adapter,
	// 		AllowedContract: adapter.AllowedContract,
	// 	})
	// }
	return latestFormat, nil
}

func (obj *MainnetDC) GetCreditAccountData(opts *bind.CallOpts,
	creditManager common.Address, borrower common.Address) (dcv2.CreditAccountData, error) {
	data, err := obj.dc.GetCreditAccountDataExtended(opts, creditManager, borrower)
	log.CheckFatal(err)
	latestFormat := dcv2.CreditAccountData{
		Addr:                       data.Addr,
		Borrower:                   data.Borrower,
		InUse:                      data.InUse,
		CreditManager:              data.CreditManager,
		Underlying:                 data.UnderlyingToken,
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
	//
	var blockNum int64
	if opts != nil {
		blockNum = opts.BlockNumber.Int64()
	}
	cfAddr := obj.creditManagerToFilter[latestFormat.CreditManager]
	mask := getMask(obj.client, blockNum, cfAddr, latestFormat.Addr)
	latestFormat.Balances = convertTodcv2Balance(data.Balances, mask)
	//
	return latestFormat, nil
}

func convertTodcv2Balance(balances []mainnet.DataTypesTokenBalance, mask *big.Int) (dcv2Balances []dcv2.TokenBalance) {
	maskInBits := fmt.Sprintf("%b", mask)
	maskLen := len(maskInBits)
	for i, balance := range balances {
		var isEnabled bool
		if maskLen > i {
			isEnabled = maskInBits[maskLen-i-1] == '1'
		}
		dcv2Balances = append(dcv2Balances, dcv2.TokenBalance{
			Token:     balance.Token,
			Balance:   balance.Balance,
			IsAllowed: balance.IsAllowed,
			IsEnabled: isEnabled,
		})
	}
	return
}
