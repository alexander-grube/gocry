package main

import (
	"os"

	"github.com/alexander-grube/gocry/cli"
)

func main() {
	defer os.Exit(0)
	cli := cli.CommandLine{}
	cli.Run()
}
