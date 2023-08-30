package gooseneck

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func MakeCounter(namespace string, subsystem string, name string, help string, constLabels prometheus.Labels) prometheus.Counter {
	return promauto.NewCounter(prometheus.CounterOpts{
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	})
}

func MakeCounterVec(namespace string, subsystem string, name string, help string, constLabels prometheus.Labels, labelNames []string) *prometheus.CounterVec {
	return promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}, labelNames)
}

var MetricsHandler = promhttp.Handler
