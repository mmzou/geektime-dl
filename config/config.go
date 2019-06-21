package config

import (
	"os"
	"path/filepath"
	"sync"

	jsoniter "github.com/json-iterator/go"
)

const (
	// EnvConfigDir 配置路径环境变量
	EnvConfigDir = "GEEKTIME_GO_CONFIG_DIR"
	// ConfigName 配置文件名
	ConfigName = "config.json"
)

var (
	configFilePath = filepath.Join(GetConfigDir(), ConfigName)

	//Config 配置信息 全局调用
	Config = NewConfig(configFilePath)
)

//ConfigsData 配置数据
type ConfigsData struct {
	Geektimes Geektimes

	configFilePath string
	configFile     *os.File
	fileMu         sync.Mutex
}

//Init 初始化配置
func (c *ConfigsData) Init() error {
	if c.configFilePath == "" {
		return ErrConfigFilePathNotSet
	}

	c.initDefaultConfig()
	err := c.loadConfigFromFile()
	if err != nil {
		return err
	}

	return nil
}

//Save 保存配置
func (c *ConfigsData) Save() error {
	err := c.lazyOpenConfigFile()
	if err != nil {
		return err
	}

	c.fileMu.Lock()
	defer c.fileMu.Unlock()

	//todo 保存配置的数据
	// data, err := jsoniter.MarshalIndent((*pcsConfigJSONExport)(unsafe.Pointer(c)), "", "")
	data, err := jsoniter.MarshalIndent(c, "", "")

	if err != nil {
		panic(err)
	}

	// 减掉多余的部分
	err = c.configFile.Truncate(int64(len(data)))
	if err != nil {
		return err
	}

	_, err = c.configFile.Seek(0, os.SEEK_SET)
	if err != nil {
		return err
	}

	_, err = c.configFile.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (c *ConfigsData) initDefaultConfig() {
	//todo 默认配置
}

func (c *ConfigsData) loadConfigFromFile() error {
	//todo 从配置文件中加载配置
	err := c.lazyOpenConfigFile()
	if err != nil {
		return err
	}

	return nil
}

func (c *ConfigsData) lazyOpenConfigFile() (err error) {
	if c.configFile != nil {
		return nil
	}
	c.fileMu.Lock()
	os.MkdirAll(filepath.Dir(c.configFilePath), 0700)
	c.configFile, err = os.OpenFile(c.configFilePath, os.O_CREATE|os.O_RDWR, 0600)
	c.fileMu.Unlock()

	if err != nil {
		if os.IsPermission(err) {
			return ErrConfigFileNoPermission
		}
		if os.IsExist(err) {
			return ErrConfigFileNotExist
		}
		return err
	}

	return nil
}

//NewConfig new config
func NewConfig(configFilePath string) *ConfigsData {
	c := &ConfigsData{
		configFilePath: configFilePath,
	}

	return c
}

//Geektimes 极客时间用户
type Geektimes []*Geektime

//GetConfigDir 配置文件夹
func GetConfigDir() string {
	configDir, ok := os.LookupEnv(EnvConfigDir)
	if ok {
		if filepath.IsAbs(configDir) {
			return configDir
		}
	}

	home, ok := os.LookupEnv("HOME")
	if ok {
		return filepath.Join(home, ".config", "geekbang")
	}

	return filepath.Join("/tmp", "geekbang")
}
