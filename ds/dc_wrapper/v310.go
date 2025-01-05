package dc_wrapper

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/artifacts/creditAccountCompressor"
	"github.com/Gearbox-protocol/sdk-go/artifacts/multicall"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/pkg/dc"
	"github.com/ethereum/go-ethereum/common"
)

func (dcw *DataCompressorWrapper) addFieldsToAccountv310(blockNum int64, callData creditAccountCompressor.CreditAccountData) (dc.CreditAccountCallData, error) {
	poolQuotaBytes, err := core.CallFuncGetSingleValue(dcw.client, "be8da14b", callData.CreditManager, blockNum, nil) // poolQuotaKeeper
	log.CheckFatal(err)
	poolKeeperAddr := common.BytesToAddress(poolQuotaBytes)
	//
	poolQuotaABI := core.GetAbi("PoolQuotaKeeper")
	// indexes
	var tokenQuotaIndexcalls []multicall.Multicall2Call
	var tokens []common.Address
	for _, token := range callData.Tokens {
		if token.Quota != nil && token.Quota.Cmp(new(big.Int)) > 0 {
			getQuotadata, err := poolQuotaABI.Pack("getQuota", callData.CreditAccount, token.Token)
			log.CheckFatal(err)
			tokens = append(tokens, token.Token)
			tokenQuotaIndexcalls = append(tokenQuotaIndexcalls, multicall.Multicall2Call{
				Target:   poolKeeperAddr,
				CallData: getQuotadata,
			})
		}
	}
	// cm base index
	cmABI := core.GetAbi("CreditManagerv3")
	data, err := cmABI.Pack("creditAccountInfo", callData.CreditAccount)
	log.CheckFatal(err)
	tokenQuotaIndexcalls = append(tokenQuotaIndexcalls, multicall.Multicall2Call{
		Target:   callData.CreditManager,
		CallData: data,
	})
	//
	results := core.MakeMultiCall(dcw.client, blockNum, false, tokenQuotaIndexcalls)
	// indexes
	tokenQuotaIndex := map[string]*big.Int{}
	for ind, token := range tokens {
		if results[ind].Success {
			tokenQuotaIndex[token.Hex()] = new(big.Int).SetBytes(results[ind].ReturnData[32:])
		}
	}

	// creditAccountInfo
	var cumulativeIndexLastUpdate, cumulativeQuotaInterest *big.Int
	baseCumIndexInfo := results[len(tokens)]
	if baseCumIndexInfo.Success {
		values, err := cmABI.Unpack("creditAccountInfo", baseCumIndexInfo.ReturnData)
		log.CheckFatal(err)
		cumulativeIndexLastUpdate = values[1].(*big.Int)
		cumulativeQuotaInterest = values[2].(*big.Int)
	}
	// return
	return dc.GetAccountDataFromDCCall(dcw.client, callData.CreditFacade, blockNum, dc.CreditAccountv310{
		CreditAccountData:         callData,
		CumulativeIndexLastUpdate: cumulativeIndexLastUpdate,
		CumulativeQuotaInterest:   cumulativeQuotaInterest,
		QuotaCumIndexMap:          tokenQuotaIndex,
	})
}
