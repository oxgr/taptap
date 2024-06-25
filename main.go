package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "watch",
				Aliases: []string{"s"},
				Usage:   "start a tap tempo watcher",
				Action: func(cCtx *cli.Context) error {
					return watch()
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
