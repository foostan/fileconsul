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
		Name:  "basepath",
		Value: ".",
		Usage: "base directory path of target files",
	},
}

func StatusCommand(c *cli.Context) {
	addr := c.String("addr")
	dc := c.String("dc")
	prefix := c.String("prefix")
	basepath := c.String("basepath")

	client, err := NewClient(&ClientConfig{
		ConsulAddr: addr,
		ConsulDC:   dc,
	})
	if err != nil {
		log.Fatal(err)
	}

	lfList, err := ReadLFList(basepath)
	if err != nil {
		log.Fatal(err)
	}

	mfList, err := client.ReadMFList(prefix)
	if err != nil {
		log.Fatal(err)
	}

	mfDiff := mfList.Diff(lfList.ToMFList())

	for _, metafile := range mfDiff.Eq {
		println("remote/local:\t" + metafile.Path)
	}

	for _, metafile := range mfDiff.Add {
		println("remote:\t" + metafile.Path)
	}

	for _, metafile := range mfDiff.Del {
		println("local:\t" + metafile.Path)
	}

	for _, metafile := range mfDiff.New {
		println("remote/local:(modified)\t" + metafile.Path)
	}
}
