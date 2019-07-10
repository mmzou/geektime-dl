package service

import "io"

//RequestUser get user info
func (s *Service) RequestUser() (io.ReadCloser, Error) {
	res, err := s.client.Req("POST", "https://account.geekbang.org/account/user", nil, map[string]string{"Origin": "https://account.geekbang.org"})

	if err != nil {
		defer res.Body.Close()
		return nil, &ErrorInfo{Err: err}
	}

	return res.Body, nil
}
