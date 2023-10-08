package cm_v2

import (
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *CMv2) addProtocolAdapters(state dc.CMCallData) {
	newProtocols := map[string]bool{}
	for _, entry := range state.Adapters {
		newProtocols[entry.TargetContract.Hex()] = true
	}
	mdl.allowedProtocols = newProtocols
}
func (mdl *CMv2) addProtocolAdaptersLocally() {
	// cm is registered with dataCompressor after discoveredAt, so we can get adapters for blockNum more than discoveredAt
	blockToFetchCMData := mdl.DiscoveredAt
	if blockToFetchCMData < mdl.LastSync {
		blockToFetchCMData = mdl.LastSync
	}
	//
	call, resultFn, err := mdl.Repo.GetDCWrapper().GetCreditManagerData(mdl.GetVersion(),
		blockToFetchCMData, common.HexToAddress(mdl.GetAddress()))
	if err != nil {
		log.Fatal("Failed preparing credit manager data", err)
	}
	results := core.MakeMultiCall(mdl.Client, blockToFetchCMData, false, []multicall.Multicall2Call{call})
	state, err := resultFn(results[0].ReturnData)
	if err != nil {
		log.Fatalf("Failed call for cm(%s)@%d version %d: %s", mdl.GetAddress(), blockToFetchCMData, mdl.GetVersion(), err)
	}
	mdl.addProtocolAdapters(state)
}
