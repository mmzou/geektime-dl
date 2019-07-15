package service

import (
	"io"
	"net/http"
)

func deferResponseClose(s *http.Response) {
	if s != nil {
		defer s.Body.Close()
	}
}

//handleHTTPResponse handle
func handleHTTPResponse(res *http.Response, err error) (io.ReadCloser, Error) {
	if err != nil {
		deferResponseClose(res)
		return nil, &ErrorInfo{Err: err}
	}

	if res.StatusCode == 452 {
		return nil, &ErrorInfo{Err: ErrLoginOffline}
	}

	return res.Body, nil
}
