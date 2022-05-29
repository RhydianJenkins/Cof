package main

import (
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli/v2"
)

var app *cli.App

func init() {
	app = &cli.App{
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
					fmt.Println("TODO")
					return nil
				},
			},
		},
		Flags: []cli.Flag{},
	}
}

func main() {
	app.Run(os.Args)
}
