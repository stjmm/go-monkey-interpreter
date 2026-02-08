package main

import (
	"fmt"
	"monkey/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! You already know what language it is...\n", 
		user.Username)
	fmt.Printf("Now type in commands!\n")
	repl.Start(os.Stdin, os.Stdout)
}
