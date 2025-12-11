package main

import (
	"fmt"
	"goTaskTrackerCLI/app"
	cliparser "goTaskTrackerCLI/cliParser"
	"os"
	"strconv"
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
		convertedString, err := strconv.Atoi(command.Parameters[0])
		if err != nil {
			fmt.Println("Can't convert this to an int!")
		}
		app.AddTask(convertedString, command.Parameters[1], "todo", time.Now().String())
	}

	if command.Mode == "-update" {
		convertedString, err := strconv.Atoi(command.Parameters[0])
		if err != nil {
			fmt.Println("Can't convert this to an int!")
		}

		app.UpdateTask(convertedString, command.Parameters[1], "todo", time.Now().String())
	}

	if command.Mode == "-delete" {
		convertedString, err := strconv.Atoi(command.Parameters[0])
		if err != nil {
			fmt.Println("Can't convert this to an int!")
		}

		app.DeleteTask(convertedString)
	}

}
