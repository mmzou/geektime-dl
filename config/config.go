package config

import (
	"os"
	"path/filepath"
	"sync"
	"unsafe"

	jsoniter "github.com/json-iterator/go"
	"github.com/mmzou/geektime-dl/service"
)

const (
	// EnvConfigDir 配置路径环境变量
	EnvConfigDir = "GEEKTIME_GO_CONFIG_DIR"
	// ConfigName 配置文件名
	ConfigName = "config.json"
)

var (
	configFilePath = filepath.Join(GetConfigDir(), ConfigName)

	//Instance 配置信息 全局调用
	Instance = NewConfig(configFilePath)
)

//ConfigsData 配置数据
type ConfigsData struct {
	AcitveUID    int
	Geektimes    Geektimes
	DownloadPath string

	activeUser     *Geektime
	configFilePath string
	configFile     *os.File
	fileMu         sync.Mutex
	service        *service.Service
}

//Init 初始化配置
func (c *ConfigsData) Init() error {
	if c.configFilePath == "" {
		return ErrConfigFilePathNotSet
	}

	//初始化默认配置
	c.initDefaultConfig()
	//从配置文件中加载配置
	err := c.loadConfigFromFile()
	if err != nil {
		return err
	}

	//初始化登陆用户信息
	err = c.initActiveUser()
	if err != nil {
		return nil
	}

	if c.activeUser != nil {
		c.service = c.activeUser.Service()
	}

	return nil
}

func (c *ConfigsData) initActiveUser() error {
	//如果已经初始化过，则跳过
	if c.AcitveUID > 0 && c.activeUser != nil && c.activeUser.ID == c.AcitveUID {
		return nil
	}

	if c.AcitveUID == 0 && c.activeUser != nil {
		c.AcitveUID = c.activeUser.ID
		return nil
	}

	if c.AcitveUID > 0 {
		for _, geek := range c.Geektimes {
			if geek.ID == c.AcitveUID {
				c.activeUser = geek
				return nil
			}
		}
	}

	if len(c.Geektimes) > 0 {
		return ErrHasLoginedNotLogin
	}

	return ErrNotLogin
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
	data, err := jsoniter.MarshalIndent((*configJSONExport)(unsafe.Pointer(c)), "", " ")

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
	err := c.lazyOpenConfigFile()
	if err != nil {
		return err
	}

	info, err := c.configFile.Stat()
	if err != nil {
		return err
	}

	if info.Size() == 0 {
		return c.Save()
	}

	c.fileMu.Lock()
	defer c.fileMu.Unlock()

	_, err = c.configFile.Seek(0, os.SEEK_SET)
	if err != nil {
		return nil
	}

	//从配置文件中加载配置
	decoder := jsoniter.NewDecoder(c.configFile)
	decoder.Decode((*configJSONExport)(unsafe.Pointer(c)))

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

//ActiveUser active user
func (c *ConfigsData) ActiveUser() *Geektime {
	return c.activeUser
}

//ActiveUserService user service
func (c *ConfigsData) ActiveUserService() *service.Service {
	if c.service == nil {
		c.service = c.activeUser.Service()
	}
	return c.service
}
