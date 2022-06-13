package repository

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// This function is used for getting the collateral value in usd and underlying
func (repo *Repository) GetValueInCurrency(blockNum int64, version int16, token, currency string, amount *big.Int) *big.Int {
	oracle, err := repo.GetPriceOracleByVersion(version)
	if err != nil {
		log.Fatalf("err %s version: %d", err, version)
	}
	poContract, err := priceOracle.NewPriceOracle(common.HexToAddress(oracle), repo.client)
	log.CheckFatal(err)
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	currencyAddr := common.HexToAddress(repo.GetUSDCAddr())
	if currency != "USDC" {
		currencyAddr = common.HexToAddress(currency)
	}
	usdcAmount, err := poContract.Convert(opts, amount, common.HexToAddress(token), currencyAddr)
	log.CheckFatal(err)
	// convert to 8 decimals
	return usdcAmount
}
