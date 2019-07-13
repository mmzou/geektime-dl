package config

import "errors"

var (
	//ErrNotLogin 未登录帐号错误
	ErrNotLogin = errors.New("请先登录极客时间账户")
	//ErrHasLoginedNotLogin 有登录用户，但是当前并未有有效用户
	ErrHasLoginedNotLogin = errors.New("存在登录的用户，可以进行切换登录用户")
	//ErrConfigFilePathNotSet 未设置配置文件
	ErrConfigFilePathNotSet = errors.New("config file not set")
	//ErrConfigFileNotExist 未设置Config, 未初始化
	ErrConfigFileNotExist = errors.New("config file not exist")
	//ErrConfigFileNoPermission Config文件无权限访问
	ErrConfigFileNoPermission = errors.New("config file permission denied")
	//ErrConfigContentsParseError 解析Config数据错误
	ErrConfigContentsParseError = errors.New("config contents parse error")
)
