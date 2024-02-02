package main

import (
	"fmt"
	"github.com/makarellav/monkey-go/lexer"
)

func main() {
	l := lexer.New(`=+(){},;`)

	fmt.Printf("%+v\n", l)
}
