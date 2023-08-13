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
	PoolLMRewards      = "PoolLMRewards"
	//
	RebaseToken = "RebaseToken"
	// Wrapper
	AggregatedQueryFeedWrapper = "AggregatedQueryFeedWrapper"
	AdminWrapper               = "AdminWrapper"
	CFWrapper                  = "CFWrapper"
	CMWrapper                  = "CMWrapper"
	PoolWrapper                = "PoolWrapper"
)

func IsWrapperAdapter(name string) bool {
	return strings.HasSuffix(name, "Wrapper")
}

const (
	UnknownPF            = "UnknownPF"
	YearnPF              = "YearnPF"
	CurvePF              = "CurvePF"
	ZeroPF               = "ZeroPF"
	CompositeChainlinkPF = "CompositeChainlinkPF"
	AlmostZeroPF         = "AlmostZeroPF"
)

const (
	FacadeMulticallCall        = "FacadeMulticall"
	FacadeOpenMulticallCall    = "FacadeOpenMulticall"
	FacadeLiquidateCall        = "FacadeLiquidate"
	FacadeLiquidateExpiredCall = "FacadeLiquidateExpired"
	FacadeCloseAccountCall     = "FacadeCloseAccount"
)

func FacadeAccountMethodSigToCallName(funcSig string) string {
	switch funcSig {
	// common
	case "liquidateCreditAccount":
		return FacadeLiquidateCall
	case "closeCreditAccount":
		return FacadeCloseAccountCall
	case "multicall":
		return FacadeMulticallCall
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
