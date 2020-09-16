package middleware

import (
	"ginx/metric"
	"ginx/utility"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

// Log access log
func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		// 后面的中间件可能抛出错误
		defer func() {
			cost := time.Since(start)

			utility.Logger().Info("accesslog", // ElasticSearch-friendly
				zap.String("method", c.Request.Method),
				zap.String("path", c.Request.URL.Path),
				zap.String("query", c.Request.URL.RawQuery),
				zap.String("ip", c.ClientIP()),
				zap.String("fwd", c.Request.Header.Get("X-Forwarded-Server")),
				zap.Int("reqsize", computeApproximateRequestSize(c.Request)),
				zap.Int("repsize", c.Writer.Size()),
				zap.Int("code", c.Writer.Status()),
				zap.String("ua", c.Request.UserAgent()),
				zap.Time("start", start),
				zap.Duration("duration", cost),
			)

			metric.RequestDurationSeconds.With(prometheus.Labels{
				"code":     strconv.Itoa(c.Writer.Status()),
				"handler":  c.Request.URL.Path,
				"instance": c.Request.Host,
			}).Observe(cost.Seconds())
		}()

		c.Next()
	}
}

// From https://github.com/DanielHeckrath/gin-prometheus/blob/master/gin_prometheus.go
func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.String())
	}

	s += len(r.Method)
	s += len(r.Proto)
	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}
	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}
	return s
}
