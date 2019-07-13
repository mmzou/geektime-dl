package cmds

import (
	"os"
	"strconv"

	"github.com/mmzou/geektime-dl/cli/application"
	"github.com/olekukonko/tablewriter"
	"github.com/urfave/cli"
)

//NewProductCommand login command
func NewProductCommand() []cli.Command {
	return []cli.Command{
		cli.Command{
			Name:      "product",
			Usage:     "Geektime all products",
			UsageText: appName + " product [OPTIONS]",
			Action:    productAction,
			Before:    authorizationFunc,
		},
	}
}

func productAction(c *cli.Context) error {
	products, err := application.ProductAll()

	if err != nil {
		return err
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "ID", "类型", "名称", "作者"})

	i := 0
	for _, p := range products.Column.List {
		table.Append([]string{strconv.Itoa(i), strconv.Itoa(p.Extra.LastAid), products.Column.Title, p.Title, p.Extra.AuthorName})
		i++
	}

	for _, p := range products.Course.List {
		table.Append([]string{strconv.Itoa(i), strconv.Itoa(p.Extra.ColumnID), products.Course.Title, p.Title, p.Extra.AuthorName})
		i++
	}

	table.Render()

	return nil
}
