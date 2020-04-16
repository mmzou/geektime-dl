package cmds

import (
	"os"
	"strconv"
	"time"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/mmzou/geektime-dl/service"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

//NewCourseCommand login command
func NewCourseCommand() []cli.Command {
	return []cli.Command{
		{
			Name:      "column",
			Usage:     "获取专栏列表",
			UsageText: appName + " column",
			Action:    columnAction,
		},
		{
			Name:      "video",
			Usage:     "获取视频课程列表",
			UsageText: appName + " video",
			Action:    videoAction,
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
	table.SetHeader([]string{"#", "ID", "名称", "时间", "作者", "购买"})
	table.SetAutoWrapText(false)

	for i, p := range courses {
		isBuy := ""
		if p.HadSub {
			isBuy = "是"
		}
		table.Append([]string{strconv.Itoa(i), strconv.Itoa(p.ID), p.ColumnTitle, time.Unix(int64(p.ColumnCtime), 0).Format("2006-01-02"), p.AuthorName, isBuy})
	}

	table.Render()
}
