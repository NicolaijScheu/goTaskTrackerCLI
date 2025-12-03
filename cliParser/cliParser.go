package cliparser

import "fmt"

func ParseArgs(args []string) string {

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
		if args[i] == "-add" {
			return "add"
		}
	}
	return "error"
}
