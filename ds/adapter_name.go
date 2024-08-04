package ds

import (
	"log"
	"strings"
)

const (
	AddressProvider    = "AddressProvider"
	ContractRegister   = "ContractRegister"
	PriceOracle        = "PriceOracle"
	AccountFactory     = "AccountFactory"
	ACL                = "ACL"
	CreditManager      = "CreditManager"
	Pool               = "Pool"
	ChainlinkPriceFeed = "ChainlinkPriceFeed"
	QueryPriceFeed     = "QueryPriceFeed"
	CreditFilter       = "CreditFilter"
	GearToken          = "GearToken"
	Treasury           = "Treasury"
	AccountManager     = "AccountManager"
	CreditConfigurator = "CreditConfigurator"
	LMRewardsv2        = "LMRewardsv2"
	//
	RebaseToken = "RebaseToken"
	// Wrapper
	AggregatedQueryFeedWrapper = "AggregatedQueryFeedWrapper"
	AdminWrapper               = "AdminWrapper"
	CFWrapper                  = "CFWrapper"
	CMWrapper                  = "CMWrapper"
	PoolWrapper                = "PoolWrapper"
	PoolQuotaWrapper           = "PoolQuotaWrapper"
	// v3
	PoolQuotaKeeper = "PoolKeeper"
	LMRewardsv3     = "LMRewardsv3"
)

// beef2 -- v2LMRewards
// beef3 -- v3LMRewards

func IsWrapperAdapter(name string) bool {
	return strings.HasSuffix(name, "Wrapper")
}

const (
	UnknownPF            = "UnknownPF"
	YearnPF              = "YearnPF"
	SingleAssetPF        = "SingleAssetPF"
	CurvePF              = "CurvePF"
	RedStonePF           = "RedStonePF"
	ZeroPF               = "ZeroPF"
	CompositeChainlinkPF = "CompositeChainlinkPF"
	CompositeRedStonePF  = "CompositeRedStonePF"
	AlmostZeroPF         = "AlmostZeroPF"
)

const (
	FacadeMulticallCall        = "FacadeMulticall"
	FacadeBotMulticallCall        = "FacadeBotMulticall"
	FacadeOpenMulticallCall    = "FacadeOpenMulticall"
	FacadeLiquidateCall        = "FacadeLiquidate"
	FacadeLiquidateExpiredCall = "FacadeLiquidateExpired"
	FacadeCloseAccountCall     = "FacadeCloseAccount"
)

func FacadeAccountMethodSigToCallName(funcSig string) string {
	switch funcSig {
	// common v2/v3
	case "liquidateCreditAccount":
		return FacadeLiquidateCall
	case "closeCreditAccount":
		return FacadeCloseAccountCall
	case "multicall":
		return FacadeMulticallCall
	case "botMulticall":
		return FacadeBotMulticallCall
	// v2
	case "openCreditAccountMulticall":
		return FacadeOpenMulticallCall
	case "liquidateExpiredCreditAccount":
		return FacadeLiquidateExpiredCall
	// for v3
	case "openCreditAccount":
		return FacadeOpenMulticallCall
	}
	log.Fatal()
	return ""
}
