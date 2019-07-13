package service

import "io"

//requestUser get user info
func (s *Service) requestUser() (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", "https://account.geekbang.org/account/user", nil, map[string]string{"Origin": "https://account.geekbang.org"})

	if err != nil {
		defer res.Body.Close()
		return nil, &ErrorInfo{Err: err}
	}

	return res.Body, nil
}

//requestProductAll 所有订阅课程
func (s *Service) requestProductAll() (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", "https://time.geekbang.org/serv/v1/my/products/all", nil, map[string]string{"Origin": "https://account.geekbang.org"})

	if err != nil {
		defer res.Body.Close()
		return nil, &ErrorInfo{Err: err}
	}

	return res.Body, nil
}
