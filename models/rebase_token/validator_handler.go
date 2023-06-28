package rebase_token

import (
	"bytes"
	"embed"
	"fmt"
	"sort"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
)

type ValidatorHandler struct {
	validators []BlockAndValidator
	ind        int
}

//go:embed validators_mainnet.json
var validatorFiles embed.FS

func (h *ValidatorHandler) GetValuesBefore(block int64) []BlockAndValidator {
	checkIn := h.validators[h.ind:]
	endInd := sort.Search(len(checkIn), func(i int) bool {
		return checkIn[i].Block > block
	})
	h.ind += endInd
	if endInd > 0 {
		log.Info(block, endInd)
	}
	return checkIn[:endInd]
}

func NewValidatorHandler(chainId int64) *ValidatorHandler {
	validators := []BlockAndValidator{}
	data, err := validatorFiles.ReadFile(
		fmt.Sprintf("validators_%s.json", strings.ToLower(log.GetNetworkName(chainId))),
	)
	log.CheckFatal(err)
	utils.ReadJsonReaderAndSetInterface(bytes.NewBuffer(data), &validators)
	return &ValidatorHandler{
		validators: validators,
	}
}

type BlockAndValidator struct {
	Block       int64        `json:"block"`
	Validator   int64        `json:"validators"`
	TotalShares *core.BigInt `json:"totalShares"`
}
