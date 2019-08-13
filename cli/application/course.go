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

//CourseWithArticles course and articles info
func CourseWithArticles(id int) (*service.Course, []*service.Article, error) {
	course, err := getService().ShowCourse(id)
	if err != nil {
		return nil, nil, err
	}

	articles, err := getService().Articles(id)
	if err != nil {
		return course, nil, err
	}

	return course, articles, nil
}
