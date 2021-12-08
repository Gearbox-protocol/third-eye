package repository

import (
	"github.com/Gearbox-protocol/third-eye/artifacts/dataCompressor"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/ethereum/go-ethereum/common"
	"sort"
)

func (repo *Repository) loadCreditSessions(lastDebtSync int64) {
	data := []*core.CreditSession{}
	err := repo.db.Find(&data, "status = ? OR (status <> ? AND closed_at > ?)",
		core.Active, core.Active, lastDebtSync).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, session := range data {
		repo.AddCreditSession(session)
	}
}

func (repo *Repository) AddDataCompressor(blockNum int64, addr string) {
	dc, err := dataCompressor.NewDataCompressor(common.HexToAddress(addr), repo.client)
	if err != nil {
		log.Fatal(err)
	}
	repo.dc[blockNum] = dc
	repo.dcBlockNum = append(repo.dcBlockNum, blockNum)
	arr := repo.dcBlockNum
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	repo.dcBlockNum = arr
}

func (repo *Repository) AddCreditSession(session *core.CreditSession) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.sessions[session.ID] == nil {
		log.Info("Add creditSession(%s) with id %s", session.Account, session.ID)
		repo.sessions[session.ID] = session
	} else {
		log.Fatalf("Credit session already present %s", session.ID)
	}
}

func (repo *Repository) GetDataCompressor(blockNum int64) *dataCompressor.DataCompressor {
	var dc *dataCompressor.DataCompressor
	for _, num := range repo.dcBlockNum {
		// dc should be deployed before it is queried
		if num < blockNum {
			dc = repo.dc[num]
		}
	}
	return dc
}

func (repo *Repository) GetCreditSession(sessionId string) *core.CreditSession {
	return repo.sessions[sessionId]
}
