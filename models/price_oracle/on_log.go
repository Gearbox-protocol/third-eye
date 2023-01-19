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
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

func (mdl *PriceOracle) OnLog(txLog types.Log) {
	blockNum := int64(txLog.BlockNumber)
	switch txLog.Topics[0] {
	case core.Topic("NewPriceFeed(address,address)"):
		newPriceFeedEvent, err := mdl.contractETH.ParseNewPriceFeed(txLog)
		if err != nil {
			log.Fatal("[PriceOracle]: Cant unpack NewPriceFeed event", err)
		}
		mdl.Repo.AddDAOOperation(&schemas.DAOOperation{
			BlockNumber: blockNum,
			LogID:       txLog.Index,
			TxHash:      txLog.TxHash.Hex(),
			Contract:    mdl.Address,
			Type:        schemas.NewPriceFeed,
			Args: &core.Json{
				"priceFeed": newPriceFeedEvent.PriceFeed.Hex(),
				"token":     newPriceFeedEvent.Token.Hex(),
			},
		})

		token := newPriceFeedEvent.Token.Hex()
		oracle := newPriceFeedEvent.PriceFeed.Hex()
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
		if utils.Contains([]string{"VM execution error.", "execution reverted"}, err.Error()) {
			yearnContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(oracle), mdl.Client)
			if err != nil {
				return ds.UnknownPF, false, err
			}
			_, err = yearnContract.YVault(opts)
			if err != nil {
				description, err := yearnContract.Description(opts)
				if strings.Contains(description, "USD Composite") {
					return ds.CompositeChainlinkPF, false, nil
				} else if strings.Contains(description, "CurveLP pricefeed") {
					return ds.CurvePF, false, nil
				} else if strings.Contains(description, "Wrapped liquid staked Ether 2.0") { // steth price feed will behandled like YearnPF
					return ds.YearnPF, false, nil
				} else if strings.Contains(description, "Bounded") {
					return ds.ChainlinkPriceFeed, true, nil
				} else if strings.Contains(description, "Zero pricefeed") {
					return ds.ZeroPF, false, nil
				} else if strings.Contains(description, "ZERO (one) priceFeed") {
					return ds.AlmostZeroPF, false, nil
				} else {
					log.Info(description, oracle)
					return ds.UnknownPF, false, fmt.Errorf("neither chainlink nor yearn nor curve price feed %v, got %s", err, description)
				}
			}
			return ds.YearnPF, false, nil
		}
	} else {
		description, err := core.CallFuncWithExtraBytes(mdl.Client, "7284e416", common.HexToAddress(oracle), discoveredAt, nil) // description()
		return ds.ChainlinkPriceFeed,
			(err == nil) && strings.Contains(string(description), "Bounded"),
			nil
	}
	return ds.UnknownPF, false, fmt.Errorf("PriceFeed type not found")
}
