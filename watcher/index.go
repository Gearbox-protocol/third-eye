package watcher

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Gearbox-protocol/sdk-go/utils"
	"github.com/Gearbox-protocol/third-eye/config"
	"github.com/Gearbox-protocol/third-eye/ds"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var Version = "dev"

type Metrics struct {
	Version     string `json:"version"`
	LatestBlock int64  `json:"latestBlock"`
	Uptime      string `json:"uptime"`
}

func newMetEngine(eng ds.EngineI, _cfg *config.Config) {
	//
	mux := http.NewServeMux()
	startedAt := time.Now().UTC()
	startUnix := float64(startedAt.Unix())

	reg := prometheus.NewRegistry()
	reg.MustRegister(
		// pseudo-metric that provides metadata about the running binary
		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "third_eye_build_info",
			Help: "Build info",
			ConstLabels: prometheus.Labels{
				"version": Version,
			},
		}, func() float64 { return 1.0 }),

		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "eth_block_number",
			Help: "Latest processed block",
		}, func() float64 { return float64(eng.LastSyncedBlock()) }),

		prometheus.NewGaugeFunc(prometheus.GaugeOpts{
			Name: "start_time",
			Help: "Start time, in unixtime (seconds)",
		}, func() float64 { return startUnix }),
	)
	mux.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	mux.HandleFunc("/health", func(hw http.ResponseWriter, hr *http.Request) {
		resp := Metrics{
			Version:     Version,
			LatestBlock: eng.LastSyncedBlock(),
			Uptime:      time.Since(startedAt).Round(time.Second).String(),
		}
		d, _ := json.Marshal(resp)
		fmt.Fprint(hw, string(d))
	})

	utils.ServerFromMux(mux, _cfg.Port)
}
