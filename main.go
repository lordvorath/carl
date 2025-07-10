package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"slices"
	"strconv"
	"strings"
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
	printReport(cfg.pages, cfg.baseURL.String())

}

func printReport(pages map[string]int, baseURL string) {
	fmt.Print("=============================\n")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Print("=============================\n")
	visits := []visit{}
	for u, n := range pages {
		visits = append(visits, visit{n, u})
	}
	slices.SortFunc(visits, func(a, b visit) int {
		diff := b.numberOfVisits - a.numberOfVisits
		if diff != 0 {
			return diff
		}
		return strings.Compare(a.URL, b.URL)
	})
	for _, v := range visits {
		fmt.Printf("Found %d internal links to %s\n", v.numberOfVisits, v.URL)
	}
}

type visit struct {
	numberOfVisits int
	URL            string
}
