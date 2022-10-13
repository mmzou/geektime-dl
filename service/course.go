package service

import (
	"strconv"

	"github.com/mmzou/geektime-dl/utils"
)

// Columns 获取专栏
func (s *Service) Columns() ([]*Course, error) {
	return s.getCourses(1)
}

// Videos 获取专栏
func (s *Service) Videos() ([]*Course, error) {
	return s.getCourses(3)
}

// 获取课程信息
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

// ShowCourse 获取课程信息
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

// Article get course article by id
func (s *Service) Article(id int) (*Article, error) {
	body, err := s.requestArticleDetail(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer body.Close()

	articleDetail := &Article{}
	if err := handleJSONParse(body, articleDetail); err != nil {
		return nil, err
	}

	return articleDetail, nil
}

// ArticleCommentsWithDiscussion get course article ArticleComments by id
func (s *Service) ArticleCommentsWithDiscussion(id int) (*CommentList, error) {
	body, err := s.requestArticleComments(strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	defer body.Close()

	list := &CommentList{}
	if err := handleJSONParse(body, list); err != nil {
		return nil, err
	}

	for i, comment := range list.List {
		if comment.DiscussionCount > 0 {
			body, err = s.requestCommentDiscussion(comment.Id)
			if err != nil {
				return nil, err
			}
			defer body.Close()

			discussionList := &CommentDiscussions{}
			if err := handleJSONParse(body, discussionList); err != nil {
				return nil, err
			}
			list.List[i].CommentDiscussions = discussionList
		}
	}

	return list, nil
}

// CommentsDiscussion get comment discussion  by id
func (s *Service) CommentsDiscussion(id int) (*CommentDiscussions, error) {
	body, err := s.requestCommentDiscussion(id)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	list := &CommentDiscussions{}
	if err := handleJSONParse(body, list); err != nil {
		return nil, err
	}

	return list, nil
}

// Articles get course articles
func (s *Service) Articles(id int) ([]*Article, error) {
	body, err := s.requestCourseArticles(id)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	articleRes := &articleResult{}
	if err = handleJSONParse(body, articleRes); err != nil {
		return nil, err
	}

	return articleRes.Articles, nil
}

// VideoPlayAuth 获取视频的播放授权信息
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

// VideoPlayInfo 获取视频播放信息
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
