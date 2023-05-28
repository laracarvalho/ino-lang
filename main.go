package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello %s! This is the INO programming language!\n")
	fmt.Printf("Feel free to type in commands:\n")
	repl.Start(os.Stdin, os.Stdout)
}