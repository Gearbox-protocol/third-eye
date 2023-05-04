package repository

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOraclev2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// This function is used for getting the collateral value in usd and underlying
func (repo *Repository) GetValueInCurrency(blockNum int64, version core.VersionType, token, currency string, amount *big.Int) *big.Int {
	oracle, err := repo.GetPriceOracleByDiscoveredAt(blockNum)
	if err != nil {
		log.Fatalf("err %s version: %d", err, version)
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	currencyAddr := common.HexToAddress(repo.GetUSDCAddr())
	if currency != "USDC" {
		currencyAddr = common.HexToAddress(currency)
	}
	if version.IsGBv1() {
		poContract, err := priceOracle.NewPriceOracle(common.HexToAddress(oracle), repo.client)
		log.CheckFatal(err)
		usdcAmount, err := poContract.Convert(opts, amount, common.HexToAddress(token), currencyAddr)
		if err != nil {
			log.Fatalf("%v %s %d %s %s", err, oracle, amount, token, currencyAddr)
		}
		return usdcAmount
	} else { // v2 and above
		poContract, err := priceOraclev2.NewPriceOraclev2(common.HexToAddress(oracle), repo.client)
		log.CheckFatal(err)
		usdcAmount, err := poContract.Convert(opts, amount, common.HexToAddress(token), currencyAddr)
		if err != nil {
			log.Fatalf("%v %s %d %s %s", err, oracle, amount, token, currencyAddr)
		}
		return usdcAmount
	}
}
