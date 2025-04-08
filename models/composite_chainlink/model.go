package composite_chainlink

import (
	"context"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/ds"
	cpf "github.com/Gearbox-protocol/third-eye/models/chainlink_price_feed"
	"github.com/ethereum/go-ethereum/common"
)

type CompositeChainlinkPF struct {
	*ds.SyncAdapter
	BaseTokenMainAgg *cpf.ChainlinkMainAgg
	MainAgg          *cpf.ChainlinkMainAgg
	TokenETHPrice    *big.Int
	ETHUSDPrice      *big.Int
	decimalsOfBasePF int8
	ansBlock         []int64
	priceAdded       int64
}

// compositeChainlink price feed has token base  oracle and base usd oracle for calculating the price of token in usd.
// address is set as identifier(random), as same oracle can be added for different tokens.
func NewCompositeChainlinkPF(token, oracle string, discoveredAt int64, client core.ClientI, repo ds.RepositoryI, version core.VersionType, priceOracle schemas.PriceOracleT) *CompositeChainlinkPF {
	oracleAddr := common.HexToAddress(oracle)
	//
	adapter := &ds.SyncAdapter{
		SyncAdapterSchema: &schemas.SyncAdapterSchema{
			Contract: &schemas.Contract{
				Address:      oracleAddr.Hex(),
				DiscoveredAt: discoveredAt,
				FirstLogAt:   discoveredAt,
				ContractName: ds.CompositeChainlinkPF,
				Client:       client,
			},
			Details: map[string]interface{}{
				"oracle": oracle,
				"token":  token,
			},
			// since last_sync is set to discoveredAt not discoveredAt-1, setPrice will get tokenBase and baseUSD price at discoveredAt
			// so the db entry that is added at addPriceToDB will have the correct price while creating new compositeChainlinkPF
			LastSync: discoveredAt,
			V:        version,
		},
		Repo: repo,
	}
	compositeMdl := NewCompositeChainlinkPFFromAdapter(adapter)
	compositeMdl.addPriceToDB(discoveredAt)
	return compositeMdl
}
func (mdl *CompositeChainlinkPF) addSecAddrs(blockNum int64) {
	if mdl.Details["decimals"] != nil {
		return
	}
	oracleAddr := common.HexToAddress(mdl.Address)
	tokenETHPF := getAddrFromRPC(mdl.Client, "targetETH", oracleAddr, blockNum)
	// get decimals
	decimalsToBasePF := getDecimals(mdl.Client, tokenETHPF, blockNum)
	mdl.Details["decimals"] = decimalsToBasePF
	if mdl.Details["secAddrs"] != nil {
		return
	}
	mainAgg := cpf.NewMainAgg(mdl.Client, tokenETHPF)

	ethUSDPF := getAddrFromRPC(mdl.Client, "ETHUSD", oracleAddr, blockNum)
	baseTokenMainAgg := cpf.NewMainAgg(mdl.Client, ethUSDPF)
	//
	mainPhaseAgg := mainAgg.GetPriceFeedAddr(blockNum)
	basePhaseAgg := baseTokenMainAgg.GetPriceFeedAddr(blockNum)
	mdl.Details["secAddrs"] = map[string]interface{}{
		"target":      tokenETHPF.Hex(),
		"base":        ethUSDPF.Hex(),
		"targetPhase": mainPhaseAgg.Hex(),
		"basePhase":   basePhaseAgg.Hex(),
	}
}

func toDecimals(decimals interface{}) int8 {
	switch ans := decimals.(type) {
	case float64:
		return int8(ans)
	case int8:
		return ans
	default:
		log.Fatal("Can't convert decimals interface")
		return 0
	}
}

func NewCompositeChainlinkPFFromAdapter(adapter *ds.SyncAdapter) *CompositeChainlinkPF {
	compositeMdl := &CompositeChainlinkPF{
		SyncAdapter: adapter,
	}
	compositeMdl.addSecAddrs(utils.Max(adapter.LastSync, adapter.DiscoveredAt))
	compositeMdl.decimalsOfBasePF = toDecimals(adapter.GetDetails()["decimals"])
	compositeMdl.BaseTokenMainAgg = cpf.NewMainAgg(adapter.Client, compositeMdl.getAddrFromDetails("base"))
	compositeMdl.MainAgg = cpf.NewMainAgg(adapter.Client, compositeMdl.getAddrFromDetails("target"))
	compositeMdl.DataProcessType = ds.ViaMultipleLogs
	compositeMdl.setPrices(adapter.LastSync)
	//
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
	log.Fatalf("feed(%s) can't fetch price.", feed)
	return nil
}

func getAddrFromRPC(client core.ClientI, targetMethod string, oracle common.Address, blockNum int64) common.Address {
	chainId, err := client.ChainID(context.TODO())
	log.CheckFatal(err)
	sig := getSig(targetMethod, blockNum, chainId.Int64())
	tokenETHPFData, err := core.CallFuncGetSingleValue(client, sig, oracle, blockNum, nil)
	if err != nil {
		log.Fatalf("Oracle(%s) at blockNum %d doesn't have valid %s: %s", oracle, blockNum, targetMethod, err)
	}
	return common.BytesToAddress(tokenETHPFData)
}
func getDecimals(client core.ClientI, addr common.Address, blockNum int64) int8 {
	decimals, err := core.CallFuncGetSingleValue(client, "313ce567", addr, blockNum, nil) // decimals
	if err != nil {
		log.Fatalf("Can't get decimals for addr(%s) : %s", addr, err)
	}
	return int8(new(big.Int).SetBytes(decimals).Int64())
}

func (mdl *CompositeChainlinkPF) AfterSyncHook(syncedTill int64) {
	log.Infof("Processed event:%v, added price : %d", mdl.ansBlock, mdl.priceAdded)
	mdl.ansBlock = nil
	mdl.priceAdded = 0
	mdl.SyncAdapter.AfterSyncHook(syncedTill)
}

// there are two type of composite oracle
// 1) with targetETh and ethUSD price feed.
// 2) with baseToUSD and targetToBase price feed.
func getSig(targetMethod string, discoveredAt int64, chainId int64) (sig string) {
	// TODO anvil fork
	compositeOracleVersion := 0
	baseNet := log.GetBaseNet(chainId)
	switch baseNet {
	case "MAINNET":
		if discoveredAt <= 15997386 {
			compositeOracleVersion = 2
		} else if discoveredAt <= 18544086 {
			compositeOracleVersion = 210
		} else {
			compositeOracleVersion = 3
		}
	case "ARBITRUM", "OPTIMISM":
		compositeOracleVersion = 3
	}
	//
	switch targetMethod {
	case "targetETH":
		switch compositeOracleVersion {
		case 2:
			sig = "f1a75c6e" // targetEthPriceFeed
		case 210:
			sig = "a76d5447" // targetToBasePriceFeed
		case 3:
			sig = "385aee1b" // priceFeed0
		default:
			log.Fatal("Unknown composite oracle version", compositeOracleVersion, targetMethod, discoveredAt, chainId)
		}
	case "ETHUSD":
		switch compositeOracleVersion {
		case 2:
			sig = "42f6fb29" // ethUsdPriceFeed
		case 210:
			sig = "51a799d6" // baseToUsdPriceFeed
		case 3:
			sig = "ab0ca0e1" // priceFeed1
		default:
			log.Fatal("Unknown composite oracle version", compositeOracleVersion, targetMethod, discoveredAt, chainId)
		}
	default:
		log.Fatal(targetMethod, "not found")
	}
	return
}
