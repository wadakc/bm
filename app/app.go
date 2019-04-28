package app

import (
	"github.com/codegangsta/cli"
	"github.com/kcwebapply/bm/commands"
)

var appName = "bm"
var version = "0.0.1"

// InitApp method is for initializing and  getting App settings.
func InitApp() *cli.App {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = "github.com/kcwebapply/bm"
	app.Version = version
	// command routing.
	app.Commands = commands.Commands()

	app.Before = func(c *cli.Context) error {
		return nil
	}

	app.After = func(c *cli.Context) error {
		return nil
	}
	return app
}
