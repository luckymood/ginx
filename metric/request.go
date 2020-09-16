package metric

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	// RequestDurationSeconds metric about request duration seconds
	RequestDurationSeconds = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Help collecting request duration",
			Buckets: []float64{0.1, 0.3, 0.5, 0.7, 0.9},
		},
		[]string{"code", "handler", "instance"},
	)
)
