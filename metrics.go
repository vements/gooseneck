package gooseneck

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	MetricsHandler = promhttp.Handler
)

type (
	MetricLabels     = prometheus.Labels
	MetricLabelNames = []string
)

func MakeCounter(namespace string, subsystem string, name string, help string, constLabels MetricLabels) prometheus.Counter {
	return promauto.NewCounter(prometheus.CounterOpts{
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	})
}

func MakeCounterFunc(namespace string, subsystem string, name string, help string, constLabels MetricLabels, fn func() float64) prometheus.CounterFunc {
	return promauto.NewCounterFunc(prometheus.CounterOpts{
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}, fn)
}

func MakeCounterVec(namespace string, subsystem string, name string, help string, constLabels MetricLabels, labelNames MetricLabelNames) *prometheus.CounterVec {
	return promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}, labelNames)
}

func MakeGauge(namespace string, subsystem string, name string, help string, constLabels MetricLabels) prometheus.Gauge {
	return promauto.NewGauge(prometheus.GaugeOpts{
		Namespace:   name,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	})
}

func MakeGaugeFunc(namespace string, subsystem string, name string, help string, constLabels MetricLabels, fn func() float64) prometheus.GaugeFunc {
	return promauto.NewGaugeFunc(prometheus.GaugeOpts{
		Namespace:   namespace,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}, fn)
}

func MakeGaugeVec(namespace string, subsystem string, name string, help string, constLabels MetricLabels, labelNames MetricLabelNames) *prometheus.GaugeVec {
	return promauto.NewGaugeVec(prometheus.GaugeOpts{
		Namespace:   name,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}, labelNames)
}

func MakeHistogram(namespace string, subsystem string, name string, help string, constLabels MetricLabels) prometheus.Histogram {
	return promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace:   name,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	})
}

func MakeHistogramVec(namespace string, subsystem string, name string, help string, constLabels MetricLabels, labelNames MetricLabelNames) *prometheus.HistogramVec {
	return promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   name,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}, labelNames)
}

func MakeSummary(namespace string, subsystem string, name string, help string, constLabels MetricLabels) prometheus.Summary {
	return promauto.NewSummary(prometheus.SummaryOpts{
		Namespace:   name,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	})
}

func MakeSummaryVec(namespace string, subsystem string, name string, help string, constLabels MetricLabels, labelNames MetricLabelNames) *prometheus.SummaryVec {
	return promauto.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:   name,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}, labelNames)
}

func MakeUntypedFunc(namespace string, subsystem string, name string, help string, constLabels MetricLabels, fn func() float64) prometheus.UntypedFunc {
	return promauto.NewUntypedFunc(prometheus.UntypedOpts{
		Namespace:   name,
		Subsystem:   subsystem,
		Name:        name,
		Help:        help,
		ConstLabels: constLabels,
	}, fn)
}
