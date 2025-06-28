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
		t = core.GetToken(chainid, "USDC.e")
	} else if chainid == 56 {
		t = core.GetToken(chainid, "USDT")
	} else {
		t = core.GetToken(chainid, "USDC")
	}
	return t
}

// This function is used for getting the collateral value in usd and underlying
func (repo *Repository) GetValueInCurrency(blockNum int64, pool, token, currency string, amount *big.Int) (*big.Int, float64) {
	oracle, version := repo.GetPoolToPriceOraclev3(blockNum, pool)
	// oracle, _, err := repo.GetActivePriceOracleByBlockNum(blockNum)
	// if err != nil {
	// 	log.Fatalf("err %s version: %d", err, version)
	// }
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	currencyAddr := repo.GetUSD()
	if currency != "USD" {
		currencyAddr = common.HexToAddress(currency)
	}
	currencyDecimals := repo.GetToken(currencyAddr.Hex()).Decimals
	if currencyAddr.Hex() == token {
		return amount, utils.GetFloat64Decimal(amount, currencyDecimals)
	}
	sig := big.NewInt(1)
	if amount.Cmp(big.NewInt(0)) < 0 {
		amount = new(big.Int).Neg(amount)
		sig = big.NewInt(-1)
	}

	if version.IsGBv1() {
		// oracle = "0x0e74a08443c5E39108520589176Ac12EF65AB080" // already from getPooltopriceoraclev3
		poContract, err := priceOracle.NewPriceOracle(common.HexToAddress(string(oracle)), repo.client)
		log.CheckFatal(err)
		usdcAmount, err := poContract.Convert(opts, amount, common.HexToAddress(token), currencyAddr)
		if err != nil {
			log.Fatalf("%v %s %d %s %s at block %d", err, oracle, amount, token, currencyAddr, blockNum)
		}
		amount := new(big.Int).Mul(usdcAmount, sig)
		return amount, utils.GetFloat64Decimal(amount, currencyDecimals) // for v1 on mainnet
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
					_amount := new(big.Int).Quo(num, denom)
					return _amount, utils.GetFloat64Decimal(_amount, currencyDecimals)
				}
			}
			log.Fatalf("%v %s %d %s %s at block %d", err, oracle, amount, token, currencyAddr, blockNum)
		}
		_amount := new(big.Int).Mul(usdcAmount, sig)
		return _amount, utils.GetFloat64Decimal(_amount, currencyDecimals)
	}
}
