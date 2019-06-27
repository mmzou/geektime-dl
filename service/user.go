package service

import "net/http"

//User get user info
func (s *Service) User() (*http.Response, error) {
	return s.client.Req("POST", "https://account.geekbang.org/account/user", nil, map[string]string{"Origin": "https://account.geekbang.org"})
}
