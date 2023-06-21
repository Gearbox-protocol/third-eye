package cm_v2

import (
	dcv2 "github.com/Gearbox-protocol/sdk-go/artifacts/dataCompressor/dataCompressorv2"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl *CMv2) addProtocolAdapters(state dcv2.CreditManagerData) {
	newProtocols := map[string]bool{}
	for _, entry := range state.Adapters {
		newProtocols[entry.AllowedContract.Hex()] = true
	}
	mdl.allowedProtocols = newProtocols
}
func (mdl *CMv2) addProtocolAdaptersLocally(blockNum int64) {
	call, resultFn, err := mdl.Repo.GetDCWrapper().GetCreditManagerData(blockNum, common.HexToAddress(mdl.GetAddress()))
	if err != nil {
		log.Fatal("Failed preparing credit manager data", err)
	}
	results := core.MakeMultiCall(mdl.Client, blockNum, false, []multicall.Multicall2Call{call})
	state, err := resultFn(results[0].ReturnData)
	if err != nil {
		log.Fatal("Failed call", err)
	}
	mdl.addProtocolAdapters(state)
}
