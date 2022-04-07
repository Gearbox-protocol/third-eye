/*
* Gearbox monitoring
* Copyright (c) 2021. Harsh Jain
*
 */

package main

import (
	"context"
	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/services"
	"go.uber.org/fx"
	"time"
)

func StartServer(lc fx.Lifecycle, debtEng ds.DebtEngineI, shutdowner fx.Shutdowner) {

	// Starting server
	lc.Append(fx.Hook{
		// To mitigate the impact of deadlocks in application startup and
		// shutdown, Fx imposes a time limit on OnStart and OnStop hooks. By
		// default, hooks have a total of 15 seconds to complete. Timeouts are
		// passed via Go's usual context.Context.
		OnStart: func(context.Context) error {
			// In production, we'd want to separate the Listen and Serve phases for
			// better error-handling.
			go func() {
				debtEng.ProcessBackLogs()
				shutdowner.Shutdown()
			}()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		ethclient.Module,
		config.Module,
		repository.Module,
		services.Module,
		fx.NopLogger,
		debts.Module,
		fx.Invoke(StartServer),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}
	<-app.Done()
}
