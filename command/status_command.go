package command

import (
	"github.com/codegangsta/cli"
	"log"

	. "github.com/foostan/fileconsul/fileconsul"
)

var StatusFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "base-dir",
		Usage: "base directory of target files",
	},
	cli.StringFlag{
		Name:  "addr",
		Value: "127.0.0.1:8500",
		Usage: "consul HTTP API address with port",
	},
	cli.StringFlag{
		Name:  "dc",
		Value: "local",
		Usage: "consul datacenter, uses local if blank",
	},
}

func StatusCommand(c *cli.Context) {
	baseDir := c.String("base-dir")
	if baseDir == "" {
		log.Fatalf("Error missing flag 'base-dir'")
	}
	addr := c.String("addr")
	dc := c.String("dc")

	client, err := NewClient(&ClientConfig{
		ConsulAddr: addr,
		ConsulDC:   dc,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = client.GetFileStatus(baseDir)
	if err != nil {
		log.Fatal(err)
	}
}
