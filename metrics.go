package gooseneck

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// "github.com/prometheus/client_golang/prometheus/promhttp"
)

func MakeCounter(name string, help string) prometheus.Counter {
	return promauto.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: help,
	})
}

var MetricsHandler = promhttp.Handler
