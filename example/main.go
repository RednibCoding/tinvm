package main

import (
	"fmt"
	"os"

	"github.com/RednibCoding/tinvm"
)

func main() {
	// args := []string{"example.exe", "test.tin"}
	args := os.Args
	if len(args) < 2 {
		fmt.Println("USAGE: tin <sourcefile>")
		os.Exit(1)
	}
	source, err := os.ReadFile(args[1])
	if err != nil {
		fmt.Printf("ERROR: Can't find source file '%s'.\n", args[1])
		os.Exit(1)
	}

	vm := tinvm.New()
	vm.Run(string(source), args[1])
}
