package main

import (
	"os"

	"github.com/gouniverse/envenc"
)

func main() {
	args := os.Args
	envenc.NewCli().Run(args[0:])
}
