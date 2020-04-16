package cmds

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/config"
	"github.com/olekukonko/tablewriter"
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
func NewLoginCommand() []cli.Command {
	return []cli.Command{
		{
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
		},
		{
			Name:        "who",
			Usage:       "获取当前帐号",
			UsageText:   appName + " who",
			Description: "获取当前帐号的信息",
			Action:      whoAction,
			Before:      authorizationFunc,
		},
		{
			Name:        "users",
			Usage:       "获取帐号列表",
			UsageText:   appName + " users",
			Description: "获取当前已登录的帐号列表",
			Action:      usersAction,
		},
		{
			Name:        "su",
			Usage:       "切换极客时间帐号",
			UsageText:   appName + " su [UID]",
			Description: "切换已登录的极客时间账号",
			Action:      suAction,
			After:       configSaveFunc,
		},
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

func whoAction(c *cli.Context) error {
	activeUser := config.Instance.ActiveUser()
	fmt.Printf("当前帐号 uid: %d, 用户名: %s, 头像地址: %s \n", activeUser.ID, activeUser.Name, activeUser.Avatar)

	return nil
}

func usersAction(c *cli.Context) error {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "UID", "昵称"})
	i := 0
	for _, g := range config.Instance.Geektimes {
		table.Append([]string{strconv.Itoa(i), strconv.Itoa(g.ID), g.Name})
		i++
	}
	table.Render()

	return nil
}

func suAction(c *cli.Context) error {
	if config.Instance.LoginUserCount() == 0 {
		return errors.New("未登录任何极客时间账号，不能切换")
	}

	if c.NArg() == 0 {
		cli.HandleAction(usersAction, c)
		return errors.New("请选择登录的用户UID")
	}

	i, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		cli.HandleAction(usersAction, c)
		return errors.New("请输入用户UID")
	}

	err = config.Instance.SwitchUser(&config.User{ID: i})

	if err != nil {
		return err
	}

	fmt.Printf("成功切换登录用户：%s\n", config.Instance.ActiveUser().Name)

	return nil
}
