package config

import (
	"io/ioutil"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

// Framework framework config
var Framework struct {
	Mode string `yaml:"mode"`
	Port string `yaml:"port"`
}

// ParseFrameworkConfig parse default framework-config in conf/framework.yaml
func ParseFrameworkConfig() {
	fp, err := os.Open(path.Join(Root(), "conf/framework.yaml"))
	if err != nil {
		panic("open framework config error, " + err.Error())
	}
	defer fp.Close()
	bs, err := ioutil.ReadAll(fp)
	if err != nil {
		panic("read framework config error, " + err.Error())
	}
	if err := yaml.Unmarshal(bs, &Framework); err != nil {
		panic("parse framework config error, " + err.Error())
	}
}
