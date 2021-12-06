package repository

import "github.com/Gearbox-protocol/third-eye/core"

// for credit filter
func (repo *Repository) AddAllowedProtocol(p *core.Protocol) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[p.BlockNumber].AddAllowedProtocol(p)
}

func (repo *Repository) AddAllowedToken(atoken *core.AllowedToken) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	repo.blocks[atoken.BlockNumber].AddAllowedToken(atoken)
}
