package main

import (
	"fmt"
	"os"

	"github.com/mmzou/geektime-dl/cli/cmds"
	"github.com/mmzou/geektime-dl/config"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func init() {

	err := config.Instance.Init()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	app := cmds.NewApp()
	app.Commands = []cli.Command{}
	app.Commands = append(app.Commands, cmds.NewLoginCommand()...)
	app.Commands = append(app.Commands, cmds.NewBuyCommand()...)
	app.Commands = append(app.Commands, cmds.NewCourseCommand()...)

	app.Action = cmds.DefaultAction

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
