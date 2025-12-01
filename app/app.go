package app

import "fmt"

var ErrShutdown = fmt.Errorf("application was shutdown gracefully")

func Start() error {
	fmt.Println("Hi, Mum <3")
	return ErrShutdown
}

func Shutdown() {
	// Shutdown contexts, listeners, and such
}
