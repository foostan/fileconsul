package command

import (
	"github.com/codegangsta/cli"
	"log"

	. "github.com/foostan/fileconsul/fileconsul"
	"path/filepath"
)

var PushFlags = []cli.Flag{
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

func PushCommand(c *cli.Context) {
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
		println("push new file:\t" + fh.Path)
		err = client.PutKV(filepath.Join(prefix, "hash", fh.Path), fh.Hash)
		if err != nil {
			log.Fatal(err)
		}

		err = client.PutKV(filepath.Join(prefix, "host", fh.Path), fh.Host)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, fh := range delFhs {
		println("delete file:\t" + fh.Path)
		err = client.DeleteKV(filepath.Join(prefix, "hash", fh.Path))
		if err != nil {
			log.Fatal(err)
		}

		err = client.DeleteKV(filepath.Join(prefix, "host", fh.Host))
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, fh := range modFhs {
		println("modify file:\t" + fh.Path)
		err = client.PutKV(filepath.Join(prefix, "hash", fh.Path), fh.Hash)
		if err != nil {
			log.Fatal(err)
		}

		err = client.PutKV(filepath.Join(prefix, "host", fh.Path), fh.Host)
		if err != nil {
			log.Fatal(err)
		}
	}
}
