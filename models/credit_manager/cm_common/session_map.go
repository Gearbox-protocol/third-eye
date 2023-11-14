package cm_common

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type SessionUpdateDetails struct {
	Count              int
	LiqUpdatev3Details *SessionLiqUpdatev3Details
}

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

func (mdl CommonCMAdapter) SetSessionIsUpdated(sessionId string, data ...*SessionLiqUpdatev3Details) {
	if _, ok := mdl.UpdatedSessions[sessionId]; !ok {
		mdl.UpdatedSessions[sessionId] = &SessionUpdateDetails{}
	}
	mdl.UpdatedSessions[sessionId].Count++
	if len(data) > 0 {
		if mdl.UpdatedSessions[sessionId].LiqUpdatev3Details != nil {
			log.Fatal("can't set the liquidatev3 details for fetching dcv3 data twice", sessionId, utils.ToJson(data))
		}
		mdl.UpdatedSessions[sessionId].LiqUpdatev3Details = data[0]
	}
}

func (mdl CommonCMAdapter) SetSessionIsClosed(sessionId string, details *SessionCloseDetails) {
	mdl.ClosedSessions[sessionId] = details
}

// used for v2 liquidations setting expired,paused or noraml liquidation
func (mdl CommonCMAdapter) UpdateClosedSessionStatus(sessionId string, status int) {
	mdl.ClosedSessions[sessionId].Status = status
}

// collateral
func (mdl CommonCMAdapter) AddCollateralToSession(blockNum int64, sessionId, token string, amount *big.Int) {
	if !mdl.Repo.IsDieselToken(token) && mdl.Repo.GetGearTokenAddr() != token {
		session := mdl.Repo.GetCreditSession(sessionId)
		//
		if session.Collateral == nil {
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

func (mdl CommonCMAdapter) AddCollateralForOpenCreditAccount(blockNum int64, mainAction *schemas.AccountOperation) {
	collateral := mdl.GetCollateralAmount(blockNum, mainAction)
	(*mainAction.Args)["amount"] = collateral.String()
	mdl.Repo.UpdateCreditSession(mainAction.SessionId, map[string]interface{}{"InitialAmount": collateral})
}

// TO CHECK
func (mdl CommonCMAdapter) GetCollateralAmount(blockNum int64, mainAction *schemas.AccountOperation) *big.Int {
	balances := map[string]*big.Int{}
	for _, event := range mainAction.MultiCall {
		if event.Action == "AddCollateral(address,address,uint256)" {
			for token, amount := range *event.Transfers {
				if balances[token] == nil {
					balances[token] = new(big.Int)
				}
				balances[token] = new(big.Int).Add(balances[token], amount)
			}
		}
	}
	tokens := make([]string, 0, len(balances)+1)
	for token := range balances {
		tokens = append(tokens, token)
	}
	underlyingToken := mdl.GetUnderlyingToken()
	if balances[underlyingToken] == nil {
		tokens = append(tokens, underlyingToken)
	}
	//
	prices := mdl.Repo.GetPricesInUSD(blockNum, tokens)
	underlyingDecimals := mdl.GetUnderlyingDecimal()
	//
	totalValue := new(big.Float)
	// sigma(tokenAmount(i)*price(i)/exp(tokendecimals- underlyingToken))/price(underlying)
	for token, amount := range balances {
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
	initialAmount, _ := new(big.Float).Quo(totalValue, big.NewFloat(prices[underlyingToken])).Int(nil)

	if balances[underlyingToken] != nil { // directly add collateral for underlying token
		initialAmount = new(big.Int).Add(initialAmount, balances[underlyingToken])
	}
	if initialAmount == nil || initialAmount.Cmp(new(big.Int)) == 0 {
		log.Fatal("Collateral for opencreditaccount v2 is zero or nil")
	}
	return initialAmount
}
