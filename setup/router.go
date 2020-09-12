package setup

import (
	"ginx/controller"
)

// initRouters init routers
func initRouters() {

	r := Engine()

	r.GET("/health", controller.Health)
}
