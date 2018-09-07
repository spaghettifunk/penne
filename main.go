package main

import (
	"os/user"
	"fmt"
	"github.com/spaghettifunk/penne/repl"
	"os"
)

func main() {
	u, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Howdy %s! Welcome to the Penne programming language!\n", u.Username)
	fmt.Printf("Start typing in some commands\n")

	repl.Start(os.Stdin, os.Stdout)
}