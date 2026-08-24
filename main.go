package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: iac-validator <config-file>")
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer f.Close()
	resources := 0
	issues := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.Contains(line, "resource") || strings.Contains(line, "Resource") {
			resources++
		}
		if strings.Contains(line, "password") && !strings.Contains(line, "var.") {
			issues++
			fmt.Printf("  [WARN] Hardcoded credential at: %s\n", line)
		}
		if strings.Contains(line, "0.0.0.0/0") {
			issues++
			fmt.Printf("  [WARN] Open CIDR (0.0.0.0/0) detected\n")
		}
	}
	fmt.Printf("\nIaC Validator\n")
	fmt.Printf("Resources: %d\n", resources)
	fmt.Printf("Issues:    %d\n", issues)
}