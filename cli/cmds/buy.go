package cmds

import (
	"os"
	"strconv"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

//NewBuyCommand login command
func NewBuyCommand() []cli.Command {
	return []cli.Command{
		{
			Name:      "buy",
			Usage:     "获取已购买过的专栏和视频课程",
			UsageText: appName + " buy",
			Action:    buyAction,
			Before:    authorizationFunc,
		},
	}
}

func buyAction(c *cli.Context) error {
	products, err := application.BuyProductAll()

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "类型", "名称", "作者"})

	i := 0
	for _, p := range products.Columns.List {
		table.Append([]string{strconv.Itoa(i), strconv.Itoa(p.Extra.ColumnID), products.Columns.Title, p.Title, p.Extra.AuthorName})
		i++
	}

	for _, p := range products.Videos.List {
		table.Append([]string{strconv.Itoa(i), strconv.Itoa(p.Extra.ColumnID), products.Videos.Title, p.Title, p.Extra.AuthorName})
		i++
	}

	table.Render()

	return nil
}
