package main

import (
	"ascii-art-output/internal"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [OPTIONS] [STRING] [BANNER]")
		os.Exit(1)
	}

	// Store options
	options := make(map[string]string)
	var textFromOutside, bannerType string

	// Parse command-line arguments dynamically
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--") {
			// Extract key-value pairs for options
			parts := strings.SplitN(arg[2:], "=", 2)
			if len(parts) == 2 {
				options[parts[0]] = parts[1]
			} else {
				options[parts[0]] = ""
			}
		} else if textFromOutside == "" {
			textFromOutside = arg
		} else {
			bannerType = arg
		}
	}

	// Ensure a valid input string
	if textFromOutside == "" {
		fmt.Println("Error: Missing required STRING argument.")
		os.Exit(1)
	}

	// Default banner type if not specified
	if bannerType == "" {
		bannerType = "standard"
	}

	// Validate input for non-ASCII characters
	if containsNonASCII(textFromOutside) {
		fmt.Println("Error: Input contains non-ASCII characters.")
		os.Exit(1)
	}

	// Get available banners dynamically
	bannerFilePath := filepath.Join("banners", bannerType+".txt")
		// Banner hashes map for validation
		bannerHashes := map[string]string{
			"standard":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
			"shadow":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
			"thinkertoy": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
		}
	
		// Validate the banner type
		expectedHash, exists := bannerHashes[bannerType]
		if !exists {
			fmt.Printf("Error: Banner type '%s' not recognized.\n", bannerType)
			os.Exit(1)
		}
		// Validate the integrity of the banner file
	err := internal.ValidateFileHash(bannerFilePath, expectedHash)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	if _, err := os.Stat(bannerFilePath); os.IsNotExist(err) {
		fmt.Printf("Error: Banner type '%s' not found.\n", bannerType)
		os.Exit(1)
	}

	// Read and parse the banner file
	fileLines := internal.ReadBannerFile(bannerFilePath)
	
	asciiTemplates := internal.ParseBanner(fileLines)

	// Render ASCII art
	asciiArt := internal.RenderASCIIArt(textFromOutside, asciiTemplates)
	

	// Handle output option
	if outputFileName, exists := options["output" ]; exists {
		if !strings.HasSuffix(outputFileName, ".txt") {
			fmt.Println("Error: Output file must have a .txt extension.")
			os.Exit(1)
		}

		err := os.WriteFile(outputFileName, []byte(asciiArt), 0644)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
		fmt.Printf("Generated ASCII art saved to %s\n", outputFileName)
	} else {
		fmt.Println(asciiArt)
	}
}

// containsNonASCII checks if the input string contains non-ASCII characters
func containsNonASCII(text string) bool {
	for _, r := range text {
		if r > unicode.MaxASCII {
			return true
		}
	}
	return false
}
