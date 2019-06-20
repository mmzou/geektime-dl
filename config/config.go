package config

import (
	"os"
	"path/filepath"
)

const (
	// EnvConfigDir 配置路径环境变量
	EnvConfigDir = "GEEKTIME_GO_CONFIG_DIR"
	// ConfigName 配置文件名
	ConfigName = "pcs_config.json"
)

var (
	Config = NewConfig(ConfigName)
)

type ConfigData struct {
	Geektimes Geektimes

	configFilePath string
}

func NewConfig(configFilePath string) *ConfigData {
	c := &ConfigData{
		configFilePath: configFilePath,
	}

	return c
}

//Geektimes 极客时间用户
type Geektimes []*Geektime

func GetConfigDIr() string {

	home, err := os.LookupEnv("HOME")
	if !err {
		home = "/tmp"
	}

	return filepath.Join(home, ".config", "geekbang")
}
