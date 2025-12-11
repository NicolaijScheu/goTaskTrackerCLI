package cliparser

import (
	"errors"
)

func ParseArgs(args []string) (Command, error) {

	//switch statement looking at args[0]
	switch args[0] {
	case "-add":
		//check args length
		if len(args) > 1 {
			return Command{Mode: args[0], Parameters: args[1:]}, nil
		}
		return Command{Mode: "", Parameters: []string{}}, errors.New("no values to add a task")
	case "-update":
		//check args length
		if len(args) > 1 {
			return Command{Mode: args[0], Parameters: args[1:]}, nil
		}
		return Command{Mode: "", Parameters: []string{}}, errors.New("no values to update the task")
	case "-delete":
		return Command{Mode: args[0], Parameters: args[1:]}, nil
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
