package main

import (
	"fmt"

	"github.com/mmzou/geektime-dl/config"
)

func init() {
	err := config.Config.Init()
	if err != nil {
		fmt.Println(err)
	}
	err = config.Config.Save()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	// login := login.NewLoginClient()

	// phone := "13240929572"
	// password := "111111"
	// result := login.Login(phone, password)

	// fmt.Println(result)
}
