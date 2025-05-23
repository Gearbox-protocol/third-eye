package repository

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOraclev2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// This function is used for getting the collateral value in usd and underlying
func (repo *Repository) GetValueInCurrency(blockNum int64, version core.VersionType, token, currency string, amount *big.Int) *big.Int {
	oracle, _, err := repo.GetActivePriceOracleByBlockNum(blockNum)
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
	if currencyAddr.Hex() == token {
		return amount
	}
	sig := big.NewInt(1)
	if amount.Cmp(big.NewInt(0)) < 0 {
		amount = new(big.Int).Neg(amount)
		sig = big.NewInt(-1)
	}

	if version.IsGBv1() {
		oracle = "0x0e74a08443c5E39108520589176Ac12EF65AB080"
		poContract, err := priceOracle.NewPriceOracle(common.HexToAddress(string(oracle)), repo.client)
		log.CheckFatal(err)
		usdcAmount, err := poContract.Convert(opts, amount, common.HexToAddress(token), currencyAddr)
		if err != nil {
			log.Fatalf("%v %s %d %s %s at block %d", err, oracle, amount, token, currencyAddr, blockNum)
		}
		return new(big.Int).Mul(usdcAmount, sig)
	} else { // v2 and above
		poContract, err := priceOraclev2.NewPriceOraclev2(oracle.Hex(), repo.client)
		log.CheckFatal(err)
		usdcAmount, err := poContract.Convert(opts, amount, common.HexToAddress(token), currencyAddr)
		if err != nil {
			{ // one token is redstone token
				tokenPrice := repo.GetRedStonePrice(blockNum, oracle, token)
				currencyPrice := repo.GetRedStonePrice(blockNum, oracle, currency)
				if tokenPrice != nil || currencyPrice != nil {
					if tokenPrice == nil {
						tokenPrice, err = poContract.GetPrice(opts, common.HexToAddress(token))
						if err != nil {
							log.Fatalf("%v %s %d %s %s at block %d", err, oracle, amount, token, currencyAddr, blockNum)
						}
					}
					if currencyPrice == nil {
						currencyPrice, err = poContract.GetPrice(opts, currencyAddr)
						if err != nil {
							log.Fatalf("%v %s %d %s %s at block %d", err, oracle, amount, token, currencyAddr, blockNum)
						}
					}
					// amount *token Price *currenyDecimals/ currencyPrice * tokenDecimals
					num := utils.GetInt64(new(big.Int).Mul(amount, tokenPrice), -repo.GetToken(currencyAddr.Hex()).Decimals)
					denom := utils.GetInt64(currencyPrice, -repo.GetToken(token).Decimals)
					// log.Fatal(new(big.Int).Quo(num, denom), token, currencyAddr, tokenPrice, currencyPrice, amount)
					return new(big.Int).Quo(num, denom)
				}
			}
			log.Fatalf("%v %s %d %s %s at block %d", err, oracle, amount, token, currencyAddr, blockNum)
		}
		return new(big.Int).Mul(usdcAmount, sig)
	}
}
