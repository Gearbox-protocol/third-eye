package aggregated_block_feed

import (
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/common"
)

func GetCurveVirtualPrice(blockNum int64, oracleAddr common.Address, version core.VersionType, client core.ClientI) *big.Int {
	curvePool := func() common.Address {
		if !version.MoreThanEq(core.NewVersion(300)) {
			curvePoolBytes, err := core.CallFuncWithExtraBytes(client, "218751b2", oracleAddr, blockNum, nil) // curvePool from curvev1Adapter abi
			log.CheckFatal(err)
			return common.BytesToAddress(curvePoolBytes)
		} else {
			// LPCONTRACT_LOGIC
			lpCOntractBytes, err := core.CallFuncWithExtraBytes(client, "8acee3cf", oracleAddr, blockNum, nil) // lpContract
			log.CheckFatal(err)
			return common.BytesToAddress(lpCOntractBytes)
		}
	}()
	virtualPrice, err := core.CallFuncWithExtraBytes(client, "bb7b8b80", curvePool, blockNum, nil) // getVirtualPrice
	log.CheckFatal(err)
	return new(big.Int).SetBytes(virtualPrice)
}
