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
		Name:  "path",
		Usage: "pushed file path, full file path is `prefix + path` in K/V store",
	},
	cli.StringFlag{
		Name:  "url",
		Usage: "pushed file url",
	},
}

func PushCommand(c *cli.Context) {
	addr := c.String("addr")
	dc := c.String("dc")
	prefix := c.String("prefix")
	path := c.String("path")
	if path == "" {
		log.Fatalf("Error missing flag 'path'")
	}
	url := c.String("url")
	if url == "" {
		log.Fatalf("Error missing flag 'url'")
	}

	client, err := NewClient(&ClientConfig{
		ConsulAddr: addr,
		ConsulDC:   dc,
	})
	if err != nil {
		log.Fatal(err)
	}

	hash, err:= UrlToHash(url)
	if err != nil {
		log.Fatal(err)
	}

	rfValue := RFValue{Url: url, Hash: hash}

	err = client.PutKV(filepath.Join(prefix, path), rfValue.ToStr())
	if err != nil {
		log.Fatal(err)
	}
}
