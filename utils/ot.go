package utils

import (
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerlog "github.com/uber/jaeger-client-go/log"
	"github.com/uber/jaeger-lib/metrics"
)

// SetGlobalTracer set global tracer
func SetGlobalTracer() {
	cfg := jaegercfg.Configuration{
		Reporter: &jaegercfg.ReporterConfig{
			LocalAgentHostPort: "localhost:6831",
		},
	}
	jLogger := jaegerlog.StdLogger
	jMetricsFactory := metrics.NullFactory

	_, err := cfg.InitGlobalTracer(
		"GINX",
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		Logger().Panic("set global tracer failed.")
	}
}
