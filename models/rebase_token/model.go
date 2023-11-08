package rebase_token

import (
	"fmt"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/rebaseToken"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type RebaseToken struct {
	*ds.SyncAdapter
	contract            *rebaseToken.RebaseToken
	state               *schemas.RebaseTokenDetails
	knownImplmentations []string
	lastBlockNum        int64
	prevRatio           *big.Int
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
				FirstLogAt:   0,
				Address:      addr,
				ContractName: ds.RebaseToken,
				Client:       client,
			},
		},
		DataProcessType: ds.ViaLog,
		Repo:            repo,
	}

	obj := NewRebaseTokenFromAdapter(adapter)

	return obj
}

func (obj *RebaseToken) operationsAtInit(startFrom int64) {
	// get state
	obj.state = obj.getStateAt(startFrom)
	// set kernel
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(startFrom),
	}
	if kernel, err := obj.contract.Kernel(opts); err == nil {
		obj.state.Kernel = kernel
	}
	// check implementation addr
	implementationAddr, err := obj.contract.Implementation(opts)
	log.CheckFatal(err)
	obj.checkImplementationAddr(implementationAddr.Hex(), fmt.Sprintf("%d", startFrom))
	log.Infof("RebaseToken init state at %d:  %s", startFrom, utils.ToJson(obj.state))
	// discard before updates
	// obj.validatorHandler.GetValuesBefore(startFrom)
}

func (mdl RebaseToken) checkImplementationAddr(implementationAddr string, txHash string) {
	if !utils.Contains(mdl.knownImplmentations, implementationAddr) {
		log.Fatalf("implmentation addr(%s) at txHash(%s) is unknown", implementationAddr, txHash)
	}
}
func NewRebaseTokenFromAdapter(adapter *ds.SyncAdapter) *RebaseToken {
	contract, err := rebaseToken.NewRebaseToken(common.HexToAddress(adapter.Address), adapter.Client)
	log.CheckFatal(err)
	obj := &RebaseToken{
		SyncAdapter: adapter,
		contract:    contract,
		knownImplmentations: []string{
			"0xC7B5aF82B05Eb3b64F12241B04B2cF14469E39F7", // at addressprovider discoveredAt for mainnet , this implmenetation doesn't have event for tracking deposited_validators_position
			"0x47EbaB13B806773ec2A2d16873e2dF770D130b50", // same as above implemenation but was added at 14860268
			"0x17144556fd3424EDC8Fc8A4C940B2D04936d17eb", // 2 for steth on mainnet, 17266004
		},
		prevRatio: new(big.Int),
		// validatorHandler: NewValidatorHandler(core.GetChainId(adapter.Client)),
	}
	if obj.Details["kernel"] == nil {
		obj.operationsAtInit(obj.LastSync)
	} else {
		obj.state = &schemas.RebaseTokenDetails{}
		obj.state.Unserialize(obj.Details)
	}
	return obj
}

// buffered -> transient -> cl
func (mdl RebaseToken) Topics() [][]common.Hash {
	return [][]common.Hash{{
		// kernel app change for rebaseToken steth
		core.Topic("SetApp(bytes32,bytes32,address)"),
		// deposit
		// - position: DEPOSITED_VALIDATORS_POSITION
		core.Topic("DepositedValidatorsChanged(uint256)"), // DepositedValidatorsChanged [final value]
		// - position: BUFFERED_ETHER_POSITION
		core.Topic("Unbuffered(uint256)"),                                             // Unbuffered [-delta]
		core.Topic("Submitted(address,uint256,address)"),                              // Submitted [delta]
		core.Topic("ETHDistributed(uint256,uint256,uint256,uint256,uint256,uint256)"), // ETHDistributed(last field) [final value] subtracted value due to withdrawals
		//
		// - position: CL_VALIDATORS_POSITION
		core.Topic("CLValidatorsUpdated(uint256,uint256,uint256)"), // CLValidatorsUpdated(last field)  [final value]
		// - position: CL_BALANCE_POSITION
		// ETHDistributed(postCLBalance)
		//
		// - check
		// TokenRebased
		core.Topic("TokenRebased(uint256,uint256,uint256,uint256,uint256,uint256,uint256)"),
		// - for totalShares
		// shares
		core.Topic("SharesBurnt(address,uint256,uint256,uint256)"),
		core.Topic("TransferShares(address,address,uint256)"),
	}}
}

func (mdl RebaseToken) GetAllAddrsForLogs() []common.Address {
	return []common.Address{
		common.HexToAddress(mdl.Address),
		mdl.state.Kernel,
	}
}

func (mdl *RebaseToken) SetUnderlyingState(state interface{}) {
	switch v := state.(type) {
	case *schemas.RebaseDetailsForDB:
		mdl.prevRatio = getETHToSharesRatio(v)
	}
}
