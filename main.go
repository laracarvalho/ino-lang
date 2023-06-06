package main

import (
	"fmt"
	"os"

	"github.com/laracarvalho/ino-lang/repl"
)

func main() {
	fmt.Printf("Hello! This is the INO programming language!\n")
	fmt.Printf("Feel free to type in commands:\n")

	repl.Start(os.Stdin, os.Stdout)
}
