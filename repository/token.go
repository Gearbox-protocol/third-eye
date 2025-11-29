package repository

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOraclev2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (repo *Repository) GetUSD() common.Address {
	chainid := core.GetChainId(repo.client)
	if chainid == 1337 {
		for _, addr := range repo.GetTokens() {
			token := repo.GetToken(addr)
			if token.Symbol == "USDC" {
				return common.HexToAddress(token.Address)
			}
		}
	}
	chainid = core.GetBaseChainId(repo.client)
	var t common.Address
	if chainid == 146 {
		t = core.GetToken(chainid, "USDC_e")
	} else if chainid == 1135 || chainid == 43111 {
		t = core.GetToken(chainid, "USDC.e")
	} else if chainid == 56 {
		t = core.GetToken(chainid, "USDT")
	} else if chainid == 9745 {
		t = core.GetToken(chainid, "USDT0")
	} else {
		t = core.GetToken(chainid, "USDC")
	}
	return t
}

// This function is used for getting the collateral value in usd and underlying
func (repo *Repository) GetValueInCurrency(blockNum int64, pool string, version core.VersionType, token, currency string, amount *big.Int) float64 {
	var neg bool
	if amount.Cmp(big.NewInt(0)) < 0 {
		amount = new(big.Int).Neg(amount)
		neg = true
	}
	oracle, pfVersion := repo.GetPoolToPriceOraclev3(blockNum, pool)
	if core.GetChainId(repo.client) == 1337 {
		pfVersion = version
	}
	usdAmount := repo.getValueInCurrency(blockNum, oracle, pfVersion, token, currency, amount)
	if neg {
		usdAmount = -usdAmount
	}
	return usdAmount
}
func (repo *Repository) getValueInCurrency(blockNum int64, oracle schemas.PriceOracleT, version core.VersionType, token, currency string, amount *big.Int) float64 {

	// oracle, _, err := repo.GetActivePriceOracleByBlockNum(blockNum)
	// if err != nil {
	// 	log.Fatalf("err %s version: %d", err, version)
	// }
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	currencyAddr := repo.GetUSD()
	if oracle.Hex().Hex() == "0x6385892aCB085eaa24b745a712C9e682d80FF681" && token == "0x514910771AF9Ca656af840dff83E8264EcF986CA" &&
		(currency == "USD" || currency == "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48") { // LINK oracle missing
		return 0 // link token price and amount for usd conversion
	}
	if currency != "USD" {
		currencyAddr = common.HexToAddress(currency)
	}
	currencyDecimals := repo.GetToken(currencyAddr.Hex()).Decimals
	if currencyAddr.Hex() == token {
		// return amount, utils.GetFloat64Decimal(amount, currencyDecimals)
		return utils.GetFloat64Decimal(amount, currencyDecimals)
	}

	if version.IsGBv1() {
		// oracle = "0x0e74a08443c5E39108520589176Ac12EF65AB080" // already from getPooltopriceoraclev3
		poContract, err := priceOracle.NewPriceOracle(common.HexToAddress(string(oracle)), repo.client)
		log.CheckFatal(err)
		usdcAmount, err := poContract.Convert(opts, amount, common.HexToAddress(token), currencyAddr)
		if err != nil {
			log.Fatalf("%v %s %d %s %s at block %d", err, oracle, amount, token, currencyAddr, blockNum)
		}
		return utils.GetFloat64Decimal(usdcAmount, currencyDecimals) // for v1 on mainnet
	} else if currency != "USD" { // v2 and above
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
							log.Fatalf("%v %s %d %s %s at block %d. %s", err, oracle, amount, token, currencyAddr, blockNum, currency)
						}
					}
					if currencyPrice == nil {
						currencyPrice, err = poContract.GetPrice(opts, currencyAddr)
						if err != nil {
							log.Fatalf("%v %s %d %s %s at block %d. %s", err, oracle, amount, token, currencyAddr, blockNum, currency)
						}
					}
					// amount *token Price *currenyDecimals/ currencyPrice * tokenDecimals
					num := utils.GetInt64(new(big.Int).Mul(amount, tokenPrice), -repo.GetToken(currencyAddr.Hex()).Decimals)
					denom := utils.GetInt64(currencyPrice, -repo.GetToken(token).Decimals)
					// log.Fatal(new(big.Int).Quo(num, denom), token, currencyAddr, tokenPrice, currencyPrice, amount)
					_amount := new(big.Int).Quo(num, denom)
					return utils.GetFloat64Decimal(_amount, currencyDecimals)
				}
			}
			log.Fatalf("%v %s %d %s %s at block %d", err, oracle, amount, token, currencyAddr, blockNum)
		}
		// _amount := new(big.Int).Mul(usdcAmount)
		return utils.GetFloat64Decimal(usdcAmount, currencyDecimals)
	} else {
		poContract, err := priceOraclev2.NewPriceOraclev2(oracle.Hex(), repo.client)
		log.CheckFatal(err)
		price, err := poContract.GetPrice(opts, common.HexToAddress(token))
		if err != nil {
			price = repo.GetRedStonePrice(blockNum, oracle, token)
			if price == nil {
				log.Fatalf("err:%v oracle:%s amount:%d token:%s currencyAddr:%s at block %d", err, oracle, amount, token, currencyAddr, blockNum)
			}
		}
		amountInCurrency := utils.GetInt64(new(big.Int).Mul(amount, price), repo.GetToken(token).Decimals)
		// return amountInCurrency, utils.GetFloat64Decimal(amountInCurrency, 8)
		return utils.GetFloat64Decimal(amountInCurrency, 8)
	}
}
