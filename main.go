package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mmzou/geektime-dl/config"
)

func init() {

	err := config.Instance.Init()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// login := login.NewLoginClient()
	// phone := "13240929572"
	// password := "123123"
	// result := login.Login(phone, password)
	// if !result.IsLoginSuccess() {
	// 	fmt.Println(result.Error.Msg)
	// 	return
	// }

	// geektime := &config.Geektime{
	// 	User: config.User{
	// 		ID:     result.Data.UID,
	// 		Name:   result.Data.Name,
	// 		Avatar: result.Data.Avatar,
	// 	},
	// 	GCID:         result.Data.GCID,
	// 	GCESS:        result.Data.GCESS,
	// 	Ticket:       result.Data.Ticket,
	// 	ServerID:     result.Data.ServerID,
	// 	CookieString: result.Data.CookieString,
	// }

	// fmt.Println("geektime", geektime)

	// config.Instance.Geektimes = append(config.Instance.Geektimes, geektime)
	// config.Instance.AcitveUID = geektime.ID

	// err := config.Instance.Save()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(geektime)
	res, err := config.Instance.ActiveUserService().User()
	if err != nil {
		fmt.Println("error", err)
	}
	defer res.Body.Close()
	user, err := ioutil.ReadAll(res.Body)
	fmt.Println(string(user))
}
