package requester

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

//HTTPClient client
type HTTPClient struct {
	http.Client
	UserAgent string
}

//NewHTTPClient new client
func NewHTTPClient() *HTTPClient {
	c := &HTTPClient{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}

	c.ResetCookieJar()

	return c
}

// SetUserAgent 设置 UserAgent 浏览器标识
func (h *HTTPClient) SetUserAgent(ua string) {
	h.UserAgent = ua
}

//SetCookiejar 设置 Cookie
func (h *HTTPClient) SetCookiejar(jar http.CookieJar) {
	h.Client.Jar = jar
}

//ResetCookieJar 重置cookie jar
func (h *HTTPClient) ResetCookieJar() {
	h.Jar, _ = cookiejar.New(nil)
}

//SetTimeout 设置超时时间
func (h *HTTPClient) SetTimeout(t time.Duration) {
	h.Client.Timeout = t
}
