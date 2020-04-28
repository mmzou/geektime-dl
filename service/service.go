package service

import (
	"net/http"
	"net/url"

	"github.com/mmzou/geektime-dl/requester"
)

var (
	geekBangCommURL = &url.URL{
		Scheme: "https",
		Host:   "geekbang.org",
	}
)

//Service geek time service
type Service struct {
	client *requester.HTTPClient
}

//NewService new service
func NewService(gcid, gcess, serviceID string) *Service {
	client := requester.NewHTTPClient()
	client.ResetCookieJar()
	cookies := []*http.Cookie{}
	cookies = append(cookies, &http.Cookie{
		Name:   "GCID",
		Value:  gcid,
		Domain: "." + geekBangCommURL.Host,
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "GCESS",
		Value:  gcess,
		Domain: "." + geekBangCommURL.Host,
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "SERVERID",
		Value:  serviceID,
		Domain: "." + geekBangCommURL.Host,
	})
	client.Jar.SetCookies(geekBangCommURL, cookies)

	return &Service{client: client}
}

//Cookies get cookies string
func (s *Service) Cookies() map[string]string {
	cookies := s.client.Jar.Cookies(geekBangCommURL)

	cstr := map[string]string{}

	for _, cookie := range cookies {
		cstr[cookie.Name] = cookie.Value
	}

	return cstr
}
