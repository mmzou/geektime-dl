package config

import (
	"errors"

	"github.com/mmzou/geektime-dl/service"
)

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

//SetUserByGcidAndGcess set user
func (c *ConfigsData) SetUserByGcidAndGcess(gcid, gcess, serverID string) (*Geektime, error) {
	ser := service.NewService(gcid, gcess, serverID)
	user, err := ser.User()
	if err != nil {
		return nil, err
	}

	c.DeleteUser(&User{ID: user.UID})

	geektime := &Geektime{
		User: User{
			ID:     user.UID,
			Name:   user.Nickname,
			Avatar: user.Avatar,
		},
		GCID:     gcid,
		GCESS:    gcess,
		ServerID: serverID,
	}

	c.Geektimes = append(c.Geektimes, geektime)

	c.setActiveUser(geektime)

	return geektime, nil
}

//DeleteUser delete
func (c *ConfigsData) DeleteUser(u *User) {
	for k, gk := range c.Geektimes {
		if gk.ID == u.ID {
			c.Geektimes = append(c.Geektimes[:k], c.Geektimes[k+1:]...)
			break
		}
	}
}

func (c *ConfigsData) setActiveUser(g *Geektime) {
	c.AcitveUID = g.ID
	c.activeUser = g
}

//LoginUserCount 登录用户数量
func (c *ConfigsData) LoginUserCount() int {
	return len(c.Geektimes)
}

//SwitchUser switch user
func (c *ConfigsData) SwitchUser(u *User) error {
	for _, g := range c.Geektimes {
		if g.ID == u.ID {
			c.setActiveUser(g)
			return nil
		}
	}

	return errors.New("用户不存在")
}
