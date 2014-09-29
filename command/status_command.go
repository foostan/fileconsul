package command

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/codegangsta/cli"

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
	args := c.Args()
	pattern := args.First()

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

	if pattern != "" {
		rfList, err = rfList.Filter(pattern)
		if err != nil {
			log.Fatal(err)
		}

		lfrfList, err = lfrfList.Filter(pattern)
		if err != nil {
			log.Fatal(err)
		}
	}

	rfDiff := lfrfList.Diff(rfList)
	switch {
	case lfrfList.Equal(rfList):
	default:
		fmt.Println("Changes to be pushed:\n  (use \"fileconsul push [command options]\" to synchronize local files)")

		for _, remotefile := range rfDiff.Add {
			println("\tadd remote file:\t" + filepath.Join(basepath, remotefile.Path))
		}

		for _, remotefile := range rfDiff.Del {
			println("\tdelete remote file:\t" + filepath.Join(basepath, remotefile.Path))
		}

		for _, remotefile := range rfDiff.New {
			println("\tmodify remote file:\t" + filepath.Join(basepath, remotefile.Path))
		}
	}

	rfDiff = rfList.Diff(lfrfList)
	switch {
	case lfrfList.Equal(rfList):
	case rfList.Empty():
	default:
		fmt.Println("Changes to be pulled:\n  (use \"fileconsul pull [command options]\" to synchronize remote files)")

		for _, remotefile := range rfDiff.Add {
			println("\tadd local file:\t" + filepath.Join(basepath, remotefile.Path))
		}

		for _, remotefile := range rfDiff.Del {
			println("\tdelete local file:\t" + filepath.Join(basepath, remotefile.Path))
		}

		for _, remotefile := range rfDiff.New {
			println("\tmodify local file:\t" + filepath.Join(basepath, remotefile.Path))
		}
	}
}
