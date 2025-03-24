package engine

import "github.com/prometheus/client_golang/prometheus"

func getRegistry() *prometheus.Registry {
	return prometheus.NewRegistry()
}
