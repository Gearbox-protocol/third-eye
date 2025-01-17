package cm_common

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type SessionLiqUpdatev3Details SessionCloseDetails

// credit session close details
type SessionCloseDetails struct {
	RemainingFunds *big.Int
	Status         int
	LogId          uint
	TxHash         string
	Borrower       string
	// used for v1 repayment , so that we can set repayment amount for repaid session accountOperation
	AccountOperation *schemas.AccountOperation
}

func (x SessionCloseDetails) String() string {
	return fmt.Sprintf("ClosingDetails(Status: %d LogId %d TxHash %s Borrower %s RemainingFunds %s)",
		x.Status, x.LogId, x.TxHash, x.Borrower, x.RemainingFunds)
}

func (mdl CommonCMAdapter) SetSessionIsUpdated(sessionId string) {
	mdl.updatedSessions[sessionId] += 1
}
func (mdl CommonCMAdapter) SetSessionIsLiqv3(sessionId string, details *SessionLiqUpdatev3Details) {
	if mdl.liqv3Sessions[sessionId] == nil {
		mdl.liqv3Sessions[sessionId] = details
		mdl.updatedSessions[sessionId] += 1
	} else {
		log.Fatal("can't set the liquidatev3 details for fetching dcv3 data twice", sessionId, utils.ToJson(details))
	}
}

func (mdl CommonCMAdapter) SetSessionIsClosed(sessionId string, details *SessionCloseDetails) {
	mdl.closedSessions[sessionId] = details
}

// used for v2 liquidations setting expired,paused or noraml liquidation
// SET_LIQ_STATUS_AFTER_CALL
func (mdl CommonCMAdapter) UpdateClosedSessionStatus(sessionId string, status int) {
	mdl.closedSessions[sessionId].Status = status
}

// collateral
func (mdl CommonCMAdapter) AddCollateralToSession(blockNum int64, sessionId, token string, amount *big.Int) {
	if !mdl.Repo.IsDieselToken(token) && mdl.Repo.GetGearTokenAddr() != token {
		session := mdl.Repo.GetCreditSession(sessionId)
		//
		if session.Collateral == nil || *session.Collateral == nil {
			session.Collateral = &core.JsonBigIntMap{}
		}
		(*session.Collateral)[token] = (*core.BigInt)(new(big.Int).Add(
			core.NewBigInt((*session.Collateral)[token]).Convert(),
			amount,
		))
		//
		valueInUSD := mdl.Repo.GetValueInCurrency(blockNum, session.Version, token, "USDC", amount)
		session.CollateralInUSD = session.CollateralInUSD + utils.GetFloat64Decimal(valueInUSD, 6)
		valueInUnderlyingAsset := mdl.Repo.GetValueInCurrency(blockNum, session.Version, token, mdl.GetUnderlyingToken(), amount)
		session.CollateralInUnderlying += utils.GetFloat64Decimal(valueInUnderlyingAsset, mdl.GetUnderlyingDecimal())
	}
}

// only used for v2/v3
// if mainEvent is opencreditaccount , called from facade_actions
// if openCreditAccountWithoutMulticall called for v2
func (mdl CommonCMAdapter) AddCollateralForOpenCreditAccount(blockNum int64, version core.VersionType, mainAction *schemas.AccountOperation) {
	collateral := mdl.getCollateralAmountOnOpen(blockNum, version, mainAction)
	(*mainAction.Args)["userFunds"] = collateral.String()
	if version.MoreThanEq(core.NewVersion(300)) {
		borrowedAmount := mdl.getBorrowAmountOnOpen(mainAction)
		(*mainAction.Args)["borrowAmount"] = borrowedAmount.String()
	}
	mdl.Repo.UpdateCreditSession(mainAction.SessionId, map[string]interface{}{"InitialAmount": collateral})
}

func (mdl CommonCMAdapter) getBorrowAmountOnOpen(mainAction *schemas.AccountOperation) *big.Int {
	borrowedAmount := new(big.Int)
	for _, event := range mainAction.MultiCall {
		if event.Action == "IncreaseDebt(address,uint256)" {
			if len(*event.Transfers) != 1 {
				log.Fatal(mainAction.TxHash, mainAction.LogId, " has changed borrowedAmount for more than 1 token.")
			}
			for _, amount := range *event.Transfers {
				borrowedAmount = new(big.Int).Add(borrowedAmount, amount)
			}
		}
	}
	return borrowedAmount
}

func (mdl CommonCMAdapter) getCollateralAmountOnOpen(blockNum int64, version core.VersionType, mainAction *schemas.AccountOperation) *big.Int {
	userFunds := map[string]*big.Int{}
	for _, event := range mainAction.MultiCall {
		if event.Action == "AddCollateral(address,address,uint256)" || // v2,v3
			event.Action == "WithdrawCollateral(address,address,uint256,address)" { // v3
			for token, amount := range *event.Transfers {
				if userFunds[token] == nil {
					userFunds[token] = new(big.Int)
				}
				userFunds[token] = new(big.Int).Add(userFunds[token], amount)
			}
		}
	}
	underlyingToken := mdl.GetUnderlyingToken()
	underlyingDecimals := mdl.GetUnderlyingDecimal()
	//
	prices := func() core.JsonFloatMap {
		tokens := make([]string, 0, len(userFunds)+1)
		for token := range userFunds {
			tokens = append(tokens, token)
		}
		if userFunds[underlyingToken] == nil {
			tokens = append(tokens, underlyingToken)
		}
		//
		return mdl.Repo.GetPricesInUSD(blockNum, mdl.State.PoolAddress, tokens)
	}()
	//
	totalValue := new(big.Float)
	// sigma(tokenAmount(i)*price(i)/exp(tokendecimals- underlyingToken))/price(underlying)
	for token, amount := range userFunds {
		if token == underlyingToken { // directly add collateral for underlying token
			continue
		}
		calcValue := utils.GetFloat64(amount, -1*underlyingDecimals)
		nomunerator := new(big.Float).Mul(calcValue, big.NewFloat(prices[token]))
		//
		tokenDecimals := utils.GetExpFloat(mdl.Repo.GetToken(token).Decimals)
		//
		totalValue = new(big.Float).Add(totalValue, new(big.Float).Quo(nomunerator, tokenDecimals))
	}
	if prices[underlyingToken] == 0 {
		log.Info(underlyingToken, utils.ToJson(prices))
	}
	initialAmount, _ := new(big.Float).Quo(totalValue, big.NewFloat(prices[underlyingToken])).Int(nil)

	if userFunds[underlyingToken] != nil { // directly add collateral for underlying token
		initialAmount = new(big.Int).Add(initialAmount, userFunds[underlyingToken])
	}
	if (initialAmount == nil || initialAmount.Cmp(new(big.Int)) == 0) && !version.MoreThanEq(core.NewVersion(300)) {
		log.Fatal("Collateral for opencreditaccount v2 is zero or nil", blockNum, utils.ToJson(mainAction))
	}
	return initialAmount
}
