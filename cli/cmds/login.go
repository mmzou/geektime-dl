package cmds

import (
	"errors"
	"fmt"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/config"
	"github.com/urfave/cli"
)

//Login login data
type Login struct {
	phone    string
	password string
	gcid     string
	gcess    string
	serverID string
}

//IsByPhoneAndPassword 通过手机号和密码登录
func (l *Login) IsByPhoneAndPassword() bool {
	return l.phone != "" && l.password != ""
}

//IsByCookie cookie login
func (l *Login) IsByCookie() bool {
	return l.gcid != "" && l.gcess != ""
}

//LoginConfig config
var LoginConfig Login

//NewLoginCommand login command
func NewLoginCommand() cli.Command {
	return cli.Command{
		Name:      "login",
		Usage:     "Login geektime",
		UsageText: appName + " login [OPTIONS]",
		Action:    loginAction,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "phone",
				Usage:       "登录手机号",
				Destination: &LoginConfig.phone,
			},
			cli.StringFlag{
				Name:        "password",
				Usage:       "登录密码",
				Destination: &LoginConfig.password,
			},
			cli.StringFlag{
				Name:        "gcid",
				Usage:       "GCID Cookie",
				Destination: &LoginConfig.gcid,
			},
			cli.StringFlag{
				Name:        "gcess",
				Usage:       "GCESS Cookie",
				Destination: &LoginConfig.gcess,
			},
			cli.StringFlag{
				Name:        "serverId",
				Usage:       "SERVERID Cookie",
				Destination: &LoginConfig.serverID,
			},
		},
		After: configSaveFunc,
	}
}

func loginAction(c *cli.Context) error {
	//通过手机号和密码登录
	var (
		gcid     = LoginConfig.gcid
		gcess    = LoginConfig.gcess
		serverID = LoginConfig.serverID
		err      error
	)
	if LoginConfig.IsByPhoneAndPassword() {
		gcid, gcess, serverID, err = application.Login(LoginConfig.phone, LoginConfig.password)
		if err != nil {
			return err
		}
	}

	if gcid != "" && gcess != "" {
		geektime, err := config.Instance.SetUserByGcidAndGcess(gcid, gcess, serverID)
		if err != nil {
			return err
		}

		fmt.Println("极客时间账号登录成功：", geektime.Name)

		return nil
	}

	return errors.New("请输入登录凭证信息")
}
