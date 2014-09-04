package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "fileconsul"
	app.Version = Version
	app.Usage = ""
	app.Author = "foostan"
	app.Email = "ks@fstn.jp"
	app.Commands = Commands

	app.Run(os.Args)
}
