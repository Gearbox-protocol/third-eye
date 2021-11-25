package repository

import (
	"github.com/Gearbox-protocol/gearscan/core"
	"github.com/Gearbox-protocol/gearscan/log"
)

func (repo *Repository) loadCreditManagers() {
	data := []*core.CreditManager{}
	err := repo.db.Find(&data).Error
	if err != nil {
		log.Fatal(err)
	}
	for _, cm := range data {
		repo.AddCreditManager(cm)
	}
}

func (repo *Repository) AddCreditManager(cm *core.CreditManager) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.creditManagers[cm.Address] != nil {
		log.Fatal("credit manager already set")
	}
	repo.creditManagers[cm.Address] = cm
}

func (repo *Repository) AddCreditOwnerSession(cmAddr, owner, sessionId string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.creditManagers[cmAddr] == nil {
		log.Fatal("credit manager not found ", cmAddr)
	}
	repo.creditManagers[cmAddr].Sessions.Set(owner,sessionId)
}

func (repo *Repository) RemoveCreditOwnerSession(cmAddr, owner string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.creditManagers[cmAddr] == nil {
		log.Fatal("credit manager not found ", cmAddr)
	}
	repo.creditManagers[cmAddr].Sessions.Remove(owner)
}

func (repo *Repository) GetCreditOwnerSession(cmAddr, owner string) string {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	if repo.creditManagers[cmAddr] == nil {
		log.Fatal("credit manager not found ", cmAddr)
	}
	return repo.creditManagers[cmAddr].Sessions.Get(owner)
}