package service

import "errors"

var (
	//ErrNotLogin not login
	ErrNotLogin = errors.New("当前未登录获取登录已失效，请先登录")
	//ErrLoginOffline not login
	ErrLoginOffline = errors.New("该账号已在其他同类设备登录，如非本人操作，则密码可能已经泄露，建议立即更换密码")
)

// Error 错误信息接口
type Error interface {
	error
	IsUnlogin() bool
}

//ErrorInfo error info
type ErrorInfo struct {
	Err error
}

//IsUnlogin 是否未登录
func (e *ErrorInfo) IsUnlogin() bool {
	return e.Err == ErrNotLogin
}

func (e *ErrorInfo) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return ""
}

func (e *ErrorInfo) String() string {
	return e.Error()
}
