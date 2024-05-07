package debts

import (
	"math"

	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

// token threshold
// before blockNum let the latest LT of all cm/token pair
func (eng *DebtEngine) loadAllowedTokenThreshold(lastDebtSync int64) {
	defer utils.Elapsed("Debt(loadAllowedTokenThreshold)")()
	data := []*schemas.AllowedToken{}
	query := `SELECT * FROM allowed_tokens 
	JOIN (SELECT max(block_num) as bn, token, credit_manager FROM allowed_tokens 
		WHERE block_num <= ? group by token,credit_manager) as atokens
	ON atokens.bn = allowed_tokens.block_num
	AND atokens.credit_manager = allowed_tokens.credit_manager
	AND atokens.token = allowed_tokens.token
	WHERE block_num <= ? ORDER BY block_num;`
	err := eng.db.Raw(query, lastDebtSync, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, atoken := range data {
		eng.AddAllowedTokenThreshold(atoken)
	}
}

func (eng *DebtEngine) AddAllowedTokenThreshold(atoken *schemas.AllowedToken) {
	var lt uint16 = 0
	if atoken.LiquidityThreshold != nil {
		lt = uint16(atoken.LiquidityThreshold.Convert().Int64())
	}
	eng.AddTokenLTRamp(&schemas_v3.TokenLTRamp{
		BlockNum:      atoken.BlockNumber,
		CreditManager: atoken.CreditManager,
		Token:         atoken.Token,
		LtInitial:     lt,
		LtFinal:       lt,
		RampStart:     math.MaxInt64,
		RampEnd:       0,
	})
}

func (eng *DebtEngine) loadLastLTRamp(lastDebtSync int64) {
	defer utils.Elapsed("LastLTRamp()")()
	data := []*schemas_v3.TokenLTRamp{}
	query := `SELECT DISTINCT ON (credit_manager, token) *
	 FROM token_ltramp WHERE block_num <= ? 
	 ORDER BY credit_manager, token, block_num DESC;`
	err := eng.db.Raw(query, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, atoken := range data {
		eng.AddTokenLTRamp(atoken)
	}
}

func (eng *DebtEngine) AddTokenLTRamp(atoken *schemas_v3.TokenLTRamp) {
	if eng.tokenLTRamp[atoken.CreditManager] == nil {
		eng.tokenLTRamp[atoken.CreditManager] = make(map[string]*schemas_v3.TokenLTRamp)
	}
	if eng.tokenLTRamp[atoken.CreditManager][atoken.Token] != nil {
		block := eng.tokenLTRamp[atoken.CreditManager][atoken.Token].BlockNum
		if block > atoken.BlockNum { // no need to update if the block is latest already
			return
		}
	}
	eng.tokenLTRamp[atoken.CreditManager][atoken.Token] = atoken
}

// token price from feeds
func (eng *DebtEngine) loadTokenLastPrice(lastDebtSync int64) {
	defer utils.Elapsed("Debt(loadTokenLastPrice)")()
	data := []*schemas.PriceFeed{}
	query := `select * from (SELECT distinct on (token, merged_pf_version) * FROM price_feeds WHERE block_num <= ? ORDER BY token, merged_pf_version, block_num DESC) t order by block_num;`
	err := eng.db.Raw(query, lastDebtSync, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, tokenPrice := range data {
		eng.AddTokenLastPrice(tokenPrice)
	}
}

func (eng *DebtEngine) AddTokenLastPrice(pf *schemas.PriceFeed) {
	for _, pfVersion := range pf.MergedPFVersion.MergedPFVersionToList() {
		if eng.tokenLastPrice[pfVersion] == nil {
			eng.tokenLastPrice[pfVersion] = make(map[string]*schemas.PriceFeed)
		}
		eng.tokenLastPrice[pfVersion][pf.Token] = pf
	}
}
