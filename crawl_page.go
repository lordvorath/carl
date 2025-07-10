package main

import (
	"log"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]int
	maxPages           int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {

	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != cfg.baseURL.Hostname() {
		return
	}

	normURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("failed to normalize url: %v", err)
		return
	}

	if isFirst := cfg.addPageVisit(normURL); !isFirst {
		return
	}

	log.Printf("crawling -> %s", normURL)
	bod, err := getHTML(currentURL.String())
	if err != nil {
		log.Printf("failed to get HTML: %v", err)
		return
	}

	urls, err := getURLsFromHTML(bod, currentURL.String())
	if err != nil {
		log.Printf("failed to get links: %v", err)
		return
	}

	for _, u := range urls {
		cfg.wg.Add(1)
		go cfg.crawlPage(u)
	}

}

func (cfg *config) addPageVisit(normURL string) bool {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, ok := cfg.pages[normURL]; ok {
		cfg.pages[normURL]++
		return false
	}

	cfg.pages[normURL] = 1
	return true
}
