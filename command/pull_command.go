package command

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/codegangsta/cli"

	. "github.com/foostan/fileconsul/fileconsul"
)

var PullFlags = []cli.Flag{
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

func PullCommand(c *cli.Context) {
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
	rfDiff := rfList.Diff(lfrfList)

	switch {
	case rfList.Empty():
		fmt.Println("There are no remote files. Skip synchronizing.")
	case lfrfList.Equal(rfList):
		fmt.Println("Already up-to-date.")
	default:
		fmt.Println("Synchronize remote files:")

		for _, remotefile := range rfDiff.Add {
			localfile := remotefile.ToLocalfile(basepath)
			err := localfile.Save()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("\tadd local file:\t" + filepath.Join(basepath, remotefile.Path))
		}

		for _, remotefile := range rfDiff.New {
			localfile := remotefile.ToLocalfile(basepath)
			err := localfile.Save()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("\tmodify local file:\t" + filepath.Join(basepath, remotefile.Path))
		}

		for _, remotefile := range rfDiff.Del {
			localfile := remotefile.ToLocalfile(basepath)
			err := localfile.Remove()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("\tdelete local file:\t" + filepath.Join(basepath, remotefile.Path))
		}

		fmt.Println("Already up-to-date.")
	}
}
