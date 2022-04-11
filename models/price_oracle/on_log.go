package price_oracle

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/priceFeed"
	"github.com/Gearbox-protocol/sdk-go/artifacts/yearnPriceFeed"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	ChainlinkPriceFeed = iota
	YearnPriceFeed
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
		if priceFeedType == ChainlinkPriceFeed {
			obj := chainlink_price_feed.NewChainlinkPriceFeed(token, oracle, oracle, blockNum, mdl.SyncAdapter.Client, mdl.Repo, version)
			mdl.Repo.AddSyncAdapter(obj)
		} else if priceFeedType == YearnPriceFeed {
			mdl.Repo.AddYearnFeed(token, oracle, blockNum, version)
			// obj := aggregated_block_feed.NewYearnPriceFeed(token, oracle, blockNum, mdl.SyncAdapter.Client, mdl.Repo, version)
			// mdl.Repo.AddSyncAdapter(obj)
		} else {
			log.Fatal("Unknown PriceFeed type", priceFeedType)
		}
	}
}

func (mdl *PriceOracle) checkPriceFeedContract(discoveredAt int64, oracle string) (int, error) {
	pfContract, err := priceFeed.NewPriceFeed(common.HexToAddress(oracle), mdl.Client)
	if err != nil {
		return -1, err
	}
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(discoveredAt),
	}
	_, err = pfContract.PhaseId(opts)
	if err != nil {
		if utils.Contains([]string{"VM execution error.", "execution reverted"}, err.Error()) {
			yearnContract, err := yearnPriceFeed.NewYearnPriceFeed(common.HexToAddress(oracle), mdl.Client)
			if err != nil {
				return -1, err
			}
			_, err = yearnContract.YVault(opts)
			if err != nil {
				return -1, fmt.Errorf("Neither chainlink nor yearn price feed %s", err)
			}
			return YearnPriceFeed, nil
		}
	} else {
		return ChainlinkPriceFeed, nil
	}
	return -1, fmt.Errorf("PriceFeed type not found")
}
