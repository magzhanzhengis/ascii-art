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

	if len(os.Args) != 2 {
		fmt.Print("Give more arguments")
		return
	}
	for i := 0; i < len(os.Args[1]); i++ {
		// if banner[rune((os.Args[1][i]))] == nil {
		// 	// fmt.Printf("The character '%c' not found: \n", os.Args[1][i])
		// 	// os.Exit(1)
		// 	fmt.Println(banner[rune((os.Args[1][i]))])
		// }
		fmt.Println(banner[rune((os.Args[0][i]))])
	}
	// render(os.Args[1], banner)

}
