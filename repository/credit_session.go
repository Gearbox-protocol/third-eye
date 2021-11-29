package repository

import (
	"github.com/Gearbox-protocol/gearscan/artifacts/dataCompressor"
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

func (repo *Repository) loadCreditSessions() {
	data := []*core.CreditSession{}
	err := repo.db.Find(&data, "status = ?", core.Active).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, session := range data {
		repo.AddCreditSession(session)
	}
}

func (repo *Repository) AddDataCompressor(addr string) {
	dc, err := dataCompressor.NewDataCompressor(common.HexToAddress(addr), repo.client)
	if err != nil {
		log.Fatal()
	}
	repo.dc = dc
}

func (repo *Repository) AddCreditSession(session *core.CreditSession) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.sessions[session.ID] == nil {
		repo.sessions[session.ID] = session
	} else {
		log.Fatalf("Credit session already present %s", session.ID)
	}

}

func (repo *Repository) GetCreditSessionData(blockNum int64, sessionId string) dataCompressor.DataTypesCreditAccountDataExtended {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	session := repo.GetCreditSession(sessionId)
	data, err := repo.dc.GetCreditAccountDataExtended(opts,
		common.HexToAddress(session.CreditManager),
		common.HexToAddress(session.Borrower),
	)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func (repo *Repository) GetCreditSession(sessionId string) *core.CreditSession {
	return repo.sessions[sessionId]
}
