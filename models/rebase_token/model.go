package rebase_token

import (
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/third-eye/ds"
)

type RebaseToken struct {
	*ds.SyncAdapter
}

func NewRebaseToken(addr string, client core.ClientI, repo ds.RepositoryI) *RebaseToken {
	var startFrom int64 = 0
	switch core.GetChainId(client) {
	case 1:
		startFrom = 17266004
	case 5:
		startFrom = 7692351
	}
	adapter := &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			LastSync: startFrom,
			V:        core.NewVersion(1),
			Details:  core.Json{},
			Contract: &schemas.Contract{
				DiscoveredAt: startFrom,
				FirstLogAt:   startFrom,
				Address:      addr,
				ContractName: ds.RebaseToken,
				Client:       client,
			},
		},
		DataProcessType: ds.ViaQuery,
		Repo:            repo,
	}

	mdl := NewRebaseTokenFromAdapter(adapter)
	details := mdl.GetstETHDetails(startFrom)
	mdl.Repo.AddRebaseDetailsForDB(details.ToDB())
	mdl.save(details)
	return mdl
}

func NewRebaseTokenFromAdapter(adapter *ds.SyncAdapter) *RebaseToken {
	obj := &RebaseToken{
		SyncAdapter: adapter,
	}
	obj.DataProcessType = ds.ViaQuery
	return obj
}
