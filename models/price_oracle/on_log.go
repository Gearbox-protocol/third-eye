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
		priceFeedType, err := mdl.checkPriceFeedContract(blockNum, oracle)
		if err != nil {
			log.Fatal(err)
		}
		switch priceFeedType {
		case ds.YearnPF, ds.CurvePF, ds.ChainlinkPriceFeed, ds.ZeroPF, ds.AlmostZeroPF:
			if oracle == "0xEB24b7c2fB6497f28c937942439B4EAAE9535525" {
				log.Info(token, oracle, blockNum, txLog.Index)
			}
			mdl.Repo.AddNewPriceOracleEvent(&schemas.TokenOracle{
				Token:       token,
				Oracle:      oracle,
				Feed:        oracle, // feed is same as oracle
				BlockNumber: blockNum,
				Version:     version,
				FeedType:    priceFeedType,
			})
		default:
			log.Fatal("Unknown PriceFeed type", priceFeedType)
		}
	}
}

func (mdl *PriceOracle) checkPriceFeedContract(discoveredAt int64, oracle string) (string, error) {
	pfContract, err := priceFeed.NewPriceFeed(common.HexToAddress(oracle), mdl.Client)
	if err != nil {
		return ds.UnknownPF, err
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(discoveredAt),
	}
	_, err = pfContract.PhaseId(opts)
	if err != nil {
		if utils.Contains([]string{"VM execution error.", "execution reverted"}, err.Error()) {
			yearnContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(oracle), mdl.Client)
			if err != nil {
				return ds.UnknownPF, err
			}
			_, err = yearnContract.YVault(opts)
			if err != nil {
				description, err := yearnContract.Description(opts)
				if strings.Contains(description, "CurveLP pricefeed") {
					return ds.CurvePF, nil
				} else if strings.Contains(description, "Zero pricefeed") {
					return ds.ZeroPF, nil
				} else if strings.Contains(description, "ZERO (one) priceFeed") {
					return ds.AlmostZeroPF, nil
				} else {
					log.Info(description, oracle)
					return ds.UnknownPF, fmt.Errorf("neither chainlink nor yearn nor curve price feed %v, got %s", err, description)
				}
			}
			return ds.YearnPF, nil
		}
	} else {
		return ds.ChainlinkPriceFeed, nil
	}
	return ds.UnknownPF, fmt.Errorf("PriceFeed type not found")
}
