package command

import (
	"fmt"
	"gitlab/git"
	"gitlab/project"
	"gitlab/user"

	"github.com/urfave/cli/v2"
)

func ProjectCommand() *cli.Command {
	return &cli.Command{
		Name: "project",
		Subcommands: []*cli.Command{
			searchCommand(),
		},
	}
}

func searchCommand() *cli.Command {
	return &cli.Command{
		Name:  "search",
		Usage: "string",
		Action: func(c *cli.Context) error {

			projectName := c.String("project")
			merge := c.String("merge")
			if len(merge) > 0 {
				repoName := projectName
				if len(repoName) == 0 {
					// 优先使用本地缓存读取当前项目ID
					repoID, err := git.CurrentRepoID()
					if err != nil {
						fmt.Println(err)
						return nil
					}
					if len(repoID) > 0 {
						list, err := project.SearchMergeByProjectID(repoID, merge)
						if err != nil {
							println(err.Error())
							return err
						}
						fmt.Printf("search emrge request result :%v\n", list)
						return nil
					}
					// 读取当前项目名
					repoName, err = git.CurrentRepoName()
					if err != nil {
						fmt.Printf("current folder is not a git repository, and not assign project name which you want search.")
						return err
					}
				}

				list, err := project.SearchMergeByProjectName(projectName, merge)
				if err != nil {
					println(err.Error())
					return err
				}
				fmt.Printf("search emrge request result :%v\n", list)
				return nil
			}

			if len(projectName) > 0 {
				_, err := project.Search(projectName)
				if err != nil {
					println(err.Error())
					return err
				}
				return nil
			}

			userName := c.String("user")
			if len(userName) > 0 {
				_, err := user.Search(userName)
				if err != nil {
					println(err.Error())
					return err
				}
				return nil
			}

			name := c.Args().First()
			if len(name) == 0 {
				println("please enter project name, which you want search. Or show more usage by '-h'")
				return nil
			}

			_, err := project.Search(name)
			if err != nil {
				println(err.Error())
				return err
			}

			return nil
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "project",
				Usage: "the defualt repo host",
			},
			&cli.StringFlag{
				Name:  "user",
				Usage: "the token for default host",
			},
			&cli.StringFlag{
				Name:  "merge",
				Usage: "the token for default host",
			},
		},
	}
}
