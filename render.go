package main

import "fmt"

func render(input string, banner map[rune][]string) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < 8; j++ {
			fmt.Println(banner[rune(input[i])][j])
		}
	}
}
