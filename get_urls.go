package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func getURLsFromHTML(htmlbody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %w", err)
	}
	temp := []string{}
	urls := []string{}
	doc, err := html.Parse(strings.NewReader(htmlbody))
	if err != nil {
		return nil, fmt.Errorf("failed to parse body: %w", err)
	}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if strings.ToLower(attr.Key) == "href" {
					temp = append(temp, attr.Val)
				}
			}
		}
	}
	for _, link := range temp {
		href, err := url.Parse(link)
		if err != nil {
			continue
		}
		resolvedURL := baseURL.ResolveReference(href)

		urls = append(urls, resolvedURL.String())
	}
	if len(urls) == 0 {
		return nil, nil
	}
	return urls, nil
}
