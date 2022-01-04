package main

import (
	"context"

	"flag"
	"github.com/Gearbox-protocol/third-eye/config"
	"time"

	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"go.uber.org/fx"
)

func StartServer(lc fx.Lifecycle, d *DBhandler, shutdowner fx.Shutdowner) {

	// Starting server
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			var blockNum int64
			flag.Int64Var(&blockNum, "blockNum", 0, "blockNUm to rollback to")
			go func() {
				flag.Parse()
				// credit_session -> credit_manager -> pool -> sync_adapters
				// debts
				d.delete("debt_sync", "last_calculated_at", blockNum)
				d.deleteOnBlockNum("debts", blockNum)
				d.deleteOnBlockNum("iquidable_accounts", blockNum)
				// tokens/credit filter
				d.deleteOnBlockNum("allowed_protocols", blockNum)
				d.deleteOnBlockNum("allowed_tokens", blockNum)
				d.deleteOnBlockNum("price_feeds", blockNum)
				d.deleteOnBlockNum("token_oracle", blockNum)
				// credit session
				d.deleteOnBlockNum("account_operations", blockNum)
				d.deleteOnBlockNum("credit_session_snapshots", blockNum)
				d.CreditSessionRollback(blockNum)
				// credit_manager
				d.deleteOnBlockNum("credit_manager_stats", blockNum)
				d.CreditManagerRollback(blockNum)
				// pools
				d.PoolRollback(blockNum)
				// blocks/sync adapter
				d.delete("blocks", "id", blockNum)
				d.SyncadapterRollback(blockNum)
				// shutdown
				shutdowner.Shutdown()
			}()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		config.Module,
		fx.Provide(
			NewDBhandler,
			repository.NewDBClient,
		),
		fx.Invoke(StartServer),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	<-app.Done()
}
