package main

import (
	"math/big"
	"strings"

	"github.com/Gearbox-protocol/sdk-go/core"
	"github.com/Gearbox-protocol/sdk-go/core/schemas"
	"github.com/Gearbox-protocol/sdk-go/ethclient"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/ethereum/go-ethereum/common"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	log.CheckFatal(err)
	client, err := ethclient.Dial(utils.GetEnvOrDefault("ETH_PROVIDER", ""))
	log.CheckFatal(err)
	cfg := config.Config{DatabaseUrl: utils.GetEnvOrDefault("DATABASE_URL", "")}
	db := repository.NewDBClient(&cfg)

	query := `SELECT id, css.session_id, block_num , css.balances FROM
	credit_session_snapshots css WHERE css.balances ?? (SELECT address FROM tokens where symbol='stETH');`
	css := []schemas.CreditSessionSnapshot{}
	err = db.Raw(query).Find(&css).Error
	stETH := core.GetSymToAddrByChainId(core.GetChainId(client)).Tokens["stETH"]
	log.CheckFatal(err)

	tx := db.Begin()
	for _, snapshot := range css {
		account := strings.Split(snapshot.SessionId, "_")[0]
		accountData := common.HexToHash(account)
		_v, err := core.CallFuncWithExtraBytes(
			client, "f5eb42dc", // shareOf, https://etherscan.io/token/0xae7ab96520de3a18e5e111b5eaab095312d7fe84#readProxyContract
			stETH, snapshot.BlockNum, accountData[:],
		)
		log.CheckFatal(err)
		amt := new(big.Int).SetBytes(_v)
		//
		(*snapshot.Balances)[core.NULL_ADDR.Hex()] = core.CoreIntBalance{
			IsAllowed: false,
			IsEnabled: false,
			BI:        (*core.BigInt)(amt),
			F:         utils.GetFloat64Decimal(amt, 18),
			Ind:       -1,
		}
		err = tx.Raw("UPDATE credit_session_snapshots set balances=? WHERE id = ?", snapshot.Balances, snapshot.ID).Error
		log.CheckFatal(err)
	}
	//
	query2 := `UPDATE credit_sessions cs set cs.balances= css.balances FROM (SELECT DISTINCT ON (session_id) balances, session_id from credit_session_snapshots ORDER BY session_id, block_num DESC) css WHERE cs.id= css.session_id`
	log.CheckFatal(tx.Raw(query2).Error)
	//
	info := tx.Commit()
	log.CheckFatal(info.Error)
}
