package cm_common

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
)

func ToJsonBalanceWithRepo(z core.Transfers, repo ds.RepositoryI) core.JsonFloatMap {
	dbFormat := core.JsonFloatMap{}
	for token, amt := range z {
		dbFormat[token] = utils.GetFloat64Decimal(amt, repo.GetToken(token).Decimals)
	}
	return dbFormat
}

func (mdl CommonCMAdapter) GetCreditFacadeAddr() string {
	return mdl.GetDetailsByKey("facade")
}
