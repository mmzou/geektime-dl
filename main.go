package main

import (
	"fmt"

	"github.com/mmzou/geektime-dl/login"
)

func main() {
	login := login.NewLoginClient()

	phone := "13240929572"
	password := "111111"
	result := login.Login(phone, password)

	fmt.Println(result)
}
