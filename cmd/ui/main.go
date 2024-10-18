package main

import (
	"os"

	"github.com/gouniverse/envenc"
)

// New UI in development
func main() {
	args := os.Args
	envenc.NewCliV2().Run(args[0:])
}
