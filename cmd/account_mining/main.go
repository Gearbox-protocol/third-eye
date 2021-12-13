package main

import (
	"context"
	"flag"
	"math/big"
	"time"

	// "github.com/Gearbox-protocol/third-eye/artifacts/accountMining"
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

type BlockMiningDetails struct {
	Count                  int
	TotalEthSpentOnAccount *big.Int
	BaseFee                *big.Int
}

type AccountMining struct {
	Address string
	*core.Node
	TotalCount     int64
	finished        bool
	CurrentBlockNum int64
	// contractETH     *accountMining.AccountMining
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
	if am.CurrentBlockNum == 0 {
		return
	}
	details := am.blockDetails[am.CurrentBlockNum]
	eipBaseFee := utils.GetFloat64Decimal(details.BaseFee, 9)
	avgAccountMiningPrice := utils.GetFloat64Decimal(details.TotalEthSpentOnAccount, 18)
	avgAccountMiningPrice = avgAccountMiningPrice/float64(details.Count)
	log.Msgf("Block[%d]: mined %d, total %d of 5000 accounts, avg price per account: %f ETH, eip 1559 baseFee is %f",
		am.CurrentBlockNum, details.Count, am.TotalCount, avgAccountMiningPrice, eipBaseFee)

}
func (am *AccountMining) Sync(startNum int64) {
	latestBlockNum := am.GetLatestBlockNumber()
	txLogs, err := am.GetLogs(startNum, latestBlockNum, am.Address)
	log.CheckFatal(err)
	for _, txLog := range txLogs {
		if am.CurrentBlockNum != 0 && am.CurrentBlockNum != int64(txLog.BlockNumber) {
			am.Send()
		}
		am.CurrentBlockNum = int64(txLog.BlockNumber)
		switch txLog.Topics[0] {
		case core.Topic("Claimed(uint256,address)"):
			am.ProcessLog(&txLog)
		}
	}
}


func isMiningStarted(gearToken string, wss string) {
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
			log.Info(vlog)
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
			go func(){
			var amAddr, gearToken, wss string
			var miningStarted bool
			flag.StringVar(&amAddr, "miningAddr", "", "Account mining address")
			flag.StringVar(&gearToken, "gearToken", "", "Gear token address")
			flag.StringVar(&wss, "wss", "", "Web socket url")
			flag.BoolVar(&miningStarted, "started", false, "Web socket url")
			flag.Parse()
			if !miningStarted {
				isMiningStarted(gearToken, wss)
			}
			am := AccountMining{
				Address: amAddr,
				Node: &core.Node{
					Client: client,
				},
				blockDetails: make(map[int64]*BlockMiningDetails),
			}
			am.init()
			var startNum int64 =0
			for !am.IsFinished() {
				am.Sync(startNum)
				am.Send()
				startNum = am.CurrentBlockNum+1
				am.CurrentBlockNum = 0
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
