package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Print("no website provided\n")
		os.Exit(1)
	}
	if len(args) > 1 {
		fmt.Print("too many arguments provided\n")
		os.Exit(1)
	}
	baseURL := args[0]
	fmt.Printf("starting crawl of: %s\n", baseURL)
	fmt.Print(getHTML(baseURL))

}
