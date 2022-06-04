package main

import (
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rhydianjenkins/cof/pkg/rainDrawer"
	"github.com/rhydianjenkins/cof/pkg/timeDrawer"
	"github.com/rivo/tview"
	"github.com/urfave/cli/v2"
)

const (
	FPS             = 30
	refreshInterval = time.Millisecond * 1000 / FPS
)

var (
	cliApp *cli.App
	view   *tview.Box
	app    *tview.Application
)

func init() {
	app = tview.NewApplication()

	cliApp = &cli.App{
		Name:        "cof",
		Version:     "v0.0.1",
		Description: "A collection of TUI programs",
		Compiled:    time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Rhydian",
				Email: "rhydz@msn.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:    "rain",
				Aliases: []string{"r"},
				Usage:   "Show some rain",
				Action: func(c *cli.Context) error {
					start(app, rainDrawer.Draw)
					return nil
				},
			},
			{
				Name:    "time",
				Aliases: []string{"t"},
				Usage:   "Show the current time",
				Action: func(c *cli.Context) error {
					start(app, timeDrawer.Draw)
					return nil
				},
			},
		},
		Flags: []cli.Flag{},
	}
}

type drawFn func(screen tcell.Screen, x, y, width, height int) (int, int, int, int)

func start(app *tview.Application, drawFn drawFn) {
	view = tview.NewBox().SetDrawFunc(drawFn)

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
