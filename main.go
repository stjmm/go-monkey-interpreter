package main

import (
	"fmt"
	"os"
	"os/user"
	"monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Welcome %s! This is the Monkey programming lang!\n", user.Username)
	fmt.Printf("REPL Started\n")
	repl.Start(os.Stdin, os.Stdout)
}
