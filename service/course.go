package service

import "github.com/mmzou/geektime-dl/utils"

//Columns 获取专栏
func (s *Service) Columns() ([]*Course, error) {
	return s.getCourses(1)
}

//Videos 获取专栏
func (s *Service) Videos() ([]*Course, error) {
	return s.getCourses(3)
}

//获取课程信息
func (s *Service) getCourses(courseType int) ([]*Course, error) {
	ids, err := s.courses(courseType)

	if err != nil {
		return nil, err
	}

	return s.getCourseDetail(ids)
}

func (s *Service) courses(courseType int) ([]int, error) {
	body, err := s.requestCourses(courseType)

	if err != nil {
		return nil, err
	}

	defer body.Close()

	courses := new(CourseList)
	if err := handleJSONParse(body, &courses); err != nil {
		return nil, err
	}

	var ids []int
	for _, item := range courses.List {
		ids = append(ids, item.ID)
	}

	return ids, nil
}

func (s *Service) getCourseDetail(ids []int) ([]*Course, error) {
	body, err := s.requestCourseDetail(ids)
	if err != nil {
		return nil, err
	}

	defer body.Close()

	var courses []*Course
	if err := handleJSONParse(body, &courses); err != nil {
		return nil, err
	}

	return courses, nil
}

//ShowCourse 获取课程信息
func (s *Service) ShowCourse(id int) (*Course, error) {
	body, err := s.requestCourseIntro(id)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	course := new(Course)
	if err := handleJSONParse(body, &course); err != nil {
		return nil, err
	}

	return course, nil
}

//Articles get course articles
func (s *Service) Articles(id int) ([]*Article, error) {
	body, err := s.requestCourseArticles(id)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	articleResult := &articleResult{}
	if err := handleJSONParse(body, articleResult); err != nil {
		return nil, err
	}

	return articleResult.Articles, nil
}

//VideoPlayAuth 获取视频的播放授权信息
func (s *Service) VideoPlayAuth(aid int, videoID string) (*VideoPlayAuth, error) {
	body, err := s.requestVideoPlayAuth(aid, videoID)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	videoPlayAuth := &VideoPlayAuth{}
	if err := handleJSONParse(body, videoPlayAuth); err != nil {
		return nil, err
	}

	return videoPlayAuth, nil
}

//VideoPlayInfo 获取视频播放信息
func (s *Service) VideoPlayInfo(playAuth string) (*VideoPlayInfo, error) {
	body, err := s.requestVideoPlayInfo(playAuth)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	videoPlayInfo := &VideoPlayInfo{}
	if err := utils.UnmarshalReader(body, &videoPlayInfo); err != nil {
		return nil, err
	}

	return videoPlayInfo, nil
}
