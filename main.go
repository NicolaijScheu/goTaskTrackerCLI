package main

import (
	"goTaskTrackerCLI/app"
	cliparser "goTaskTrackerCLI/cliParser"
	"os"
)

func main() {

	cliparser.ParseArgs(os.Args[1:])
	app.Start()
}
