package tests

import (
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/tests/framework"
	"testing"
)

func TestSyncAdapters(t *testing.T) {
	r, _ := framework.NewEngs(t, []string{"sync_adapters/input.json"})
	r.Eng.Sync(10)
	adapters := getAdapters(r.Repo.GetKit())
	r.Check(map[string]interface{}{"data": adapters}, "sync_adapters/adapters.json")
}

func getAdapters(kit *core.AdapterKit) (array []*core.SyncAdapter) {
	for lvlIndex := 0; lvlIndex < kit.Len(); lvlIndex++ {
		for kit.Next(lvlIndex) {
			array = append(array, kit.Get(lvlIndex).GetAdapterState())
		}
		kit.Reset(lvlIndex)
	}
	return
}
