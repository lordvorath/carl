package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Print("no website provided\n")
		return
	}
	if len(args) > 1 {
		fmt.Print("too many arguments provided\n")
		return
	}
	baseURL := args[0]
	fmt.Printf("=== Starting crawl of: %s ===\n", baseURL)

	pages := make(map[string]int)
	crawlPage(baseURL, baseURL, pages)

	log.Print("=== Crawl Complete ===")
	log.Printf("Total number of links: %d", len(pages))
	for k, v := range pages {
		log.Printf("%s - %d", k, v)
	}

}
