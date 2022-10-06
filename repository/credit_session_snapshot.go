package repository

import (
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
)

func (repo *Repository) AddCreditSessionSnapshot(css *schemas.CreditSessionSnapshot) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	css.ID = 0
	repo.SetAndGetBlock(css.BlockNum).AddCreditSessionSnapshot(css)
}
