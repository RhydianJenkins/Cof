package main

import (
	"os"
	"time"

	"github.com/rhydianjenkins/cof/pkg/timeDrawer"
	"github.com/rivo/tview"
	"github.com/urfave/cli/v2"
)

const refreshInterval = 500 * time.Millisecond

var (
	cliApp *cli.App
	view   *tview.Box
	app    *tview.Application
)

func init() {
	app = tview.NewApplication()

	cliApp = &cli.App{
		Name:     "cof",
		Version:  "v0.0.1",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Rhydian",
				Email: "rhydz@msn.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "time",
				Aliases: []string{"t"},
				Usage:   "Show the current time",
				Action: func(c *cli.Context) error {
					showTime(app)
					return nil
				},
			},
		},
		Flags: []cli.Flag{},
	}
}

func showTime(app *tview.Application) {
	view = tview.NewBox().SetDrawFunc(timeDrawer.Draw)

	go tickLoop()

	if err := app.SetRoot(view, true).Run(); err != nil {
		panic(err)
	}
}

func tickLoop() {
	tick := time.NewTicker(refreshInterval)

	for {
		select {
		case <-tick.C:
			app.Draw()
		}
	}
}

func main() {
	cliApp.Run(os.Args)
}
