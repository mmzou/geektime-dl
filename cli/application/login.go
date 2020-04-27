package application

import (
	"errors"

	"github.com/mmzou/geektime-dl/login"
)

//Login login
func Login(phone, password string) (gcess string, gcid string, serverID string, err error) {
	c := login.NewLoginClient()
	result := c.Login(phone, password)
	if !result.IsLoginSuccess() {
		return "", "", "", errors.New(result.Error.Msg)
	}

	return result.Data.GCID, result.Data.GCESS, result.Data.ServerID, nil
}

//LoginedCookies get logined cookies
func LoginedCookies() map[string]string {
	return getService().Cookies()
}
