package repository

import (
	"fmt"
	"math"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceOracle"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// return the active first oracle under blockNum
// if all disabled return the last one
func (repo *Repository) GetActivePriceOracleByBlockNum(blockNum int64) (string, error) {
	var disabledLastOracle, activeFirstOracle string
	var disabledOracleBlock, activeOracleBlock int64
	activeOracleBlock = math.MaxInt64
	oracles := repo.GetKit().GetAdapterAddressByName(ds.PriceOracle)
	for _, addr := range oracles {
		oracleAdapter := repo.GetAdapter(addr)
		if oracleAdapter.GetDiscoveredAt() <= blockNum {
			if oracleAdapter.IsDisabled() {
				if disabledOracleBlock < oracleAdapter.GetDiscoveredAt() {
					disabledOracleBlock = oracleAdapter.GetDiscoveredAt()
					disabledLastOracle = addr
				}
			} else {
				if activeOracleBlock > oracleAdapter.GetDiscoveredAt() {
					activeOracleBlock = oracleAdapter.GetDiscoveredAt()
					activeFirstOracle = addr
				}
			}
		}
	}
	if activeFirstOracle != "" {
		return activeFirstOracle, nil
	} else if disabledLastOracle != "" {
		return disabledLastOracle, nil
	} else {
		return "", fmt.Errorf("Not Found")
	}
}

func (repo *Repository) GetPriceOracleByVersion(version int16) (string, error) {
	addrProviderAddr := repo.GetKit().GetAdapterAddressByName(ds.AddressProvider)
	addrProvider := repo.GetKit().GetAdapter(addrProviderAddr[0])
	details := addrProvider.GetDetails()
	if details != nil {
		priceOracles, ok := details["priceOracles"].(map[string]interface{})
		if ok {
			value, ok := priceOracles[fmt.Sprintf("%d", version)].(string)
			if ok {
				return value, nil
			}
		}
	}
	return "", fmt.Errorf("Not Found")
}

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
