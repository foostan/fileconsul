package main

import (
	"github.com/codegangsta/cli"

	"github.com/foostan/fileconsul/command"
)

var Commands = []cli.Command{
	cli.Command{
		Name:        "status",
		Usage:       "Show status of local files",
		Description: "Show the difference between local files and remote files that is stored in K/V Store of a consul cluster.",
		Flags:       command.StatusFlags,
		Action:      command.StatusCommand,
	},
	cli.Command{
		Name:        "pull",
		Usage:       "Pull files from a consul cluster",
		Description: "Pull remote files from K/V Store of a consul cluster.",
		Flags:       command.PullFlags,
		Action:      command.PullCommand,
	},
	cli.Command{
		Name:        "push",
		Usage:       "Push file to a consul cluster",
		Description: "Push remote files to K/V Store of a consul cluster.",
		Flags:       command.PushFlags,
		Action:      command.PushCommand,
	},
}
