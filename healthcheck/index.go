package healthcheck

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Gearbox-protocol/sdk-go/log"
	"github.com/Gearbox-protocol/third-eye/config"
	"go.uber.org/fx"
)

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
		w.WriteHeader(http.StatusOK)
		io.WriteString(w, "OK")
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
