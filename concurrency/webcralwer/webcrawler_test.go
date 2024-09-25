package main

import (
	"cmp"
	"errors"
	"slices"
	"testing"
)

func compareCrawlResults(a, b *CrawlResult) int {
	if x := cmp.Compare(a.Body, b.Body); x != 0 {
		return x
	}
	if x := slices.Compare(a.URLs, b.URLs); x != 0 {
		return x
	}
	if a.Err == nil && b.Err == nil {
		return 0
	}
	return cmp.Compare(a.Err.Error(), b.Err.Error())
}

func equalCrawlResults(a, b *CrawlResult) bool {
	return compareCrawlResults(a, b) == 0
}

func TestCrawlResults(t *testing.T) {
	tests := []struct {
		Name string
		URL  string
		Want []*CrawlResult
	}{
		{
			Name: "base case",
			URL:  "https://golang.org/",
			Want: []*CrawlResult{
				{Body: "Command go", URLs: []string{"https://golang.org/", "https://golang.org/pkg/"}, Err: error(nil)},
				{Body: "Package fmt", URLs: []string{"https://golang.org/", "https://golang.org/pkg/"}, Err: error(nil)},
				{Body: "Package os", URLs: []string{"https://golang.org/", "https://golang.org/pkg/"}, Err: error(nil)},
				{Body: "Packages", URLs: []string{"https://golang.org/", "https://golang.org/cmd/", "https://golang.org/pkg/fmt/", "https://golang.org/pkg/os/"}, Err: error(nil)},
				{Body: "The Go Programming Language", URLs: []string{"https://golang.org/pkg/", "https://golang.org/cmd/"}, Err: error(nil)},
			},
		},
		{
			Name: "nonexisting URL",
			URL:  "http://example.com",
			Want: []*CrawlResult{
				{Body: "", URLs: nil, Err: errors.New("not found: http://example.com")},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			results := CrawlResults(tt.URL, 4, fetcher)
			slices.SortFunc(results, compareCrawlResults)
			if !slices.EqualFunc(results, tt.Want, equalCrawlResults) {
				t.Errorf("%s - bad results, result slices not equal: got %s, want %s", tt.Name, results, tt.Want)
			}
		})
	}
}
