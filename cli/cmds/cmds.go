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
	_debug         bool
	_info          bool
	_stream        string
	_pdf           bool
	_mp3           bool
	_markdown      bool
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

// NewApp cli app
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
			Name:        "debug,d",
			Usage:       "Turn on debug logs",
			Destination: &_debug,
		},
		cli.BoolFlag{
			Name:        "info, i",
			Usage:       "只输出视频信息",
			Destination: &_info,
		},
		cli.StringFlag{
			Name:        "stream, s",
			Usage:       "选择要下载的指定类型",
			Destination: &_stream,
		},
		cli.BoolFlag{
			Name:        "pdf, p",
			Usage:       "下载专栏PDF文档",
			Destination: &_pdf,
		},
		cli.BoolFlag{
			Name:        "mp3, m",
			Usage:       "下载专栏MP3音频",
			Destination: &_mp3,
		},
		cli.BoolFlag{
			Name:        "markdown, md",
			Usage:       "下载专栏markdown文档",
			Destination: &_markdown,
		},
	}

	app.Before = func(c *cli.Context) error {
		if _debug {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	return app
}

// DefaultAction default action
func DefaultAction(c *cli.Context) error {
	if len(c.Args()) == 0 {
		err := cli.ShowAppHelp(c)
		return err
	}

	dlc := &NewDownloadCommand()[0]
	if dlc != nil {
		return dlc.Run(c)
	}

	return nil
}
