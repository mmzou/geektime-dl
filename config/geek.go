package config

import "github.com/mmzou/geektime-dl/service"

//User geek time user info
type User struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
}

//Geektime geek time info
type Geektime struct {
	User
	GCID         string `json:"gcid"`
	GCESS        string `json:"gcess"`
	ServerID     string `json:"serverId"`
	Ticket       string `json:"ticket"`
	CookieString string `json:"cookieString"`
}

//Service geek time service
func (g *Geektime) Service() *service.Service {
	ser := service.NewService(g.GCID, g.GCESS, g.ServerID)

	return ser
}
