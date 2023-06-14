package handlers

import (
	"sync"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
)

type ExtrasRepo struct {
	dcWrapper     *dc_wrapper.DataCompressorWrapper
	mu            *sync.Mutex
	executeParser ds.ExecuteParserI
}

func NewExtraRepo(client core.ClientI, ep ds.ExecuteParserI) *ExtrasRepo {
	return &ExtrasRepo{
		dcWrapper:     dc_wrapper.NewDataCompressorWrapper(client),
		executeParser: ep,
	}
}

func (repo *ExtrasRepo) GetDCWrapper() *dc_wrapper.DataCompressorWrapper {
	return repo.dcWrapper
}

func (repo *ExtrasRepo) GetExecuteParser() ds.ExecuteParserI {
	return repo.executeParser
}
