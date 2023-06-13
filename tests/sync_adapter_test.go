package tests

import (
	"testing"

	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/tests/framework"
)

func TestSyncAdapters(t *testing.T) {
	r, _ := framework.NewEngs(t, []string{"sync_adapters/input.json"})
	r.Eng.Sync(10)
	adapters := getAdapters(r.Repo.(*repository.Repository).GetKit())
	r.Check(map[string]interface{}{"data": adapters}, "sync_adapters/adapters.json")
	r.Check(r.Repo.GetBlocks()[6], "sync_adapters/blocks.json")
}

func getAdapters(kit *ds.AdapterKit) (array []*ds.SyncAdapter) {
	for lvlIndex := 0; lvlIndex < kit.Len(); lvlIndex++ {
		for kit.Next(lvlIndex) {
			array = append(array, kit.Get(lvlIndex).GetAdapterState()...)
		}
		kit.Reset(lvlIndex)
	}
	return
}
