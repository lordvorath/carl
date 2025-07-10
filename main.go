package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"sync"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Print("not enough arguments: crawler <url> <maxConcurrency> <maxPages>\n")
		return
	}
	if len(args) > 3 {
		fmt.Print("too many arguments: : crawler <url> <maxConcurrency> <maxPages>\n")
		return
	}

	baseURL, err := url.Parse(args[0])
	if err != nil {
		log.Printf("bad base url: %v", err)
		return
	}

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		log.Printf("not a valid number: %v", err)
		return
	}
	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		log.Printf("not a valid number: %v", err)
		return
	}

	fmt.Printf("=== Starting crawl of: %s ===\n", baseURL.String())

	cfg := config{
		pages:              map[string]int{},
		maxPages:           maxPages,
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}
	u := baseURL.String()
	cfg.crawlPage(u)
	cfg.wg.Wait()

	log.Print("=== Crawl Complete ===")
	log.Printf("Total number of links: %d", len(cfg.pages))
	for k, v := range cfg.pages {
		log.Printf("%s - %d", k, v)
	}

}
