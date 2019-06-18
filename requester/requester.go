package requester

var (
	// UserAgent 浏览器标识
	UserAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1"

	//DefaultClient 默认 http 客户端
	DefaultClient = NewHTTPClient()
)
