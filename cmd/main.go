/*
* Gearbox monitoring
* Copyright (c) 2021. Mikael Lazarev
*
 */

package main

import (
	"context"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/core"
	"github.com/Gearbox-protocol/third-eye/engine"
	"github.com/Gearbox-protocol/third-eye/ethclient"
	"github.com/Gearbox-protocol/third-eye/log"
	"github.com/Gearbox-protocol/third-eye/repository"
	"github.com/Gearbox-protocol/third-eye/services"
	_ "github.com/heroku/x/hmetrics/onload"
	"go.uber.org/fx"
	"time"
)

func StartServer(lc fx.Lifecycle, engine core.EngineI, config *config.Config) {

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
		fx.Invoke(StartServer),
	)
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	<-app.Done()
}
