package service

import "net/http"

func deferResponseClose(s *http.Response) {
	if s != nil {
		defer s.Body.Close()
	}
}
