package main

import (
	"fmt"
	"goTaskTrackerCLI/app"
	cliparser "goTaskTrackerCLI/cliParser"
	"os"
	"time"
)

func main() {

	command, err := cliparser.ParseArgs(os.Args[1:])

	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}

	app.Start()

	if command.Mode == "-add" {
		app.AddTask(command.Parameters[0], command.Parameters[1], "todo", time.Now().String())
	}

	// if command.Mode == "update" {
	// 	app.UpdatedTask("1", "Test Task", "todo", time.Now().String(), time.Now().String())
	// }

	// if command.Mode == "delete" {
	// 	app.DeleteTask("1")
	// }

}
