package service

import (
	"errors"
	"io"

	"github.com/mmzou/geektime-dl/utils"
)

type resultData []byte

func (rd *resultData) UnmarshalJSON(data []byte) error {
	*rd = data

	return nil
}

func (rd resultData) String() string {
	return string(rd)
}

// Result 从百度服务器解析的数据结构
type Result struct {
	Code  int        `json:"code"`
	Data  resultData `json:"data"`
	Error struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	} `json:"error"`
	Extra struct {
		Cost      float64 `json:"cost"`
		RequestID string  `json:"request-id"`
	} `json:"extra"`
}

func (r *Result) isSuccess() bool {
	return r.Code == 0
}

//User user info
type User struct {
	UID       int    `json:"uid"`
	Nickname  string `json:"nickname"`
	Avatar    string `json:"avatar"`
	Cellphone string `json:"cellphone"`
}

func handleJSONParse(reader io.Reader, v interface{}) Error {
	result := new(Result)

	err := utils.UnmarshalReader(reader, &result)
	if err != nil {
		return &ErrorInfo{Err: err}
	}

	if !result.isSuccess() {
		//未登录或者登录凭证无效
		if result.Error.Code == -3050 {
			return &ErrorInfo{Err: ErrNotLogin}
		}
		return &ErrorInfo{Err: errors.New(result.Error.Msg)}
	}

	err = utils.UnmarshalJSON(result.Data, v)
	if err != nil {
		return &ErrorInfo{Err: err}
	}

	return nil
}
