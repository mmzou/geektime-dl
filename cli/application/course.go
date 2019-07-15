package application

import (
	"github.com/mmzou/geektime-dl/service"
)

//Columns 专栏列表
func Columns() ([]*service.Course, error) {
	return getService().Columns()
}

//Videos 视频课程列表
func Videos() ([]*service.Course, error) {
	return getService().Videos()
}
