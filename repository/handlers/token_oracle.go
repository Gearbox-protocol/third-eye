package handlers

import (
	"math/big"
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"github.com/Gearbox-protocol/third-eye/models/composite_chainlink"
	"gorm.io/gorm"
)

type TokenOracleRepo struct {
	// version  to token to oracle
	tokensCurrentOracle map[schemas.PFVersion]map[string]*schemas.TokenOracle // done
	mu                  *sync.Mutex
	adapters            *SyncAdaptersRepo
	blocks              *BlocksRepo
	repo                ds.RepositoryI
	client              core.ClientI
	zeroPFs             map[string]bool
}

func NewTokenOracleRepo(adapters *SyncAdaptersRepo, blocks *BlocksRepo, repo ds.RepositoryI, client core.ClientI) *TokenOracleRepo {
	return &TokenOracleRepo{
		tokensCurrentOracle: make(map[schemas.PFVersion]map[string]*schemas.TokenOracle),
		mu:                  &sync.Mutex{},
		adapters:            adapters,
		blocks:              blocks,
		repo:                repo,
		client:              client,
		zeroPFs:             map[string]bool{},
	}
}

// for price oracle/feeds
func (repo *TokenOracleRepo) LoadCurrentTokenOracle(db *gorm.DB) {
	defer utils.Elapsed("loadCurrentTokenOracle")()
	data := []*schemas.TokenOracle{}
	query := `SELECT distinct on (token, version, reserve) * FROM token_oracle order by token, version, reserve, block_num desc;`
	err := db.Raw(query).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, tokenOracle := range data {
		repo.addTokenCurrentOracle(tokenOracle)
	}
	repo.loadZeroPFs(db)
}

func (repo *TokenOracleRepo) loadZeroPFs(db *gorm.DB) {
	data := []schemas.TokenOracle{}
	err := db.Where("feed_type in ('ZeroPF', 'AlmostZeroPF')").Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, tokenOracle := range data {
		repo.zeroPFs[tokenOracle.Feed] = true
	}
}

func (repo *TokenOracleRepo) addTokenCurrentOracle(oracle *schemas.TokenOracle) {
	pfVersion := schemas.VersionToPFVersion(oracle.Version, oracle.Reserve)
	if repo.tokensCurrentOracle[pfVersion] == nil {
		repo.tokensCurrentOracle[pfVersion] = map[string]*schemas.TokenOracle{}
	}
	repo.tokensCurrentOracle[pfVersion][oracle.Token] = oracle
}

// if same feed is active for current token and version
func (repo *TokenOracleRepo) alreadyActiveFeedForToken(newTokenOracle *schemas.TokenOracle) bool {
	feedType := newTokenOracle.FeedType
	pfVersion := schemas.VersionToPFVersion(newTokenOracle.Version, newTokenOracle.Reserve)
	//
	if repo.tokensCurrentOracle[pfVersion] != nil &&
		repo.tokensCurrentOracle[pfVersion][newTokenOracle.Token] != nil {
		oldTokenOracle := repo.tokensCurrentOracle[pfVersion][newTokenOracle.Token]

		if oldTokenOracle.Feed == newTokenOracle.Feed {
			log.Debugf("Same %s(%s) added for token(%s)", feedType, newTokenOracle.Feed, newTokenOracle.Token)
			return true
		}
	}
	return false
}

func (repo *TokenOracleRepo) disablePrevAdapterAndAddNewTokenOracle(newTokenOracle *schemas.TokenOracle) {
	pfVersion := schemas.VersionToPFVersion(newTokenOracle.Version, newTokenOracle.Reserve)
	//
	if repo.tokensCurrentOracle[pfVersion] != nil &&
		repo.tokensCurrentOracle[pfVersion][newTokenOracle.Token] != nil {
		oldTokenOracle := repo.tokensCurrentOracle[pfVersion][newTokenOracle.Token]
		oldFeed := oldTokenOracle.Feed

		adapter := repo.adapters.GetAdapter(oldFeed)
		// disable the corresponding adapter
		if adapter == nil && repo.zeroPFs[oldFeed] {
			// no adapter is used for zeroPF as the price is always zero.
			// we can just work with 'adapter==nil' but we want to check if the adapter is null for other pricefeed by mistake. like disabled for chainlink etc.
		} else if adapter == nil {
			log.Error("Adapter not found for", oldFeed, utils.ToJson(oldTokenOracle))
		} else if adapter.GetName() != ds.QueryPriceFeed {
			if mdl, ok := adapter.(*chainlink_price_feed.ChainlinkPriceFeed); ok {
				mdl.DisableToken(oldTokenOracle.Token, newTokenOracle.BlockNumber, pfVersion)
			}
			if mdl, ok := adapter.(*composite_chainlink.CompositeChainlinkPF); ok {
				mdl.DisableToken(oldTokenOracle.Token, newTokenOracle.BlockNumber, pfVersion)
			}
		} else {
			repo.adapters.GetAggregatedFeed().DisableYearnFeed(
				newTokenOracle.Token, oldFeed, newTokenOracle.BlockNumber, pfVersion)
		}
	}
	// set current state of oracle for token.
	repo.addTokenCurrentOracle(
		newTokenOracle,
	)
	// token oracle
	repo.blocks.SetAndGetBlock(newTokenOracle.BlockNumber).AddTokenOracle(
		newTokenOracle,
	)
}

func (repo *TokenOracleRepo) DirectlyAddTokenOracle(newTokenOracle *schemas.TokenOracle) {
	repo.disablePrevAdapterAndAddNewTokenOracle(newTokenOracle)
}

