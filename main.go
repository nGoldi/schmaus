package main

import (
	"fmt"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Schmaus programming language!\n",
		user.Username)
	fmt.Printf("Feel free to type in commands\n")
}
