package command

import (
	"github.com/codegangsta/cli"
	"log"
	"fmt"
	"path/filepath"

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

	rfList, err := client.ReadRFList(prefix)
	if err != nil {
		log.Fatal(err)
	}

	lfrfList := lfList.ToRFList(prefix)
	rfDiff := lfrfList.Diff(rfList)

	if !lfrfList.Equal(rfList) {
		fmt.Println("Changes to be pushed:\n  (use \"fileconsul pull [command options]\" to synchronize remote files)")
	}

	for _, remotefile := range rfDiff.Add {
		println("\tnew file:\t" + filepath.Join(basepath, remotefile.Path))
	}

	for _, remotefile := range rfDiff.Del {
		println("\tdeleted:\t" + filepath.Join(basepath, remotefile.Path))
	}

	for _, remotefile := range rfDiff.New {
		println("\tmodified:\t" + filepath.Join(basepath, remotefile.Path))
	}
}
