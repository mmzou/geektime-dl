package requester

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

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

	return h.Client.Do(req)
}

// Fetch 实现 http／https 访问，
// 根据给定的 method (GET, POST, HEAD, PUT 等等),
// urlStr (网址),
// post (post 数据),
// header (header 请求头数据), 进行网站访问。
// 返回值分别为 *http.Response, 错误信息
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
