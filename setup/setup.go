package setup

import (
	"ginx/config"
	"ginx/utils"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	logger := utils.Logger()

	gin.SetMode(config.Framework.Mode)

	setEngine()
	logger.Info("gin engine is ready.")

	utils.SetGlobalTracer()
	logger.Info("global tracer is ready.")
}

// setEngine run new engine
func setEngine() {
	engine = gin.New()
	initRouters()
}

// Engine get engine
func Engine() *gin.Engine {
	if engine == nil {
		panic("please set engine first")
	}
	return engine
}

// Run run engine
func Run() {
	if engine == nil {
		panic("please set engine first")
	}
	if err := engine.Run(config.Framework.Port); err != nil {
		panic("run engine error, " + err.Error())
	}
}
