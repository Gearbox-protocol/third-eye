package healthcheck

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/version"
	"go.uber.org/fx"
)

type healthResp struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

func newHealthcheckEndpoint(lc fx.Lifecycle, config *config.Config) {
	if config.Port == "0" {
		return
	}

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Port),
		Handler: mux,
	}
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(healthResp{"OK", version.Version})
	})

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			log.Infof("healthcheck endpoint up at :%s", config.Port)
			go server.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("healthcheck endpoint down")
			return server.Shutdown(ctx)
		},
	})
}
