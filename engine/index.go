package engine

import (
	"github.com/Gearbox-protocol/gearscan/utils"
	"github.com/Gearbox-protocol/gearscan/config"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/repository"
	"github.com/Gearbox-protocol/gearscan/models/address_provider"
	"fmt"
)

type EngineI interface {
	Sync()
}

type Engine struct {
	config *config.Config
	client *ethclient.Client
	repository repository.RepositoryI
	executeParser *utils.ExecuteParser
}

func NewEngine(config *config.Config,
	ec *ethclient.Client,
	repo repository.RepositoryI,
	ep *utils.ExecuteParser) EngineI {
	return &Engine {
		config: config,
		client: ec,
		repository: repo,
		executeParser: ep,
	}
}
func (e *Engine) init() {
	if len(e.repository.GetSyncAdapters()) == 0 {
		address_provider.NewAddressProvider()
	}
	
}
func (e *Engine) Sync() {
	e.init()
	fmt.Println("sleep")
}