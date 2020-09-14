package setup

import (
	"ginx/config"
	"ginx/utility"

	"github.com/gin-gonic/gin"
)

var engine *gin.Engine

func init() {
	gin.SetMode(config.Framework.Mode)
	setEngine()
}

// setEngine run new engine
func setEngine() {
	engine = gin.New()
	initRouters()
	utility.Logger().Info("successfully set engine...")
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
