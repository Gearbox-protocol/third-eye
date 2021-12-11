package main

import (
	"context"
	"flag"
	"math/big"
	"time"

	"github.com/Gearbox-protocol/third-eye/artifacts/accountMining"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/services"
	"github.com/Gearbox-protocol/third-eye/utils"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"go.uber.org/fx"
)

var accountMiningAddr string

type BlockMiningDetails struct {
	Count                  int
	TotalEthSpentOnAccount *big.Int
	BaseFee                *big.Int
}

type AccountMining struct {
	Address string
	*core.Node
	latestIndex     int64
	finished        bool
	CurrentBlockNum int64
	firstSkipped    bool
	contractETH     *accountMining.AccountMining
	blockDetails    map[int64]*BlockMiningDetails
	txProcessed     map[string]bool
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

	am.blockDetails[blockNum].Count++
	data, err := am.contractETH.ParseClaimed(*txLog)
	log.CheckFatal(err)
	am.latestIndex = data.Index.Int64()
	if am.txProcessed[txLog.TxHash.Hex()] {
		am.txProcessed[txLog.TxHash.Hex()] = true
		return
	}
	if am.blockDetails[blockNum] == nil {
		am.blockDetails[blockNum] = &BlockMiningDetails{}
	}
	ethUsed := am.EthUsed(txLog.TxHash, am.BaseFee(int64(txLog.BlockNumber)))
	am.blockDetails[blockNum].TotalEthSpentOnAccount = new(big.Int).Add(ethUsed,
		am.blockDetails[blockNum].TotalEthSpentOnAccount)
	if am.latestIndex == 5000 {
		am.Send()
		am.finished = true
		log.Info("Mining finished")
	}
}
func (am *AccountMining) init() {
	contract, err := accountMining.NewAccountMining(common.HexToAddress(am.Address), am.Client)
	log.CheckFatal(err)
	am.contractETH = contract
}
func (am *AccountMining) Send() {
	details := am.blockDetails[am.CurrentBlockNum]
	eipBaseFee := utils.GetFloat64Decimal(details.BaseFee, 9)
	avgAccountMiningPrice := utils.GetFloat64Decimal(details.TotalEthSpentOnAccount, 18)
	log.Msgf("Block[%d]: mined %d, last index:%d of 5000 accounts, avg price per account: %f ETH, eip 1559 baseFee is %f",
		am.CurrentBlockNum, details.Count, am.latestIndex, avgAccountMiningPrice, eipBaseFee)

}
func (am *AccountMining) Sync() {
	latestBlockNum := am.GetLatestBlockNumber()
	txLogs, err := am.GetLogs(am.CurrentBlockNum+1, latestBlockNum, accountMiningAddr)
	log.CheckFatal(err)
	for _, txLog := range txLogs {
		if am.firstSkipped && am.CurrentBlockNum != int64(txLog.BlockNumber) {
			am.Send()
		}
		am.firstSkipped = true
		am.CurrentBlockNum = int64(txLog.BlockNumber)
		switch txLog.Topics[0] {
		case core.Topic("Claimed(uint256,address"):
			am.ProcessLog(&txLog)
		}
	}
}

func isMiningStart(gearToken string, wss string) {
	client, err := ethclient.Dial(wss)
	log.CheckFatal(err)
	contractAddress := common.HexToAddress(gearToken)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	ch := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, ch)
	log.CheckFatal(err)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vlog := <-ch:
			if vlog.Topics[0] == core.Topic("MinerSet(address)") {
				log.Msg("Mining Started")
				close(ch)
				return
			}
		}
	}
}

func StartServer(lc fx.Lifecycle, client *ethclient.Client, config *config.Config) {

	// Starting server
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			var amAddr, gearToken, wss string
			var startBlock int64
			flag.StringVar(&amAddr, "miningAddr", "", "Account mining address")
			flag.StringVar(&gearToken, "gearToken", "", "Gear token address")
			flag.StringVar(&wss, "gearToken", "", "Web socket url")
			flag.Int64Var(&startBlock, "block", 0, "Account mining address")
			flag.Parse()
			isMiningStart(gearToken, wss)
			am := AccountMining{
				Address: amAddr,
				Node: &core.Node{
					Client: client,
				},
				blockDetails: make(map[int64]*BlockMiningDetails),
			}
			am.init()
			for am.IsFinished() {
				am.Sync()
				time.Sleep(time.Second * 30)
			}
			return nil
		},
	})
}

func main() {
	app := fx.New(
		services.Module,
		config.Module,
		fx.Invoke(StartServer),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
