package command

import (
	"github.com/codegangsta/cli"
	"log"

	. "github.com/foostan/fileconsul/fileconsul"
)

var StatusFlags = []cli.Flag{
	cli.StringFlag{
		Name:  "addr",
		Value: "localhost:8500",
		Usage: "consul HTTP API address with port",
	},
	cli.StringFlag{
		Name:  "dc",
		Value: "dc1",
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

	localFhs, err := LocalFileHashs(baseDir)
	if err != nil {
		log.Fatal(err)
	}

	remoteFhs, err := RemoteFileHashs(client, prefix)
	if err != nil {
		log.Fatal(err)
	}

	addFhs, delFhs, modFhs := DiffFileHashs(localFhs, remoteFhs)
	for _, fh := range addFhs {
		println("new file:\t" + fh.Path)
	}
	for _, fh := range delFhs {
		println("deleted:\t" + fh.Path)
	}
	for _, fh := range modFhs {
		println("modified:\t" + fh.Path)
	}
}
