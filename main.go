package main

import (
	"os"

	"github.com/Waramoto/hryvnia-svc/internal/cli"
)

func main() {
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}
