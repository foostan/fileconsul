package command

import (
	"log"

	"github.com/codegangsta/cli"
)

var StatusFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "base-dir",
		Usage: "base directory of target files",
	},
}

func StatusCommand(c *cli.Context) {
	base_dir := c.String("base-dir")
	if base_dir == "" {
		log.Fatalf("Error missing flag 'base-dir'")
	}
}
