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
	adapters := getAdapters(r.Repo.(*repository.Repository))
	r.Check(map[string]interface{}{"data": adapters}, "sync_adapters/adapters.json")
	r.Check(r.Repo.GetBlocks()[6], "sync_adapters/blocks.json")
}

func getAdapters(repo *repository.Repository) (array []*ds.SyncAdapter) {
	kit := repo.GetKit()
	for lvlIndex := 0; lvlIndex < kit.Len(); lvlIndex++ {
		for kit.Next(lvlIndex) {
			adapter := kit.Get(lvlIndex)
			if ds.IsWrapperAdapter(adapter.GetName()) {
				continue
			} else {
				array = append(array, adapter.GetAdapterState())
			}
		}
		kit.Reset(lvlIndex)
	}
	for _, adapter := range repo.GetAdaptersFromWrapper() {
		array = append(array, adapter.GetAdapterState())
	}
	return
}
