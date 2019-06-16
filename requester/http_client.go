package requester

import (
	"net/http"
	"net/http/cookiejar"
	"time"
)

//HTTPCLient client
type HTTPCLient struct {
	http.Client
	UserAgent string
}

//NewHTTPClient new client
func NewHTTPClient() *HTTPCLient {
	c := &HTTPCLient{
		Client: http.Client{
			Timeout: 10 * time.Second,
		},
	}

	c.ResetCookieJar()

	return c
}

// SetUserAgent 设置 UserAgent 浏览器标识
func (h *HTTPCLient) SetUserAgent(ua string) {
	h.UserAgent = ua
}

//SetCookiejar 设置 Cookie
func (h *HTTPCLient) SetCookiejar(jar http.CookieJar) {
	h.Client.Jar = jar
}

//ResetCookieJar 重置cookie jar
func (h *HTTPCLient) ResetCookieJar() {
	h.Jar, _ = cookiejar.New(nil)
}

//SetTimeout 设置超时时间
func (h *HTTPCLient) SetTimeout(t time.Duration) {
	h.Client.Timeout = t
}
