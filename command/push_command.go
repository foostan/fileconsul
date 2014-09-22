package command

import (
	"log"
	"fmt"
	"path/filepath"

	"github.com/codegangsta/cli"

	. "github.com/foostan/fileconsul/fileconsul"
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
		Name:  "basepath",
		Value: ".",
		Usage: "base directory path of target files",
	},
}

func PushCommand(c *cli.Context) {
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

	for _, remotefile := range rfDiff.Add {
		fmt.Println("add remote file:\t" + filepath.Join(basepath, remotefile.Path))
		err = client.PutKV(filepath.Join(prefix, remotefile.Path), remotefile.Data)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, remotefile := range rfDiff.New {
		fmt.Println("modify remote file:\t" + filepath.Join(basepath, remotefile.Path))
		err = client.PutKV(filepath.Join(prefix, remotefile.Path), remotefile.Data)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, remotefile := range rfDiff.Del {
		fmt.Println("delete remote file:\t" + filepath.Join(basepath, remotefile.Path))
		err = client.DeleteKV(filepath.Join(prefix, remotefile.Path))
		if err != nil {
			log.Fatal(err)
		}
	}
}
