package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

func init() {
}

func main() {
	app := &cli.App{
		Name:     "cof",
		Version:  "v0.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Rhydian",
				Email: "rhydz@msn.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "add",
				Aliases: []string{"a"},
				Usage:   "Add a task to today",
				Action: func(c *cli.Context) error {
					fmt.Print("TODO")
					return nil
				},
			},
		},
		Flags: []cli.Flag{},
	}

	app.Run(os.Args)
}
