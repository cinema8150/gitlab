package command

import (
	"fmt"
	"gitlab/config"

	"github.com/urfave/cli/v2"
)

func ConfigCommand() *cli.Command {
	return &cli.Command{
		Name:  "config",
		Usage: "config gitlab settings.",
		Subcommands: []*cli.Command{
			getCommand(),
			setCommand(),
		},
	}
}

//TODO: 设置默认使用当前项目host还是globe.host,默认优先用当前项目信息
func getCommand() *cli.Command {
	return &cli.Command{
		Name:  "get",
		Usage: "string",
		Action: func(c *cli.Context) error {
			value := "unfound value for this key."
			key := c.Args().Get(0)
			switch key {
			case "host":
				info := config.Shared().Host
				if len(info) > 0 {
					value = info
				}
			case "token":
				info := config.Shared().Token
				if len(info) > 0 {
					value = info
				}
			}
			fmt.Printf("%s value is : %s\n", key, value)

			// //TODO: token获取加入密码验证
			// token := c.String("token")
			// if len(token) > 0 {
			// 	fmt.Printf("globe.token:%s\n", config.Shared().Token)
			// 	return nil
			// }

			// host := c.String("host")
			// if len(host) > 0 {
			// 	fmt.Printf("globe.host:%s\n", config.Shared().Host)
			// 	return nil
			// }

			return nil
		},
		// Flags: []cli.Flag{
		// 	&cli.StringFlag{
		// 		Name:  "host",
		// 		Usage: "the defualt repo host",
		// 		Value: "globe",
		// 	},
		// 	&cli.StringFlag{
		// 		Name:  "token",
		// 		Usage: "the token for default host",
		// 	},
		// 	&cli.StringFlag{
		// 		Name:  "all",
		// 		Usage: "all repo host.",
		// 	},
		// },
	}
}

//TODO: 设置默认使用当前项目host还是globe.host,默认优先用当前项目信息
func setCommand() *cli.Command {
	return &cli.Command{
		Name:  "set",
		Usage: "string",
		Action: func(c *cli.Context) error {
			token := c.String("token")
			if len(token) > 0 {
				config.Shared().Token = token
			}
			host := c.String("host")
			if len(host) > 0 {
				config.Shared().Host = host
			}
			config.Shared().Update()
			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Usage: "the defualt repo host",
			},
			&cli.StringFlag{
				Name:  "token",
				Usage: "the token for default host",
			},
		},
	}
}
