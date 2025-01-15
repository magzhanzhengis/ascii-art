package main

import (
	"bufio"
	"fmt"
	"os"
)

func loadbanner() [][]string {
	bannerfile, err := os.Open("standard.txt")
	if err != nil {
		fmt.Printf("Error openning the file %v", err)
		os.Exit(1)
	}
	defer bannerfile.Close()
	scanner := bufio.NewScanner(bannerfile)
	var lines []string
	groupedlines := [][]string{}
	count := 0
	for i := 0; scanner.Scan(); i++ {
		if count%8 != 0 && scanner.Text() != "\n" {
			lines = append(lines, scanner.Text())
			count++
		}
		groupedlines = append(groupedlines, lines)
		lines = nil
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning %v\n", err)
	}
	// m := make(map[rune][]string)
	//
	//	for i = 32, i<=127, i++{
	//		m[32] = lines[i-32]
	//	}
	return groupedlines
}

// func removeByValue(slice []string, value string) []string {
// 	for i, v := range slice {
// 		if v == value {
// 			return append(slice[:i], slice[i+1:]...)
// 		}
// 	}
// 	return slice // Return the original slice if value not found
// }
