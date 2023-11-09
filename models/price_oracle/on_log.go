package price_oracle

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnPriceFeed"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"gorm.io/gorm/utils"
)

func (mdl *PriceOracle) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("NewPriceFeed(address,address)"), core.Topic("SetPriceFeed(address,address,uint32,bool,bool)"):
		token := common.BytesToAddress(txLog.Topics[1].Bytes()).Hex()  // token
		oracle := common.BytesToAddress(txLog.Topics[2].Bytes()).Hex() // priceFeed

		//
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.NewPriceFeed,
			Args: &core.Json{
				"priceFeed": oracle,
				"token":     token,
			},
		})

		version := mdl.GetVersion()
		priceFeedType, bounded, err := mdl.checkPriceFeedContract(blockNum, oracle)
		if err != nil {
			log.Fatalf("Oracle %s, err: %s", oracle, err)
		}
		switch priceFeedType {
		// almost zero price feed is for blocker token on credit account
		case ds.YearnPF, ds.CurvePF, ds.ChainlinkPriceFeed, ds.ZeroPF, ds.AlmostZeroPF, ds.CompositeChainlinkPF:
			// four types of oracles
			// - Zero or almost zero price feed: constant price value
			// - Chainlink price feed: market based price value
			// - Composite price feed: price calculated from multiple price feeds
			// - Query price feed: price fetched from curve or yearn
			mdl.Repo.GetToken(token)
			mdl.Repo.AddNewPriceOracleEvent(&schemas.TokenOracle{
				Token:       token,
				Oracle:      oracle,
				Feed:        oracle, // feed is same as oracle
				BlockNumber: blockNum,
				Version:     version,
				FeedType:    priceFeedType,
			}, bounded)
		default:
			log.Fatal("Unknown PriceFeed type", priceFeedType)
		}
	}
}

// YearnPF covers LIDO, AAVE, COMPOUND, YEARN, ERC4626, Balancer(Stable, weighted)
// CurvePF covers curve and convex
// ChainlinkPF cover chainlink
func (mdl *PriceOracle) checkPriceFeedContract(discoveredAt int64, oracle string) (string, bool, error) { // type, bounded , error
	if oracle == "0xE26FB07da646138553f635c94E2a345270240e30" { // for goerli , the chainlink bounded oracle doesn't have phaseId method // LUSD price oracle
		return ds.ChainlinkPriceFeed, true, nil
	}
	pfContract, err := priceFeed.NewPriceFeed(common.HexToAddress(oracle), mdl.Client)
	if err != nil {
		return ds.UnknownPF, false, err
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(discoveredAt),
	}
	_, err = pfContract.PhaseId(opts)
	if err != nil {
		if strings.Contains(err.Error(), "VM execution error.") ||
			strings.Contains(err.Error(), "execution reverted") {
			yearnContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(oracle), mdl.Client)
			log.CheckFatal(err)
			_, err = yearnContract.YVault(opts)
			if err != nil {
				description, err := yearnContract.Description(opts)
				if strings.Contains(description, "USD Composite") {
					// https://github.com/Gearbox-protocol/core-v2/blob/main/contracts/oracles/CompositePriceFeed.sol
					return ds.CompositeChainlinkPF, false, nil
				} else if strings.Contains(description, "CurveLP pricefeed") || utils.Contains([]string{
					"PRICEFEED_OHMFRAXBP",
					"PRICEFEED_MIM_3LP3CRV",
					"PRICEFEED_crvCRVETH",
					"PRICEFEED_crvCVXETH",
					"PRICEFEED_crvUSDTWBTCWETH",
					"PRICEFEED_LDOETH",
					"PRICEFEED_crvUSDETHCRV",
					"crvPlain3andSUSD price feed",
				}, description) {
					// https://github.com/Gearbox-protocol/integrations-v2/tree/main/contracts/oracles/curve
					return ds.CurvePF, false, nil
				} else if strings.Contains(description, "Wrapped liquid staked Ether 2.0") { // steth price feed will behandled like YearnPF
					//https://github.com/Gearbox-protocol/integrations-v2/blob/main/contracts/oracles/lido/WstETHPriceFeed.sol
					return ds.YearnPF, false, nil
				} else if strings.Contains(description, "Bounded") {
					// https://github.com/Gearbox-protocol/core-v2/blob/main/contracts/oracles/BoundedPriceFeed.sol
					return ds.ChainlinkPriceFeed, true, nil
				} else if strings.Contains(description, "Zero pricefeed") {
					// zero for G-OBS
					// https://github.com/Gearbox-protocol/core-v2/blob/main/contracts/oracles/ZeroPriceFeed.sol
					return ds.ZeroPF, false, nil
				} else if strings.Contains(description, "ZERO (one) priceFeed") {
					// deprecated not used.
					return ds.AlmostZeroPF, false, nil
				} else {
					log.Info(description, oracle)
					return ds.UnknownPF, false, fmt.Errorf("neither chainlink nor yearn nor curve price feed %v, got %s", err, description)
				}
			}
			return ds.YearnPF, false, nil
		}
	} else { //chainlink description
		yearnContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(oracle), mdl.Client)
		log.CheckFatal(err)
		description, err := yearnContract.Description(opts)
		return ds.ChainlinkPriceFeed,
			(err == nil) && strings.Contains(string(description), "Bounded"),
			nil
	}
	return ds.UnknownPF, false, fmt.Errorf("PriceFeed type not found")
}
