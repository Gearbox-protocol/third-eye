package rebase_token

import (
	"context"
	"math/big"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (mdl RebaseToken) getStateAt(blockNum int64) *schemas.RebaseTokenDetails {
	opts := &bind.CallOpts{
		BlockNumber: big.NewInt(blockNum),
	}
	totalShares, err := mdl.contract.GetTotalShares(opts)
	log.CheckFatal(err)
	poolEth, err := mdl.contract.GetTotalPooledEther(opts)
	log.CheckFatal(err)

	calc := &schemas.RebaseTokenDetails{
		RebaseDetailsForDB: schemas.RebaseDetailsForDB{
			TotalShares: (*core.BigInt)(totalShares),
		},
		DepositBalance:    (*core.BigInt)(mdl.getStorageAt(blockNum, depositedBalanceKey)), // users balances
		DepositValidators: mdl.getStorageAt(blockNum, depositedValidatorKey).Int64(),       // used in validators
		CLBalance:         (*core.BigInt)(mdl.getStorageAt(blockNum, clBalanceKey)),        // waiting to be used in validators
		CLValidators:      mdl.getStorageAt(blockNum, validatorsKey).Int64(),               // new validators in wait
	}
	if calc.GetPostTotalEther().Cmp(poolEth) != 0 {
		log.Fatalf("poolTotal Eth(%d) calculated is different from call fetched %d", calc.GetPostTotalEther(), poolEth)
	}
	return calc
}
func (mdl RebaseToken) getStorageAt(blockNum int64, key string) *big.Int {
	return GetStorageAt(blockNum, key, mdl.Client, mdl.Address)
}
func GetStorageAt(blockNum int64, key string, client core.ClientI, addr string) *big.Int {
	if key[:2] == "0x" {
		key = key[2:]
	}
	value, err := client.StorageAt(context.Background(), common.HexToAddress(addr), common.HexToHash(key), big.NewInt(blockNum))
	if err != nil {
		log.Fatalf("Value %d err %v", value, err)
	}
	return new(big.Int).SetBytes(value)
}
