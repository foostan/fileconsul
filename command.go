package main

import (
	"github.com/codegangsta/cli"

	"github.com/foostan/fileconsul/command"
)

var Commands = []cli.Command{
	cli.Command{
		Name:        "status",
		Usage:       "",
		Description: "Show status of local files",
		Flags:       command.StatusFlags,
		Action:      command.StatusCommand,
	},
	cli.Command{
		Name:        "pull",
		Usage:       "",
		Description: "Pull files from a consul cluster",
		Flags:       command.PullFlags,
		Action:      command.PullCommand,
	},
	cli.Command{
		Name:        "push",
		Usage:       "",
		Description: "Push files to a consul cluster",
		Flags:       command.PushFlags,
		Action:      command.PushCommand,
	},
}