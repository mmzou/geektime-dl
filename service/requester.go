package service

import (
	"io"
)

//获取用户信息
func (s *Service) requestUser() (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", "https://account.geekbang.org/account/user", nil, map[string]string{"Origin": "https://account.geekbang.org"})
	return handleHTTPResponse(res, err)
}

//所有购买的课程
func (s *Service) requestBuyAll() (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", "https://time.geekbang.org/serv/v1/my/products/all", nil, map[string]string{"Origin": "https://account.geekbang.org"})
	return handleHTTPResponse(res, err)
}

//所有课程
func (s *Service) requestCourses(couseType int) (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", "https://time.geekbang.org/serv/v1/column/newAll", map[string]int{"type": couseType}, map[string]string{"Origin": "https://time.geekbang.org"})
	return handleHTTPResponse(res, err)
}

//获取课程信息
func (s *Service) requestCourseDetail(ids []int) (io.ReadCloser, Error) {
	ii := map[string]interface{}{"ids": ids}
	res, err := s.client.Req("POST", "https://time.geekbang.org/serv/v1/column/details", ii, map[string]string{"Origin": "https://time.geekbang.org"})
	return handleHTTPResponse(res, err)
}
