package rebase_token

import (
	"context"
	"math/big"
	"testing"

	"github.com/Gearbox-protocol/sdk-go/artifacts/rebaseToken"
	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func TestRebaseToken(t *testing.T) {
	provider := utils.GetEnvOrDefault("ETH_PROVIDER", "")
	if provider == "" {
		t.Skip()
	}
	log.SetTestLogging(t)
	client := ethclient.NewEthClient(&config.Config{EthProvider: provider})
	stETh := core.GetSymToAddrByChainId(core.GetChainId(client)).Tokens["stETH"]
	//
	//
	// blockNum := big.NewInt(17293521)
	blockNum := big.NewInt(17289199)
	v, err := client.StorageAt(context.TODO(), stETh, common.HexToHash(totalSharesKey), blockNum)
	log.CheckFatal(err)
	log.Info(new(big.Int).SetBytes(v))
	v, err = client.StorageAt(context.TODO(), stETh, common.HexToHash(depositedBalanceKey), blockNum)
	log.CheckFatal(err)
	log.Info("deposited", new(big.Int).SetBytes(v))
	v, err = client.StorageAt(context.TODO(), stETh, common.HexToHash(clBalanceKey), blockNum)
	log.CheckFatal(err)
	log.Info("cl", new(big.Int).SetBytes(v))
	contract, err := rebaseToken.NewRebaseToken(stETh, client)
	log.CheckFatal(err)
	totalETh, err := contract.GetTotalPooledEther(&bind.CallOpts{BlockNumber: blockNum})
	log.CheckFatal(err)
	log.Info(totalETh)
	eth, err := contract.GetPooledEthByShares(&bind.CallOpts{BlockNumber: blockNum}, utils.StringToInt("459223736079357852562"))
	log.CheckFatal(err)
	log.Info(eth)
}
