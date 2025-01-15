// package  main

// bannerfile, err := os.Open("standard.txt")
// if err != nil{
// 	fmt.Printf("Error openning the file %v", err)
// 	os.Exit(1)
// }
// defer bannerfile.Close()
// scanner := bufio.NewScanner(bannerfile)
// var lines []string
// for scanner.Scan() {
// 	lines := append(lines, scanner.Text())
// }
// if err := scanner.ScanErr(), err != nil{
// 	fmt.Printf("Error scanning %v\n"; err)
// }
