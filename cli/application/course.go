package application

import (
	"strings"

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

	course.ColumnTitle = strings.TrimSpace(course.ColumnTitle)

	articles, err := getService().Articles(id)
	if err != nil {
		return course, nil, err
	}

	return course, articles, nil
}

//GetVideoPlayInfo 获取视频播放信息
func GetVideoPlayInfo(aid int, videoID string) (*service.VideoPlayInfo, error) {
	videoPlayAuth, err := VideoPlayAuth(aid, videoID)
	if err != nil {
		return nil, err
	}

	videoPlayInfo, err := VideoPlayInfo(videoPlayAuth.PlayAuth)
	if err != nil {
		return nil, err
	}
	return videoPlayInfo, nil
}

//VideoPlayAuth 获取视频的播放授权信息
func VideoPlayAuth(aid int, videoID string) (*service.VideoPlayAuth, error) {
	videoPlayAuth, err := getService().VideoPlayAuth(aid, videoID)

	if err != nil {
		return nil, err
	}

	return videoPlayAuth, nil
}

//VideoPlayInfo 获取视频播放信息
func VideoPlayInfo(playAuth string) (*service.VideoPlayInfo, error) {
	videoPlayInfo, err := getService().VideoPlayInfo(playAuth)

	if err != nil {
		return nil, err
	}

	return videoPlayInfo, nil
}
