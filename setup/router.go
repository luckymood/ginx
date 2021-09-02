package setup

import (
	"ginx/controller"
	"ginx/metric"
	"ginx/middleware"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
)

// initRouters init routers
func initRouters() {

	// register prometheus
	prometheus.MustRegister(
		metric.RequestDurationSeconds,
	)

	// common API
	r := Engine()
	r.Use(middleware.MakeOpenTracingMiddleware(), middleware.MakeLogMiddleware())
	// healthy check
	r.GET("/health", controller.Health)
	// expose metrics
	r.GET("/metrics", func(handler http.Handler) gin.HandlerFunc {
		return func(c *gin.Context) {
			handler.ServeHTTP(c.Writer, c.Request)
		}
	}(promhttp.Handler()))
	// issue jwt-token
	r.GET("/api/issue", controller.Issue)

	// jwt-token required
	api := r.Group("/api", middleware.JwtAuthorize())
	api.GET("/authorize", controller.CheckIssue)

}
