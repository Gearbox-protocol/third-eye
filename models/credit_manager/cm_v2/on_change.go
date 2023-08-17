package cm_v2

import (
	"math/big"

	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_common"
	"github.com/Gearbox-protocol/third-eye/models/credit_manager/cm_v1"
	"github.com/ethereum/go-ethereum/core/types"
)

type RewardClaimDetails struct {
	Args      core.Json
	TxHash    string
	BlockNum  int64
	LogID     uint
	Transfers core.Transfers
	Borrower  string
	SessionId string
}

func newRewardClaimDetails() RewardClaimDetails {
	return RewardClaimDetails{
		Args:      core.Json{},
		Transfers: core.Transfers{},
	}
}

// /////////
// Direct Token Transfer
// /////////
// if tx is nil and currentReward.TxHash is not null, add the rewards AccountOperation
func (mdl *CMv2) getDirectTokenTransferFn() cm_common.OnDirectTokenTransferFn {
	currentRewardClaim := newRewardClaimDetails()
	return func(repo ds.RepositoryI, tx *schemas.TokenTransfer, session *schemas.CreditSession) {
		if tx == nil {
			if currentRewardClaim.TxHash != "" {
				addRewardClaimAccountOperation(repo, currentRewardClaim)
				currentRewardClaim = newRewardClaimDetails()
			}
			return
		}
		if tx.From == core.NULL_ADDR.Hex() || mdl.allowedProtocols[tx.From] {
			// different rewardclaim
			if currentRewardClaim.TxHash != "" && currentRewardClaim.TxHash != tx.TxHash {
				addRewardClaimAccountOperation(repo, currentRewardClaim)
				currentRewardClaim = newRewardClaimDetails()
			}
			currentRewardClaim.Args[tx.Token] = tx.From
			currentRewardClaim.Transfers[tx.Token] = tx.Amount.Convert()
			currentRewardClaim.TxHash = tx.TxHash
			currentRewardClaim.BlockNum = tx.BlockNum
			currentRewardClaim.LogID = tx.LogID
			currentRewardClaim.SessionId = session.ID
			currentRewardClaim.Borrower = session.Borrower
		} else {
			cm_v1.OnDirectTokenTransfer(repo, tx, session)
		}
	}
}
func addRewardClaimAccountOperation(repo ds.RepositoryI, claim RewardClaimDetails) {
	repo.AddAccountOperation(&schemas.AccountOperation{
		TxHash:      claim.TxHash,
		BlockNumber: claim.BlockNum,
		LogId:       claim.LogID,
		Borrower:    claim.Borrower,
		SessionId:   claim.SessionId,
		Dapp:        core.NULL_ADDR.Hex(),
		Action:      "RewardClaimed",
		Args:        &claim.Args,
		AdapterCall: false,
		Transfers:   &claim.Transfers,
	})
}

// /////////
// OnLog
// /////////
func (mdl *CMv2) OnLog(txLog types.Log) {
	// creditConfigurator events for test
	// CreditFacadeUpgraded is emitted when creditconfigurator is initialized, so we will receive it on init
	// although we have already set creditfacadeUpgra
	if mdl.GetDetailsByKey("configurator") == txLog.Address.Hex() {
		switch txLog.Topics[0] {
		case core.Topic("CreditFacadeUpgraded(address)"):
			facade := utils.ChecksumAddr(txLog.Topics[1].Hex())
			mdl.setCreditFacadeSyncer(facade)
		case core.Topic("FeesUpdated(uint16,uint16,uint16,uint16,uint16)"):
			mdl.SetParams(&schemas.Parameters{
				BlockNum:                   int64(txLog.BlockNumber),
				CreditManager:              mdl.Address,
				FeeInterest:                bytesToUInt16(txLog.Data[:32]),
				FeeLiquidation:             bytesToUInt16(txLog.Data[32:64]),
				LiquidationDiscount:        10000 - bytesToUInt16(txLog.Data[64:96]), // 10000- liqPremium
				FeeLiquidationExpired:      bytesToUInt16(txLog.Data[96:128]),
				LiquidationDiscountExpired: 10000 - bytesToUInt16(txLog.Data[128:160]), // 10000- liqPremiumExpired
			})
		}
		return
	}
	//
	mdl.CmMVP.PrefixOnLog(txLog)
	mdl.checkLogV2(txLog)
}

func bytesToUInt16(data []byte) uint16 {
	return uint16(new(big.Int).SetBytes(data).Int64())
}

func (mdl *CMv2) SetOnChangeFn() {
	mdl.SetLastTxHashCompleted(mdl.lastTxHashCompleted)
	mdl.SetCalculateCMStatFn(func(blockNum int64, state dcv2.CreditManagerData) {
		mdl.addProtocolAdapters(state)
		mdl.CmMVP.CalculateCMStat(blockNum, state)
	})
	mdl.SetOnDirectTokenTransferFn(mdl.getDirectTokenTransferFn())
}

// /////////
// On TxHash
// /////////
func (mdl *CMv2) lastTxHashCompleted(lastTxHash string) {
	nonMulticallExecuteEvents := mdl.ProcessNonMultiCalls()
	mdl.ProcessRemainingMultiCalls(mdl.GetVersion(), lastTxHash, nonMulticallExecuteEvents)
}
