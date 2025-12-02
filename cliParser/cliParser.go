package cliparser

import "fmt"

func ParseArgs(args []string) {

	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}
}
