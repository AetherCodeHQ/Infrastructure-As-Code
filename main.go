package main

import (
	"fmt"
	"os"
)

// infrastructure_as_code - Declarative infrastructure config
func infrastructure_as_code(path string) {
	fmt.Println("========================================")
	fmt.Println("  Infrastructure-As-Code")
	fmt.Println("  Declarative infrastructure config")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	infrastructure_as_code(path)
}
