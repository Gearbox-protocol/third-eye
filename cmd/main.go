/*
* Gearbox monitoring
* Copyright (c) 2021. Harsh Jain
*
 */

package main

import (
	"context"
	_ "net/http/pprof"
	"time"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/debts"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/healthcheck"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/services"
	_ "github.com/heroku/x/hmetrics/onload"
	"go.uber.org/fx"
)

// func init() {
// 	go func() {
// 		http.ListenAndServe(":8080", nil)
// 	}()
// }

func StartServer(lc fx.Lifecycle, engine ds.EngineI, config *config.Config) {
	log.NewAMQPService(config.ChainId, config.AMPQEnable, config.AMPQUrl, "Third-eye")
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
				engine.UseThreads()
				engine.SyncHandler()
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
		engine.Module,
		fx.NopLogger,
		debts.Module,
		healthcheck.Module,
		fx.Invoke(StartServer),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
