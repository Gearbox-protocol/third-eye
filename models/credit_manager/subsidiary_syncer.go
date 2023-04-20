package credit_manager

import (
	"sort"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type SubsidiarySyncer struct {
	client  core.ClientI
	Address common.Address
	topics  [][]common.Hash
	logs    []types.Log
	to      int64
}

func NewSubsidiarySyncer(client core.ClientI, address string, topics [][]common.Hash) *SubsidiarySyncer {
	return &SubsidiarySyncer{
		client:  client,
		Address: common.HexToAddress(address),
		topics:  topics,
	}
}

func (mdl *SubsidiarySyncer) FetchLogs(from, to int64) {
	if to == 0 {
		return
	}
	if mdl.to != 0 {
		if mdl.to+1 != from {
			log.Fatal("Not implementated")
		}
	}
	if len(mdl.logs) != 0 {
		log.Fatal("Previous logs not processed")
	}
	logs, err := pkg.Node{Client: mdl.client}.GetLogs(from, to, []common.Address{mdl.Address}, mdl.topics)
	log.CheckFatal(err)
	mdl.logs = append(mdl.logs, logs...)
	mdl.to = to
}

func (mdl *SubsidiarySyncer) GetLogsBefore(marker types.Log) []types.Log {
	ind := sort.Search(len(mdl.logs), func(i int) bool {
		return mdl.logs[i].BlockNumber > marker.BlockNumber ||
			(mdl.logs[i].BlockNumber == marker.BlockNumber && mdl.logs[i].Index > marker.Index)
	})
	ans := mdl.logs[:ind]
	mdl.logs = mdl.logs[ind:]
	return ans
}
