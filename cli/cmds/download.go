package cmds

import (
	"fmt"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/urfave/cli"
)

//NewDownloadCommand login command
func NewDownloadCommand() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:      "download",
			Usage:     "下载",
			UsageText: appName + " download",
			Action:    downloadAction,
			Before:    authorizationFunc,
		},
	}
}

func downloadAction(c *cli.Context) error {
	course, articles, err := application.CourseWithArticles(186)
	fmt.Println(course, articles, err)

	for k, v := range articles {
		fmt.Println(k, v)
	}

	if err != nil {
		return err
	}

	return nil
}
