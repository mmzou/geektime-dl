package login

import (
	"fmt"
	"net/http/cookiejar"
	"net/url"
	"strings"

	jsoniter "github.com/json-iterator/go"
	"github.com/mmzou/geektime-dl/requester"
)

//Client login client
type Client struct {
	*requester.HTTPClient
}

// Result 从百度服务器解析的数据结构
type Result struct {
	Code int `json:"code"`
	Data struct {
		UID          int    `json:"uid"`
		Name         string `json:"nickname"`
		Avatar       string `json:"avatar"`
		GCID         string `json:"gcid"`
		GCESS        string `json:"gcess"`
		ServerID     string `json:"serverId"`
		Ticket       string `json:"ticket"`
		CookieString string `json:"cookieString"`
	} `json:"data"`
	Error struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"error"`
	Extra struct {
		Cost      float64 `json:"cost"`
		RequestID string  `json:"request-iud"`
	} `json:"extra"`
}

//NewLoginClient new login client
func NewLoginClient() *Client {
	c := &Client{
		HTTPClient: requester.NewHTTPClient(),
	}

	c.InitLoginPage()

	return c
}

//InitLoginPage init
func (c *Client) InitLoginPage() {
	res, _ := c.Get("https://account.geekbang.org/signin?redirect=https%3A%2F%2Ftime.geekbang.org%2F")
	res.Body.Close()
}

//Login by phone and dpassword
func (c *Client) Login(phone, password string) *Result {
	result := &Result{}
	post := map[string]string{
		"country":   "86",
		"cellphone": phone,
		"password":  password,
		"captcha":   "",
		"remeber":   "1",
		"platform":  "3",
		"appid":     "1",
	}

	header := map[string]string{
		"Referer": "https://account.geekbang.org/signin?redirect=https%3A%2F%2Ftime.geekbang.org%2F",
		"Accept":  "application/json",
		// "Content-Type": "application/json",
		"Content-Type": "application/x-www-form-urlencoded",
		"Connection":   "keep-alive",
	}
	body, err := c.Fetch("POST", "https://account.geekbang.org/account/ticket/login", post, header)
	if err != nil {
		result.Code = -1
		result.Error.Code = -1
		result.Error.Msg = "网络请求失败, " + err.Error()

		return result
	}

	if err = jsoniter.Unmarshal(body, &result); err != nil {
		result.Code = -1
		result.Error.Code = -1
		result.Error.Msg = "发送登录请求错误: " + err.Error()

		return result
	}

	if result.Code == 0 {
		result.parseCookies("https://account.geekbang.org", c.Jar.(*cookiejar.Jar))
	}

	url, _ := url.Parse("https://account.geekbang.org")
	fmt.Println(body, post, c.Jar.Cookies(url))

	return result
}

//parseCookies 解析cookie
func (r *Result) parseCookies(targetURL string, jar *cookiejar.Jar) {
	url, _ := url.Parse(targetURL)
	cookies := jar.Cookies(url)

	cookieArr := []string{}
	for _, cookie := range cookies {
		switch cookie.Name {
		case "GCID":
			r.Data.GCID = cookie.Value
		case "GCESS":
			r.Data.GCESS = cookie.Value
		case "SERVERID":
			r.Data.ServerID = cookie.Value
		}
		cookieArr = append(cookieArr, cookie.String())
	}
	r.Data.CookieString = strings.Join(cookieArr, ";")
}
