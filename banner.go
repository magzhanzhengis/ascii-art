package main

import (
	"bufio"
	"fmt"
	"os"
)

// LoadBanner loads the banner data from the file and returns a map of ASCII characters
func loadbanner() (map[rune][]string, error) {
	// Open the banner file
	bannerfile, err := os.Open("/Users/mzhengis/ascii-art/standard.txt")
	if err != nil {
		fmt.Printf("Error opening the file: %v\n", err)
		return nil, err
	}
	defer bannerfile.Close()

	// Initialize variables
	scanner := bufio.NewScanner(bannerfile)
	banner := make(map[rune][]string)
	var lines []string
	currentChar := ' ' // Start from the space character

	// Read the file line by line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if len(lines) > 0 {
				banner[currentChar] = lines
				fmt.Printf("Loaded character: '%c'\n", currentChar) // Debug print
				lines = []string{}
				currentChar++
			}
			continue
		}
		lines = append(lines, line)
	}

	// Add the last character (if any)
	if len(lines) > 0 {
		banner[currentChar] = lines
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning the file: %v\n", err)
		return nil, err
	}

	return banner, nil
}
