package main

import (
	"context"
	"math/big"
	"time"

	"github.com/Gearbox-protocol/third-eye/artifacts/addressProvider"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/services"
	"github.com/Gearbox-protocol/third-eye/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/fx"
)

type BlockMiningDetails struct {
	Count                  int
	TotalEthSpentOnAccount *big.Int
	BaseFee                *big.Int
}

type AccountMining struct {
	Address string
	*core.Node
	TotalCount      int64
	finished        bool
	CurrentBlockNum int64
	canSend         bool
	// contractETH     *accountMining.AccountMining
	blockDetails map[int64]*BlockMiningDetails
	txProcessed  map[string]bool
}

func (am *AccountMining) IsFinished() bool {
	return am.finished
}
func (am *AccountMining) BaseFee(blockNum int64) *big.Int {
	if am.blockDetails[blockNum] == nil {
		am.blockDetails[blockNum] = &BlockMiningDetails{}
	}
	if am.blockDetails[blockNum].BaseFee == nil {
		am.blockDetails[blockNum].BaseFee = am.GetHeader(blockNum).BaseFee
	}
	return am.blockDetails[blockNum].BaseFee
}

func (am *AccountMining) ProcessLog(txLog *types.Log) {
	blockNum := int64(txLog.BlockNumber)
	if am.blockDetails[blockNum] == nil {
		am.blockDetails[blockNum] = &BlockMiningDetails{}
	}
	am.TotalCount++
	am.blockDetails[blockNum].Count++
	// data, err := am.contractETH.ParseClaimed(*txLog)
	// log.CheckFatal(err)
	if am.txProcessed[txLog.TxHash.Hex()] {
		am.txProcessed[txLog.TxHash.Hex()] = true
		return
	}
	ethUsed := am.EthUsed(txLog.TxHash, am.BaseFee(int64(txLog.BlockNumber)))
	if am.blockDetails[blockNum].TotalEthSpentOnAccount != nil {
		am.blockDetails[blockNum].TotalEthSpentOnAccount = new(big.Int).Add(ethUsed,
			am.blockDetails[blockNum].TotalEthSpentOnAccount)
	} else {
		am.blockDetails[blockNum].TotalEthSpentOnAccount = ethUsed
	}
}
func (am *AccountMining) init() {
	// contract, err := accountMining.NewAccountMining(common.HexToAddress(am.Address), am.Client)
	// log.CheckFatal(err)
	// am.contractETH = contract
}
func (am *AccountMining) Send() {
	if !am.canSend {
		return
	}
	details := am.blockDetails[am.CurrentBlockNum]
	eipBaseFee := utils.GetFloat64Decimal(details.BaseFee, 9)
	avgAccountMiningPrice := utils.GetFloat64Decimal(details.TotalEthSpentOnAccount, 18)
	avgAccountMiningPrice = avgAccountMiningPrice / float64(details.Count)
	log.Msgf("Block[%d]: mined %d, total %d of 5000 accounts, avg price per account: %f ETH, eip 1559 baseFee is %f",
		am.CurrentBlockNum, details.Count, am.TotalCount, avgAccountMiningPrice, eipBaseFee)
	am.canSend = false
}
func (am *AccountMining) Sync(startNum, latestBlockNum int64) {
	txLogs, err := am.GetLogs(startNum, latestBlockNum, am.Address)
	log.CheckFatal(err)
	for _, txLog := range txLogs {
		if am.CurrentBlockNum != int64(txLog.BlockNumber) {
			am.Send()
		}
		am.canSend = true
		am.CurrentBlockNum = int64(txLog.BlockNumber)
		switch txLog.Topics[0] {
		case core.Topic("Claimed(uint256,address)"):
			am.ProcessLog(&txLog)
		}
	}
}

func (am *AccountMining) isMiningStarted(gearToken string) {
	for {
		latestBlockNum := am.GetLatestBlockNumber()
		logs, err := am.GetLogs(0, latestBlockNum, gearToken)
		log.CheckFatal(err)
		for _, txLog := range logs {
			if txLog.Topics[0] == core.Topic("MinerSet(address)") {
				if common.HexToAddress(txLog.Topics[1].Hex()).Hex() == am.Address {
					log.Msgf("Mining Started at %d", txLog.BlockNumber)
					return
				}
			}
		}
		time.Sleep(time.Second * 30)
	}
}

func StartServer(lc fx.Lifecycle, client *ethclient.Client, config *config.Config) {

	// Starting server
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				contractETH, err := addressProvider.NewAddressProvider(common.HexToAddress(config.AddressProviderAddress), client)
				log.CheckFatal(err)
				gearToken, err := contractETH.GetGearToken(&bind.CallOpts{})
				log.CheckFatal(err)
				am := AccountMining{
					Address: config.MiningAddr,
					Node: &core.Node{
						Client: client,
					},
					blockDetails: make(map[int64]*BlockMiningDetails),
				}
				am.isMiningStarted(gearToken.Hex())
				am.init()
				var startNum int64 = 0
				for !am.IsFinished() {
					latestBlockNum := am.GetLatestBlockNumber()
					am.Sync(startNum, latestBlockNum-2)
					am.Send()
					startNum = latestBlockNum - 1
					if am.TotalCount == 5000 {
						am.finished = true
						log.Msg("Mining finished")
					}
					time.Sleep(time.Second * 30)
				}
			}()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		services.Module,
		config.Module,
		ethclient.Module,
		fx.NopLogger,
		fx.Invoke(StartServer),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
