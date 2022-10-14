package service

import (
	"io"
	"net/url"
)

const (
	baseUrl  = "https://time.geekbang.org"
	loginUrl = "https://account.geekbang.org"
)

// 获取用户信息
func (s *Service) requestUser() (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", loginUrl+"/account/user", nil, map[string]string{"Origin": loginUrl})
	return handleHTTPResponse(res, err)
}

// 所有购买的课程
func (s *Service) requestBuyAll() (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", baseUrl+"/serv/v1/my/products/all", nil, map[string]string{"Origin": loginUrl})
	return handleHTTPResponse(res, err)
}

// 所有课程
func (s *Service) requestCourses(couseType int) (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", baseUrl+"/serv/v1/column/newAll", map[string]int{"type": couseType}, map[string]string{"Origin": baseUrl})
	return handleHTTPResponse(res, err)
}

// 获取课程信息
func (s *Service) requestCourseDetail(ids []int) (io.ReadCloser, Error) {
	ii := map[string]interface{}{"ids": ids}
	res, err := s.client.Req("POST", baseUrl+"/serv/v1/column/details", ii, map[string]string{"Origin": baseUrl})
	return handleHTTPResponse(res, err)
}

// 课程详细信息
func (s *Service) requestCourseIntro(id int) (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", baseUrl+"/serv/v1/column/intro", map[string]interface{}{"cid": id, "with_groupbuy": true}, map[string]string{"Origin": baseUrl})
	return handleHTTPResponse(res, err)
}

// 文章详情
func (s *Service) requestArticleDetail(id string) (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", baseUrl+"/serv/v1/article", map[string]interface{}{"id": id, "include_neighbors": true, "is_freelyread": true}, map[string]string{"Origin": baseUrl})
	return handleHTTPResponse(res, err)
}

// 文章热门评论
func (s *Service) requestArticleComments(id string) (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", baseUrl+"/serv/v1/comments", map[string]interface{}{"aid": id, "prev": 0}, map[string]string{"Origin": baseUrl})
	return handleHTTPResponse(res, err)
}

// 文章评论留言
func (s *Service) requestCommentDiscussion(id int) (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", baseUrl+"/serv/discussion/v1/root_list",
		map[string]interface{}{"page_type": 1, "target_id": id, "prev": 1, "size": 20, "target_type": 1, "use_likes_order": true},
		map[string]string{"Origin": baseUrl})
	return handleHTTPResponse(res, err)
}

// 课程的文章列表信息
func (s *Service) requestCourseArticles(id int) (io.ReadCloser, Error) {
	data := map[string]interface{}{
		"cid":    id,
		"order":  "earliest",
		"prev":   0,
		"sample": false,
		"size":   500,
	}
	res, err := s.client.Req("POST", baseUrl+"/serv/v1/column/articles", data, map[string]string{"Origin": baseUrl})
	return handleHTTPResponse(res, err)
}

// 获取视频的播放授权信息
func (s *Service) requestVideoPlayAuth(aid int, videoID string) (io.ReadCloser, Error) {
	data := map[string]interface{}{
		"source_type": 1,
		"aid":         aid,
		"video_id":    videoID,
	}
	res, err := s.client.Req("POST", baseUrl+"/serv/v3/source_auth/video_play_auth", data, map[string]string{"Origin": baseUrl})
	return handleHTTPResponse(res, err)
}

// 获取视频的播放信息
func (s *Service) requestVideoPlayInfo(playAuth string) (io.ReadCloser, Error) {
	res, err := s.client.Req("GET", "http://ali.mantv.top/play/info?playAuth="+url.QueryEscape(playAuth), nil, nil)
	return handleHTTPResponse(res, err)
}
