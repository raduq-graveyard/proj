package project

import (
	"log"

	"github.com/raduq/proj/env"
	"github.com/urfave/cli"
)

// Columns of the project
func Columns(env env.Connection) func(c *cli.Context) {
	return func(c *cli.Context) {
		cols, _, err := env.Client.Projects.ListProjectColumns(env.Context, c.Int64("id"), nil)

		if err != nil {
			log.Fatalln(err)
		}

		for _, col := range cols {
			log.Printf("ID: %v, Name: %v\n", col.GetID(), col.GetName())
		}
	}
}

// Cards of the project's column
func Cards(env env.Connection) func(c *cli.Context) {
	return func(c *cli.Context) {
		cards, _, err := env.Client.Projects.ListProjectCards(env.Context, c.Int64("id"), nil)

		if err != nil {
			log.Fatalln(err)
		}

		for _, cd := range cards {
			log.Printf("ID: %v, Note: %v\n", cd.GetID(), cd.GetNote())
		}
	}
}
