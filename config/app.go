package config

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// App global config
var App struct {
	Secret string `yaml:"secret"`
}

// ParseAppConfig parse default app-config in conf/app.yaml
func ParseAppConfig() {
	fp, err := os.Open(path.Join(Root(), "conf/app.yaml"))
	if err != nil {
		panic("open app config error, " + err.Error())
	}
	defer fp.Close()
	bs, err := ioutil.ReadAll(fp)
	if err != nil {
		panic("read app config error, " + err.Error())
	}
	if err := yaml.Unmarshal(bs, &App); err != nil {
		panic("parse app config error, " + err.Error())
	}
}
