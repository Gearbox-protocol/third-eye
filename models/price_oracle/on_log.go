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

// QueryPriceFeed stores in details reserve status via PFVersion in details.Tokens.pfversion
// chainlinkPriceFeed and compositeChainlinkPriceFeed stores pfversion as reserve status in details

func (mdl *PriceOracle) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("NewPriceFeed(address,address)"),
		core.Topic("SetPriceFeed(address,address,uint32,bool,bool)"),
		core.Topic("SetReservePriceFeed(address,address,uint32,bool)"):
		//
		token := common.BytesToAddress(txLog.Topics[1].Bytes()).Hex()  // token
		oracle := common.BytesToAddress(txLog.Topics[2].Bytes()).Hex() // priceFeed
		isReverse := core.Topic("SetReservePriceFeed(address,address,uint32,bool)") == txLog.Topics[0]
		// if isReverse {
		// 	log.Fatal("token", token, "oracle", oracle)
		// }
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
				"reserve":   isReverse,
			},
		})

		priceFeedType, bounded, err := mdl.checkPriceFeedContract(blockNum, oracle, token)
		if err != nil {
			log.Fatalf("Oracle %s, err: %s, blockNum %d", oracle, err, blockNum)
		}
		if priceFeedType == ds.RedStonePF {
			pfs := core.GetRedStonePFByChainId(core.GetChainId(mdl.Client))
			addrToSym := core.GetTokenToSymbolByChainId(core.GetChainId(mdl.Client))
			sym := addrToSym[common.HexToAddress(token)]
			_, ok := pfs[sym]
			if !ok {
				log.Warnf("RedStonePF not found in config for %s(%s). update sdk-go.", sym, token)
			}
		}
		switch priceFeedType {
		// almost zero price feed is for blocker token on credit account
		case ds.ChainlinkPriceFeed, ds.CompositeChainlinkPF,
			ds.ZeroPF, ds.AlmostZeroPF,
			ds.RedStonePF, ds.CompositeRedStonePF, ds.YearnPF, ds.SingleAssetPF, ds.CurvePF:
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
				Version:     mdl.GetVersion(),
				Reserve:     isReverse,
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
func (mdl *PriceOracle) checkPriceFeedContract(discoveredAt int64, oracle, token string) (string, bool, error) { // type, bounded , error
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
	_, err = pfContract.PhaseId(opts) // only on chainlink
	if err != nil {
		if strings.Contains(err.Error(), "VM execution error.") ||
			strings.Contains(err.Error(), "Required data unavailable") ||
			strings.Contains(err.Error(), "execution reverted") {
			if mdl.GetVersion().MoreThanEq(core.NewVersion(300)) {
				return mdl.v3PriceFeedType(opts, oracle, token)
			} else {
				return mdl.v2PriceFeedType(opts, oracle)
			}
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

// https://github.com/Gearbox-protocol/integrations-v2/tree/faa9cfd4921c62165782dcdc196ff5a0c0e6075d/contracts/oracles
// https://github.com/Gearbox-protocol/oracles-v3/tree/2ac6d1ba1108df949222084791699d821096bc8c/contracts/oracles
func (mdl *PriceOracle) v3PriceFeedType(opts *bind.CallOpts, oracle, token string) (string, bool, error) {
	data, err := core.CallFuncWithExtraBytes(mdl.Client, "3fd0875f", common.HexToAddress(oracle), 0, nil) // priceFeedType
	if err != nil {
		{ // redstone feed without price on demand doesn't have priceFeedType method.
			//https://etherscan.io/address/0xbC5FBcf58CeAEa19D523aBc76515b9AEFb5cfd58#readProxyContract
			con, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(oracle), mdl.Client)
			log.CheckFatal(err)
			if description, err := con.Description(nil); err != nil {
				log.Fatal(oracle, token, "priceFeedType failed: ", err)
			} else {
				description = strings.ToLower(string(description))
				if strings.Contains(description, "redstone") {
					return ds.CurvePF, false, nil
				} else {
					log.Fatal(oracle, token, "priceFeedType failed: ", description, err)
				}
			}
		}
	}
	pfType := new(big.Int).SetBytes(data).Int64()
	switch pfType {
	case core.V3_COMPOSITE_ORACLE:
		{ // composite feed is using redstone feed
			// https://etherscan.io/address/0x8751F736E94F6CD167e8C5B97E245680FbD9CC36#readProxyContract
			// this composite feed doesn't have known internal chainlink feeds.
			// third-eye queries price feeds in 2 ways, for chainlink and compsite feeds it filters logs.
			// For other , it fetches prices periodic.
			// The new composite has internal feeds that don’t have known chainlink abi sigs. So, composite feed adapter is failing.
			// I will treat them as query feed to be periodic synced every 10-15 blocks.

			yearnContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(oracle), mdl.Client)
			log.CheckFatal(err)
			description, err := yearnContract.Description(opts)
			log.CheckFatal(err)
			if strings.Contains(strings.ToLower(description), "redstone") {
				return ds.CompositeRedStonePF, false, nil
			}
		}
		return ds.CompositeChainlinkPF, false, nil
	case core.V3_YEARN_ORACLE:
		return ds.YearnPF, false, nil
	case core.V3_CHAINLINK_ORACLE:
		return ds.ChainlinkPriceFeed, true, nil
	case core.V3_CURVE_USD_ORACLE, core.V3_CURVE_CRYPTO_ORACLE, // usd and crypto
		core.V3_CURVE_2LP_ORACLE, core.V3_CURVE_3LP_ORACLE, core.V3_CURVE_4LP_ORACLE: // 2lp,3lp, 4lp
		return ds.CurvePF, false, nil
	case core.V3_ZERO_ORACLE:
		return ds.ZeroPF, false, nil
		// SingleAssetLPPriceFeed
	case core.V3_WSTETH_ORACLE, core.V3_WRAPPED_AAVE_V2_ORACLE, // lido, aave,
		core.V3_BALANCER_STABLE_LP_ORACLE, core.V3_BALANCER_WEIGHTED_LP_ORACLE, // balancer
		core.V3_COMPOUND_V2_ORACLE,   // compounder
		core.V3_ERC4626_VAULT_ORACLE: // erc4626
		return ds.SingleAssetPF, false, nil
	case core.V3_REDSTONE_ORACLE:
		return ds.RedStonePF, false, nil
	default:
		yearnContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(oracle), mdl.Client)
		log.CheckFatal(err)
		description, err := yearnContract.Description(opts)
		log.CheckFatal(err)
		return ds.UnknownPF, false, fmt.Errorf("unknown v3 pfType %v, oracle: %s token: %s, description: %s", pfType, oracle, token, description)
	}
}

func (mdl *PriceOracle) v2PriceFeedType(opts *bind.CallOpts, oracle string) (string, bool, error) {
	yearnContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(oracle), mdl.Client)
	log.CheckFatal(err)
	_, err = yearnContract.YVault(opts)
	if err != nil {
		description, err := yearnContract.Description(opts)
		log.Infof("Add %s with desc: %s", oracle, description)
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
			return ds.SingleAssetPF, false, nil
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
			return ds.UnknownPF, false, fmt.Errorf("neither chainlink nor yearn nor curve price feed %v, got %s", err, description)
		}
	}
	return ds.YearnPF, false, nil
}
