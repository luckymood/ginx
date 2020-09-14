package utility

import (
	"ginx/config"

	"go.uber.org/zap"
)

// logger global logger
var logger *zap.Logger

func init() {
	parseLoggerConfig()
}

func parseLoggerConfig() {
	var err error
	if !config.Framework.Zap.Custom {
		logger, err = zap.NewDevelopment()
		if err != nil {
			panic("get new development zap logger error, " + err.Error())
		}
		return
	}
	logger, err = config.Framework.Zap.Config.Build()
	if err != nil {
		panic("build zap logger error, " + err.Error())
	}
	return
}

// Logger get global logger,
// have to get a logger instance,
// or panic
func Logger() *zap.Logger {
	var err error
	if logger == nil {
		logger, err = zap.NewDevelopment()
		if err != nil {
			panic("generate default development zap logger error, " + err.Error())
		}
		return logger
	}
	return logger
}
