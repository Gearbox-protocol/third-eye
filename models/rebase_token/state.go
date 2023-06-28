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
		DepositBalance:    (*core.BigInt)(mdl.getStorageAt(blockNum, "0xed310af23f61f96daefbcd140b306c0bdbf8c178398299741687b90e794772b0")),
		DepositValidators: mdl.getStorageAt(blockNum, "0xe6e35175eb53fc006520a2a9c3e9711a7c00de6ff2c32dd31df8c5a24cac1b5c").Int64(),
		CLBalance:         (*core.BigInt)(mdl.getStorageAt(blockNum, "0xa66d35f054e68143c18f32c990ed5cb972bb68a68f500cd2dd3a16bbf3686483")),
		CLValidators:      mdl.getStorageAt(blockNum, "0x9f70001d82b6ef54e9d3725b46581c3eb9ee3aa02b941b6aa54d678a9ca35b10").Int64(),
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
