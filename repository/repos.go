package repository

import (
	"log"

	"github.com/raduq/proj/env"
	"github.com/urfave/cli"
)

// List user repositories
func List(env env.Connection) func(c *cli.Context) {
	return func(c *cli.Context) {
		repos, _, err := env.Client.Repositories.List(env.Context, c.String("user"), nil)

		if err != nil {
			log.Fatalln(err)
		}

		for _, repo := range repos {
			log.Println(repo.GetFullName())
		}
	}
}

// List repo projects
func Projects(env env.Connection) func(c *cli.Context) {
	return func(c *cli.Context) {
		projs, _, err := env.Client.Repositories.ListProjects(env.Context, c.String("owner"), c.String("repo"), nil)

		if err != nil {
			log.Fatalln(err)
		}

		for _, proj := range projs {
			log.Printf("ID: %v, Name: %v\n", proj.GetID(), proj.GetName())
		}
	}
}
