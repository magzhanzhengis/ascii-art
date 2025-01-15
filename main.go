package main

import (
	"fmt"
	"os"
)

func main() {
	// Load the banner data
	banner, err := loadbanner()
	if err != nil {
		os.Exit(1)
	}

	// Debug: Print the banner data
	for char, ascii := range banner {
		fmt.Printf("Character: %q\n", char)
		for _, line := range ascii {
			fmt.Println(line)
		}
		fmt.Println()
	}
}
