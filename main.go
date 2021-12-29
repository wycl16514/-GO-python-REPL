package main 

import (
	"repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hell %s! This is python REPL!\n", user.Username)
	fmt.Printf("Please type some python code...\n")
	repl.Start(os.Stdin, os.Stdout)
}