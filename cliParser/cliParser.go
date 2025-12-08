package cliparser

import (
	"errors"
	"fmt"
)

func ParseArgs(args []string) (Command, error) {

	//switch statement looking at args[0]
	switch args[0] {
	case "-add":
		//check args length
		if len(args) > 1 {
			return Command{Mode: args[0], Parameters: args[1:]}, nil
		}
		fmt.Println()
		return Command{Mode: "", Parameters: []string{}}, errors.New("no values to add a task")

	// case "-update":
	// 	return "update"
	// case "-delete":
	// 	return "delete"
	// case "-mark":
	// 	return "mark"
	// case "-list":
	// 	return "list"
	default:
		return Command{Mode: "", Parameters: []string{}}, errors.New("no command found. try \"-help\"")
	}
}

type Command struct {
	Mode       string
	Parameters []string
}