// called from chainlink feed and price oracle
func (repo *TokenOracleRepo) AddNewPriceOracleEvent(newTokenOracle *schemas.TokenOracle, bounded bool, forChainlinkNewFeed ...bool) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	// SPECIAL CASE
	// for token 0xf568F6C71aE0439B8d3FFD60Ceba9B1DcB5819bF on mainnet, while adding blocker token for v1
	// two feeds where added for same block(15371802)
	// ignore 0xBc1c306920309F795fB5A740083eCBf5057349e9 at log 202
	// use 0xAaaF70b91877966900F0EfC0f2E7296e4F86B119 at log 212
	if newTokenOracle.Feed == "0xBc1c306920309F795fB5A740083eCBf5057349e9" && newTokenOracle.BlockNumber == 15371802 {
		return
	}
	pfVersion := schemas.VersionToPFVersion(newTokenOracle.Version, newTokenOracle.Reserve)
	switch newTokenOracle.FeedType {
	case ds.ZeroPF, ds.AlmostZeroPF,
		ds.CurvePF, ds.SingleAssetPF, ds.YearnPF, ds.RedStonePF, ds.CompositeRedStonePF:
		if repo.alreadyActiveFeedForToken(newTokenOracle) {
			return
		}
		repo.disablePrevAdapterAndAddNewTokenOracle(newTokenOracle)
		//
		if newTokenOracle.FeedType == ds.ZeroPF || newTokenOracle.FeedType == ds.AlmostZeroPF {
			priceBI := new(big.Int)
			if ds.AlmostZeroPF == newTokenOracle.FeedType {
				priceBI = big.NewInt(100)
			}
			repo.blocks.AddPriceFeed(&schemas.PriceFeed{
				BlockNumber:     newTokenOracle.BlockNumber,
				Token:           newTokenOracle.Token,
				Feed:            newTokenOracle.Oracle,
				RoundId:         0,
				PriceBI:         (*core.BigInt)(priceBI),
				Price:           utils.GetFloat64Decimal(priceBI, pfVersion.Decimals()),
				MergedPFVersion: schemas.MergedPFVersion(pfVersion), // for 0 and almost zero pf
			})
			repo.zeroPFs[newTokenOracle.Oracle] = true // oracle and feed are same for non-chainlink price feed
		} else {
			repo.adapters.GetAggregatedFeed().AddFeedOrToken(
				newTokenOracle.Token,
				newTokenOracle.Oracle,
				newTokenOracle.FeedType,
				newTokenOracle.BlockNumber,
				pfVersion,
			)
		}
	case ds.ChainlinkPriceFeed:
		obj := chainlink_price_feed.NewChainlinkPriceFeed(
			repo.client, repo.repo,
			newTokenOracle.Token,
			newTokenOracle.Oracle,
			newTokenOracle.BlockNumber,
			schemas.MergedPFVersion(pfVersion),
			bounded,
			forChainlinkNewFeed...,
		)
		newTokenOracle.Feed = obj.Address
		//
		if repo.alreadyActiveFeedForToken(newTokenOracle) {
			return
		}
		repo.disablePrevAdapterAndAddNewTokenOracle(newTokenOracle)
		//
		if adapter := repo.adapters.GetAdapter(obj.Address); adapter != nil {
			adapter.(*chainlink_price_feed.ChainlinkPriceFeed).AddToken(newTokenOracle.Token, newTokenOracle.BlockNumber, pfVersion)
			return
		}
		// SPECIAL CASE
		// on goerli, there are two v2 priceoracles added
		// on first priceoracle cvx token is 0x9683a59Ad8D7B5ac3eD01e4cff1D1A2a51A8f1c0
		// and on second priceoracle it is 0x6D75eb70402CF06a0cB5B8fdc1836dAe29702B17
		// and both uses the same chainlink price feed
		// REASON: we only support 1 chainlink per token, two different tokens can't share chainlink feed
		// so ignore the first cvx token as it is not used anywhere
		if newTokenOracle.Token != "0x9683a59Ad8D7B5ac3eD01e4cff1D1A2a51A8f1c0" {
			repo.adapters.AddSyncAdapter(obj)
		}
	case ds.CompositeChainlinkPF:
		obj := composite_chainlink.NewCompositeChainlinkPF(
			newTokenOracle.Token,
			newTokenOracle.Oracle,
			newTokenOracle.BlockNumber,
			repo.client, repo.repo,
			newTokenOracle.Version,
			newTokenOracle.Reserve,
		)
		//
		if repo.alreadyActiveFeedForToken(newTokenOracle) {
			return
		}
		repo.disablePrevAdapterAndAddNewTokenOracle(newTokenOracle)

		if adapter := repo.adapters.GetAdapter(obj.Address); adapter != nil {
			adapter.(*composite_chainlink.CompositeChainlinkPF).AddToken(newTokenOracle.Token, newTokenOracle.BlockNumber, pfVersion)
			return
		}
		repo.adapters.AddSyncAdapter(obj)
	default:
		log.Fatal(newTokenOracle.FeedType, "not handled")
	}
}

func (repo *TokenOracleRepo) GetTokenOracles() map[schemas.PFVersion]map[string]*schemas.TokenOracle {
	return repo.tokensCurrentOracle
}

// if returned value is nil, it means that token oracle hasn't been added yet.
// func (repo *TokenOracleRepo) GetOracleForV2Token(token string) *schemas.TokenOracle {
// 	obj := repo.tokensCurrentOracle[core.NewVersion(2)][token]
// 	return obj
// }
