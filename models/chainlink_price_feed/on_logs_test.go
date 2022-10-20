package chainlink_price_feed

import (
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type OnLogsChecker struct {
	ds.DummyRepo
	pfs []*schemas.PriceFeed
}

func (x *OnLogsChecker) AddPriceFeed(pf *schemas.PriceFeed) {
	x.pfs = append(x.pfs, pf)
}
func TestOnLogs(t *testing.T) {
	validPf := &schemas.PriceFeed{
		Feed:        utils.RandomAddr(),
		Token:       utils.RandomAddr(),
		BlockNumber: 1,
		PriceBI:     (*core.BigInt)(big.NewInt(222)),
		RoundId:     3,
	}
	repo := &OnLogsChecker{}
	obj := &ChainlinkPriceFeed{SyncAdapter: &ds.SyncAdapter{
		Repo: repo,
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				Address: validPf.Feed,
			},
			Details: core.Json{"token": validPf.Token},
		},
	}}
	txLogs := []types.Log{
		{
			BlockNumber: 1,
			Index:       1,
			Topics: []common.Hash{
				core.Topic("AnswerUpdated(int256,uint256,uint256)"),
				common.BytesToHash([]byte{1}),
				common.BytesToHash([]byte{1}),
			},
		},
		{
			BlockNumber: uint64(validPf.BlockNumber),
			Index:       2,
			Topics: []common.Hash{
				core.Topic("AnswerUpdated(int256,uint256,uint256)"),
				common.BytesToHash(validPf.PriceBI.Convert().Bytes()),
				common.BytesToHash([]byte{byte(validPf.RoundId)}),
			},
		},
	}
	obj.OnLogs(txLogs)
	if len(repo.pfs) != 1 ||
		repo.pfs[0].BlockNumber != validPf.BlockNumber ||
		repo.pfs[0].Feed != validPf.Feed ||
		repo.pfs[0].RoundId != validPf.RoundId ||
		repo.pfs[0].PriceBI.Cmp(validPf.PriceBI) != 0 {
		t.Fatal(utils.ToJson(repo.pfs))
	}
}
