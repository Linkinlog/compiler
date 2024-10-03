package main

import (
	"os"

	"gitlab.com/linkinlog/compiler/repl"
)

func main() {
	if len(os.Args) > 1 {
		repl.RunFile(os.Args[1])
		return
	}
	repl.Start(os.Stdin, os.Stdout)
}
