package organization

import (
	"log"

	"github.com/raduq/proj/env"
	"github.com/urfave/cli"
)

// List user organizations
func List(env env.Connection) func(c *cli.Context) {
	return func(c *cli.Context) {
		orgs, _, err := env.Client.Organizations.List(env.Context, c.String("user"), nil)

		if err != nil {
			log.Fatalln(err)
		}

		for _, org := range orgs {
			log.Println(org.GetLogin())
		}
	}
}

// List org projects
func Projects(env env.Connection) func(c *cli.Context) {
	return func(c *cli.Context) {
		projs, _, err := env.Client.Organizations.ListProjects(env.Context, c.String("org"), nil)

		if err != nil {
			log.Fatalln(err)
		}

		for _, proj := range projs {
			log.Printf("ID: %v, Name: %v\n", proj.GetID(), proj.GetName())
		}
	}
}
