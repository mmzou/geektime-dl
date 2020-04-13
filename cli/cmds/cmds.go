package cmds

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"

	"github.com/mmzou/geektime-dl/cli/version"
	"github.com/mmzou/geektime-dl/config"
	"github.com/urfave/cli"
)

var (
	debug          bool
	appName        = filepath.Base(os.Args[0])
	configSaveFunc = func(c *cli.Context) error {
		err := config.Instance.Save()
		if err != nil {
			return errors.New("保存配置错误：" + err.Error())
		}
		return nil
	}
	authorizationFunc = func(c *cli.Context) error {
		if config.Instance.AcitveUID <= 0 {
			if len(config.Instance.Geektimes) > 0 {
				return config.ErrHasLoginedNotLogin
			}
			return config.ErrNotLogin
		}

		return nil
	}
)

//NewApp cli app
func NewApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "极客时间下载客户端"
	app.Version = fmt.Sprintf("%s", version.Version)
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("%s version %s\n", app.Name, app.Version)
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:        "debug",
			Usage:       "Turn on debug logs",
			Destination: &debug,
		},
		cli.BoolFlag{
			Name:  "info, i",
			Usage: "只输出视频信息",
		},
	}

	app.Before = func(c *cli.Context) error {
		if debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	return app
}

//DefaultAction default action
func DefaultAction(c *cli.Context) error {
	if len(c.Args()) == 0 {
		cli.ShowAppHelp(c)
		return nil
	}

	dlc := &NewDownloadCommand()[0]
	if dlc != nil {
		return dlc.Run(c)
	}

	return nil
}
