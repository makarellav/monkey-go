package main

import (
	"fmt"
	"github.com/makarellav/monkey-go/repl"
	"os"
)

func main() {
	fmt.Println("Hello! This is a Monkey programming language!")
	fmt.Println("Let's get typing!")

	repl.Start(os.Stdin, os.Stdout)
}
