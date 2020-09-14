package setup

import (
	"ginx/controller"
	"ginx/middleware"
)

// initRouters init routers
func initRouters() {

	/*
		common api
	*/

	r := Engine()

	r.GET("/health", controller.Health)
	r.GET("/issue", controller.Issue)

	/*
		jwt-token required
	*/

	api := r.Group("/api", middleware.JwtAuthorize())

	api.GET("/authorize", controller.CheckIssue)
}
