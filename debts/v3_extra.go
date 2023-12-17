package debts

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas/schemas_v3"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"gorm.io/gorm"
)

type v3DebtDetails struct {
	// v3
	// session_id -> token -> AccountQuotaInfo
	// accountQuotaToken map[string]map[string]*schemas_v3.AccountQuotaInfo
	// pool -> token -> QuotaDetails
	poolQuotaDetails map[string]map[string]*schemas_v3.QuotaDetails
}

func Newv3DebtDetails() v3DebtDetails {
	return v3DebtDetails{
		// accountQuotaToken: map[string]map[string]*schemas_v3.AccountQuotaInfo{},
		poolQuotaDetails: map[string]map[string]*schemas_v3.QuotaDetails{},
	}
}

//	func (eng *v3DebtDetails) loadAccounQuotaInfo(lastDebtSync int64, db *gorm.DB) {
//		defer utils.Elapsed("Debt(loadAccountQuotaToken)")()
//		data := []*schemas_v3.AccountQuotaInfo{}
//		query := `SELECT DISTINCT ON (token, session_id) * FROM account_quota_info aqi WHERE session_id in
//			(SELECT id from credit_sessions WHERE ? >= since and (status=0 ||  closed_at > ?)) AND
//			block_num <=? ORDER by token, session_id, block_num desc`
//		err := db.Raw(query, lastDebtSync, lastDebtSync, lastDebtSync).Find(&data).Error
//		if err != nil {
//			log.Fatal(err)
//		}
//		for _, pd := range data {
//			eng.AddAccounQuotaInfo(pd)
//		}
//	}
func (eng *v3DebtDetails) loadPoolQuotaDetails(lastDebtSync int64, db *gorm.DB) {
	defer utils.Elapsed("Debt(loadPoolQuotaDetails)")()
	data := []*schemas_v3.QuotaDetails{}
	query := `SELECT DISTINCT ON (token, pool) * FROM quota_details qd WHERE 
		block_num <=? 
		ORDER by token, pool, block_num desc`
	err := db.Raw(query, lastDebtSync).Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, pd := range data {
		eng.AddPoolQuotaDetails(pd)
	}
}

//	func (eng v3DebtDetails) AddAccounQuotaInfo(data *schemas_v3.AccountQuotaInfo) {
//		if eng.accountQuotaToken[data.SessionId] == nil {
//			eng.accountQuotaToken[data.SessionId] = map[string]*schemas_v3.AccountQuotaInfo{}
//		}
//		eng.accountQuotaToken[data.SessionId][data.Token] = data
//	}
func (eng v3DebtDetails) AddPoolQuotaDetails(data *schemas_v3.QuotaDetails) {
	if eng.poolQuotaDetails[data.Pool] == nil {
		eng.poolQuotaDetails[data.Pool] = map[string]*schemas_v3.QuotaDetails{}
	}
	eng.poolQuotaDetails[data.Pool][data.Token] = data
}

// func (eng v3DebtDetails) GetAccountQuotaInfo(sessionId string) map[string]*schemas_v3.AccountQuotaInfo {
// 	return eng.accountQuotaToken[sessionId]
// }

func (eng v3DebtDetails) GetPoolQuotaDetails(pool string) map[string]*schemas_v3.QuotaDetails {
	return eng.poolQuotaDetails[pool]
}
