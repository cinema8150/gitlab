package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "gitlab"
	app.Version = "0.1.0"
	app.Usage = "gl <command> <subcommand> [flags]"
	app.Commands = []*cli.Command{
		{
			Name:  "version",
			Usage: "show gitlab version",
			Action: func(c *cli.Context) error {
				fmt.Println("v0.1.0")
				return nil
			},
		},
	}
	app.Run(os.Args)
}
