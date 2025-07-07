package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(urlString string) (string, error) {
	urlString = strings.Trim(strings.ToLower(urlString), "/")
	urlString = strings.ReplaceAll(urlString, " ", "")
	urlStruct, err := url.Parse(urlString)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}
	normalizedURL := fmt.Sprintf("%s%s", urlStruct.Host, urlStruct.EscapedPath())
	return normalizedURL, nil
}
