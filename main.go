package main

import (
	"os"

	"github.com/raduq/proj/env"
	"github.com/raduq/proj/organization"
	"github.com/raduq/proj/project"
	"github.com/raduq/proj/repository"
	"github.com/urfave/cli"
)

func main() {
	env := env.Connect()

	app := cli.NewApp()
	app.Name = "proj"
	app.Usage = "Managed github projects"
	app.Commands = []cli.Command{
		{
			Name:    "repos",
			Aliases: []string{"r", "repositories"},
			Usage:   "List repositories",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user, u",
					Usage: "set the user to the action",
				},
			},
			Action: repository.List(env),
		},
		{
			Name:    "repos-projs",
			Aliases: []string{"rp", "repository-projects"},
			Usage:   "List repository projects",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "owner, o",
					Usage: "set the name of the organization's/user's owner of the repository",
				},
				cli.StringFlag{
					Name:  "repo, r",
					Usage: "set the name of the repository",
				},
			},
			Action: repository.Projects(env),
		},
		{
			Name:    "orgs",
			Aliases: []string{"o", "organizations"},
			Usage:   "List organizations",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "user, u",
					Usage: "set the user to the action",
				},
			},
			Action: organization.List(env),
		},
		{
			Name:    "orgs-projs",
			Aliases: []string{"op", "organization-projects"},
			Usage:   "List repository projects",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "org, o",
					Usage: "set the name of the organization",
				},
			},
			Action: organization.Projects(env),
		},
		{
			Name:    "cols",
			Aliases: []string{"cl", "columns"},
			Usage:   "List project columns",
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "id, i",
					Usage: "set the ID of the project",
				},
			},
			Action: project.Columns(env),
		},
		{
			Name:    "cards",
			Aliases: []string{"c", "cards"},
			Usage:   "List project column's cards",
			Flags: []cli.Flag{
				cli.Int64Flag{
					Name:  "id, i",
					Usage: "set the ID of the column",
				},
			},
			Action: project.Cards(env),
		},
	}
	app.Version = "1.0"
	app.Run(os.Args)
}
