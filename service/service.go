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
		Domain: ".geekbang.org",
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "GCESS",
		Value:  gcess,
		Domain: ".geekbang.org",
	})
	cookies = append(cookies, &http.Cookie{
		Name:   "SERVERID",
		Value:  serviceID,
		Domain: ".geekbang.org",
	})
	client.Jar.SetCookies(geekBangCommURL, cookies)

	return &Service{client: client}
}
