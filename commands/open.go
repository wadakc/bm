package commands

import (
	"os"
	"strconv"

	"github.com/codegangsta/cli"
	"github.com/skratchdot/open-golang/open"
)

// OpenPage open bookmark website
func OpenPage(c *cli.Context) {
	id := c.Args().Get(0)
	if id == "" {
		os.Exit(0)
	}
	openPage(id)
}

func openPage(id string) {
	pages := readLines()
	for _, page := range pages {
		if id == strconv.Itoa(page.ID) {
			open.Run(page.URL)
		}
	}
}