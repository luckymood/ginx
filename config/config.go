package config

import (
	"os"
)

var root string

func init() {
	initRoot()
	ParseFrameworkConfig()
	ParseAppConfig()
}

// initRoot init root path
func initRoot() {
	dir, err := os.Getwd()
	if err != nil {
		panic("get pwd error, " + err.Error())
	}
	root = dir
}

// Root return root path
func Root() string {
	return root
}
