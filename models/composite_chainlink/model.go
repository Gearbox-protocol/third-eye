package composite_chainlink

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ds/dc_wrapper"
	cpf "github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"github.com/ethereum/go-ethereum/common"
)

type CompositeChainlinkPF struct {
	Token string
	*ds.SyncAdapter
	BaseTokenMainAgg *cpf.ChainlinkMainAgg
	MainAgg          *cpf.ChainlinkMainAgg
	TokenETHPrice    *big.Int
	ETHUSDPrice      *big.Int
}

// compositeChainlink price feed has token eth  oracle and eth usd base oracle for calculating the price of token in usd.
func NewCompositeChainlinkPF(token, oracle string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, version int16) *CompositeChainlinkPF {
	oracleAddr := common.HexToAddress(oracle)
	tokenETHPF := getAddrFromRPC(client, "targetETH", oracleAddr, discoveredAt)
	mainAgg := cpf.NewMainAgg(client, tokenETHPF)

	ethUSDPF := getAddrFromRPC(client, "ETHUSD", oracleAddr, discoveredAt)
	baseTokenMainAgg := cpf.NewMainAgg(client, ethUSDPF)
	//
	identifier := common.BytesToAddress(append(oracleAddr.Bytes(), big.NewInt(discoveredAt).Bytes()...))

	//
	mainPhaseAgg, _ := mainAgg.GetPriceFeedAddr(discoveredAt)
	basePhaseAgg, _ := baseTokenMainAgg.GetPriceFeedAddr(discoveredAt)
	compositeMdl := &CompositeChainlinkPF{
		BaseTokenMainAgg: baseTokenMainAgg,
		MainAgg:          mainAgg,
		Token:            token,
		SyncAdapter: &ds.SyncAdapter{
			SyncAdapterSchema: &schemas.SyncAdapterSchema{
				Contract: &schemas.Contract{
					Address:      identifier.Hex(),
					DiscoveredAt: discoveredAt,
					FirstLogAt:   discoveredAt,
					ContractName: ds.CompositeChainlinkPF,
					Client:       client,
				},
				Details: map[string]interface{}{
					"oracle": oracle,
					"token":  token,
					"secAddrs": map[string]interface{}{
						"target":      tokenETHPF.Hex(),
						"base":        ethUSDPF.Hex(),
						"targetPhase": mainPhaseAgg.Hex(),
						"basePhase":   basePhaseAgg.Hex(),
					}},
				LastSync: discoveredAt,
				V:        version,
			},
			HasOnLogs: true,
			Repo:      repo,
		},
	}
	compositeMdl.setPrices(discoveredAt)
	compositeMdl.addPriceToDB(discoveredAt)
	return compositeMdl
}

func NewCompositeChainlinkPFFromAdapter(adapter *ds.SyncAdapter) *CompositeChainlinkPF {
	compositeMdl := &CompositeChainlinkPF{
		SyncAdapter: adapter,
		Token:       adapter.GetDetailsByKey("token"),
	}
	compositeMdl.BaseTokenMainAgg = cpf.NewMainAgg(adapter.Client, compositeMdl.getAddrFromDetails("base"))
	compositeMdl.MainAgg = cpf.NewMainAgg(adapter.Client, compositeMdl.getAddrFromDetails("target"))
	compositeMdl.setPrices(adapter.LastSync)
	compositeMdl.HasOnLogs = true
	return compositeMdl
}

func (mdl *CompositeChainlinkPF) getAddrFromDetails(field string) common.Address {
	m := mdl.Details["secAddrs"].(map[string]interface{})
	return ds.InterfaceToAddr(m[field])
}
func (mdl *CompositeChainlinkPF) setPrices(blockNum int64) {
	pfABI := core.GetAbi("PriceFeed")
	data, err := pfABI.Pack("latestRoundData")
	log.CheckFatal(err)
	calls := []multicall.Multicall2Call{
		{
			Target:   mdl.getAddrFromDetails("target"),
			CallData: data,
		},
		{
			Target:   mdl.getAddrFromDetails("base"),
			CallData: data,
		},
	}
	results := core.MakeMultiCall(mdl.Client, blockNum, false, calls)
	mdl.TokenETHPrice = getPrice(results[0], mdl.getAddrFromDetails("target"))
	mdl.ETHUSDPrice = getPrice(results[1], mdl.getAddrFromDetails("base"))
}

func getPrice(entry multicall.Multicall2Result, feed common.Address) *big.Int {
	pfABI := core.GetAbi("PriceFeed")
	if entry.Success {
		values, err := pfABI.Unpack("latestRoundData", entry.ReturnData)
		log.CheckFatal(err)
		return values[1].(*big.Int)
	}
	log.Fatal("feed(%s) can't fetch price.", feed)
	return nil
}

func getAddrFromRPC(client core.ClientI, targetMethod string, oracle common.Address, blockNum int64) common.Address {
	var sig string
	switch targetMethod {
	case "targetETH":
		sig = "f1a75c6e" // targetEthPriceFeed
	case "ETHUSD":
		sig = "42f6fb29" // ethUsdPriceFeed
	default:
		log.Fatal(targetMethod, "not found")
	}
	tokenETHPFData, err := dc_wrapper.CallFuncWithExtraBytes(client, sig, oracle, blockNum, nil)
	if err != nil {
		log.Fatalf("Oracle(%s) doesn't have valid %s: %s", oracle, targetMethod, err)
	}
	return common.BytesToAddress(tokenETHPFData)
}

func (mdl *CompositeChainlinkPF) AfterSyncHook(syncedTill int64) {
	mdl.SyncAdapter.AfterSyncHook(syncedTill)
}
