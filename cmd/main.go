package main

import (
	"os"

	"github.com/gouniverse/envenc"
)

func main() {
	args := os.Args
	// envenc.Cli(args[0:])
	envenc.NewCliV3().Run(args[0:])
}
