package repository

import "github.com/Gearbox-protocol/sdk-go/core/schemas"

func (repo *Repository) AddFastCheckParams(logID uint, txHash, cm, creditFilter string, fcParams *schemas.FastCheckParams) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(fcParams.BlockNum).AddFastCheckParams(fcParams)
	daoOperation := repo.ParamsRepo.AddFastCheckParams(logID, txHash, cm, creditFilter, fcParams)
	repo.addDAOOperation(daoOperation)
}

func (repo *Repository) AddParameters(logID uint, txHash string, params *schemas.Parameters, token string) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.setAndGetBlock(params.BlockNum).AddParameters(params)
	daoOperation := repo.ParamsRepo.AddParameters(logID, txHash, params, token)
	repo.addDAOOperation(daoOperation)
}

func (repo *Repository) UpdateLimits(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	newParams, daoOperation := repo.ParamsRepo.UpdateLimits(logID, txHash, creditConfigurator, params)
	repo.setAndGetBlock(params.BlockNum).AddParameters(newParams)
	repo.addDAOOperation(daoOperation)
}

func (repo *Repository) UpdateFees(logID uint, txHash, creditConfigurator string, params *schemas.Parameters) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	newParams, daoOperation := repo.ParamsRepo.UpdateFees(logID, txHash, creditConfigurator, params)
	repo.setAndGetBlock(params.BlockNum).AddParameters(newParams)
	repo.addDAOOperation(daoOperation)
}
