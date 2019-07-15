package cmds

import (
	"os"
	"strconv"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/service"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

//NewCourseCommand login command
func NewCourseCommand() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:      "column",
			Usage:     "获取专栏列表",
			UsageText: appName + " column",
			Action:    columnAction,
			Before:    authorizationFunc,
		},
		cli.Command{
			Name:      "video",
			Usage:     "获取视频课程列表",
			UsageText: appName + " video",
			Action:    videoAction,
			Before:    authorizationFunc,
		},
	}
}

func columnAction(c *cli.Context) error {
	columns, err := application.Columns()

	if err != nil {
		return err
	}

	renderCourses(columns)

	return nil
}

func videoAction(c *cli.Context) error {
	videos, err := application.Videos()

	if err != nil {
		return err
	}

	renderCourses(videos)

	return nil
}

func renderCourses(courses []*service.Course) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "名称", "购买", "作者"})

	for i, p := range courses {
		isBuy := ""
		if p.HadSub {
			isBuy = "是"
		}
		table.Append([]string{strconv.Itoa(i), strconv.Itoa(p.ID), p.ColumnTitle, isBuy, p.AuthorName})
	}

	table.Render()
}
