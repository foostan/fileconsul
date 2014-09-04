package command

import (
	"github.com/codegangsta/cli"
	"log"

	. "github.com/foostan/fileconsul/fileconsul"
)

var StatusFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "addr",
		Value: "127.0.0.1:8500",
		Usage: "consul HTTP API address with port",
	},
	cli.StringFlag{
		Name:  "dc",
		Value: "",
		Usage: "consul datacenter, uses local if blank",
	},
	cli.StringFlag{
		Name:  "prefix",
		Value: "fileconsul",
		Usage: "reading file status from Consul's K/V store with the given prefix",
	},
	cli.StringFlag{
		Name:  "base-dir",
		Value: ".",
		Usage: "base directory of target files",
	},
}

func StatusCommand(c *cli.Context) {
	addr := c.String("addr")
	dc := c.String("dc")
	prefix := c.String("prefix")
	baseDir := c.String("base-dir")

	client, err := NewClient(&ClientConfig{
		ConsulAddr: addr,
		ConsulDC:   dc,
	})
	if err != nil {
		log.Fatal(err)
	}

	err = client.GetFileStatus(prefix, baseDir)
	if err != nil {
		log.Fatal(err)
	}
}
