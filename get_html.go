package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	res, err := http.Get(rawURL)
	if err != nil {
		return "", fmt.Errorf("getHTML failed: %w", err)
	}

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("request failed with code %d", res.StatusCode)
	}

	ct := res.Header.Get("content-type")
	if !strings.Contains(ct, "text/html") {
		return "", fmt.Errorf("content-type isn't text/html: %s", ct)
	}

	cont, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't ready response body: %w", err)
	}

	return string(cont), nil
}
