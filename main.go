package main

import (
	"goTaskTrackerCLI/app"
	cliparser "goTaskTrackerCLI/cliParser"
	"os"
	"time"
)

func main() {

	mode := cliparser.ParseArgs(os.Args[1:])
	app.Start()

	if mode == "add" {
		app.AddTask("1", "Test Task", "todo", time.Now().String())
	}
}
