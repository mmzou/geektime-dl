package requester

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
)

//HTTPGet 简单实现 http 访问 GET 请求
func HTTPGet(urlStr string) ([]byte, error) {
	res, err := DefaultClient.Get(urlStr)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}

// Req 参见 *HTTPClient.Req, 使用默认 http 客户端
func Req(method string, urlStr string, post interface{}, header map[string]string) (*http.Response, error) {
	return DefaultClient.Req(method, urlStr, post, header)
}

// Fetch 参见 *HTTPClient.Fetch, 使用默认 http 客户端
func Fetch(method string, urlStr string, post interface{}, header map[string]string) ([]byte, error) {
	return DefaultClient.Fetch(method, urlStr, post, header)
}

// Headers return the HTTP Headers of the url
func Headers(url string) (http.Header, error) {
	res, err := Req(http.MethodGet, url, nil, nil)
	if err != nil {
		return nil, err
	}
	return res.Header, nil
}

// Size get size of the url
func Size(url string) (int, error) {
	h, err := Headers(url)
	if err != nil {
		return 0, err
	}
	s := h.Get("Content-Length")
	if s == "" {
		return 0, errors.New("Content-Length is not present")
	}
	size, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return size, nil
}

// Req 实现 http／https 访问，
// 根据给定的 method (GET, POST, HEAD, PUT 等等),
// urlStr (网址),
// post (post 数据),
// header (header 请求头数据), 进行网站访问。
// 返回值分别为 *http.Response, 错误信息
func (h *HTTPClient) Req(method string, urlStr string, post interface{}, header map[string]string) (*http.Response, error) {
	var (
		req   *http.Request
		obody io.Reader
	)

	if post != nil {
		switch value := post.(type) {
		case io.Reader:
			obody = value
		case map[string]string, map[string]int, map[string]interface{}, []int, []string:
			postData, err := jsoniter.Marshal(value)
			if err != nil {
				return nil, err
			}
			header["Content-Type"] = "application/json"
			obody = bytes.NewReader(postData)
		case string:
			obody = strings.NewReader(value)
		case []byte:
			obody = bytes.NewReader(value)
		default:
			return nil, fmt.Errorf("request.Req: unknow post type: %s", post)
		}
	}

	req, err := http.NewRequest(method, urlStr, obody)
	if err != nil {
		return nil, err
	}

	// 设置浏览器标识
	req.Header.Set("User-Agent", UserAgent)

	if header != nil {
		for k, v := range header {
			req.Header.Set(k, v)
		}
	}

	h.SetTimeout(20 * time.Minute)

	return h.Client.Do(req)
}

// Fetch 实现 http／https 访问，
// 根据给定的 method (GET, POST, HEAD, PUT 等等),
// urlStr (网址),
// post (post 数据),
// header (header 请求头数据), 进行网站访问。
// 返回值分别为 []byte, 错误信息
func (h *HTTPClient) Fetch(method string, urlStr string, post interface{}, header map[string]string) ([]byte, error) {
	res, err := h.Req(method, urlStr, post, header)

	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}
