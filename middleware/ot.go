package middleware

import (
	"ginx/utils"
	"path"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

// MakeOpenTracingMiddleware make open tracing middleware using jaeger
func MakeOpenTracingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		logger := utils.Logger()

		tracer := opentracing.GlobalTracer()
		if tracer == nil {
			logger.Warn("tracer is nil")
			return
		}

		var span opentracing.Span
		ctx, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err == nil {
			span = tracer.StartSpan(
				path.Join("HTTP", c.Request.Method, c.Request.URL.Path),
				//fmt.Sprintf("HTTP/%s/%s", c.Request.Method, c.Request.URL.Path),
				opentracing.ChildOf(ctx),
			)
		} else {
			span = tracer.StartSpan(
				path.Join("HTTP", c.Request.Method, c.Request.URL.Path),
				//fmt.Sprintf("HTTP/%s/%s", c.Request.Method, c.Request.URL.Path),
			)
		}
		defer span.Finish()

		// save span_context into context
		c.Set("SpanContext", span.Context())

		c.Next()

		//ext.HTTPStatusCode.Set(span, uint16(c.Writer.Status()))
		//ext.Component.Set(span, "http")
	}
}
