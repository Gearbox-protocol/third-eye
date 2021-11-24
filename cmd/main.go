/*
* Gearbox monitoring
* Copyright (c) 2021. Mikael Lazarev
*
*/

package main

import (
	"github.com/Gearbox-protocol/gearscan/config"
	"github.com/Gearbox-protocol/gearscan/ethclient"
	"github.com/Gearbox-protocol/gearscan/repository"
	"github.com/Gearbox-protocol/gearscan/engine" 
	"github.com/Gearbox-protocol/gearscan/models" 
	"github.com/Gearbox-protocol/gearscan/log"
	"github.com/Gearbox-protocol/gearscan/utils"
	"go.uber.org/fx"

	"context"
	"time" 
)
 
func StartServer(lc fx.Lifecycle, engine engine.EngineI) {

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
				log.Info("harsh")
				engine.Sync()	
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
		utils.Module,
		engine.Module,
		models.Module,
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
 