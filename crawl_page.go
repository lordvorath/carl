package main

import (
	"log"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	currentURL, err := url.Parse(rawCurrentURL)
	if err != nil {
		log.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawCurrentURL, err)
		return
	}
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		log.Printf("Error - crawlPage: couldn't parse URL '%s': %v\n", rawBaseURL, err)
		return
	}

	// skip other websites
	if currentURL.Hostname() != baseURL.Hostname() {
		return
	}

	normURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("failed to normalize url: %v", err)
		return
	}

	if _, ok := pages[normURL]; ok {
		pages[normURL]++
		return
	}

	pages[normURL] = 1
	log.Printf("crawling -> %s", normURL)
	bod, err := getHTML(rawCurrentURL)
	if err != nil {
		log.Printf("failed to get HTML: %v", err)
		return
	}

	urls, err := getURLsFromHTML(bod, rawBaseURL)
	if err != nil {
		log.Printf("failed to get links: %v", err)
		return
	}

	for _, u := range urls {
		crawlPage(rawBaseURL, u, pages)
	}
}
