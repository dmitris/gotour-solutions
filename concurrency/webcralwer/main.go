package main

import (
	"fmt"
	"strings"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// CrawlResult wraps the results of crawling a URL including
// the potential error.
type CrawlResult struct {
	Body string
	URLs []string
	Err  error
}

func (c *CrawlResult) String() string {
	return fmt.Sprintf("Body: %s, URLs: [%s], Err: %v", c.Body, strings.Join(c.URLs, ", "), c.Err)
}

// MapWithMutex combines a map[string]bool and a mutex to synchronise access.
type MapWithMutex struct {
	M  map[string]bool
	mu *sync.Mutex
}

func NewMapWithMutex() *MapWithMutex {
	return &MapWithMutex{
		M:  map[string]bool{},
		mu: new(sync.Mutex),
	}
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	seen := NewMapWithMutex() // to keep track of the already seen URLs
	ch := make(chan *CrawlResult, 8)
	go crawlStart(url, depth, fetcher, ch, seen)
	for res := range ch {
		fmt.Printf("found: %s\n", res)
	}
}

// CrawlResults uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func CrawlResults(url string, depth int, fetcher Fetcher) []*CrawlResult {
	seen := NewMapWithMutex() // to keep track of the already seen URLs
	ch := make(chan *CrawlResult, 8)
	ret := []*CrawlResult{}
	go crawlStart(url, depth, fetcher, ch, seen)
	for res := range ch {
		// fmt.Printf("found: %s\n", res)
		ret = append(ret, res)
	}
	return ret
}

func crawlStart(url string, depth int, fetcher Fetcher, ch chan *CrawlResult, seen *MapWithMutex) {
	defer close(ch)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go crawl(url, depth, fetcher, ch, wg, seen)
	wg.Wait()
}

func crawl(url string, depth int, fetcher Fetcher, ch chan<- *CrawlResult, wg *sync.WaitGroup, seen *MapWithMutex) {
	// Fetch URLs in parallel.
	// Don't fetch the same URL twice.
	defer wg.Done()
	if depth <= 0 {
		return
	}
	isSeen := func() bool {
		seen.mu.Lock()
		defer seen.mu.Unlock()
		if _, ok := seen.M[url]; ok {
			return true
		}
		seen.M[url] = true
		return false
	}()
	if isSeen {
		return
	}
	body, urls, err := fetcher.Fetch(url)

	// fmt.Printf("found: %s %q\n", url, body)
	ch <- &CrawlResult{
		Body: body,
		URLs: urls,
		Err:  err,
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, u := range urls {
		wg.Add(1)
		go crawl(u, depth-1, fetcher, ch, wg, seen)
	}
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/cmd/": &fakeResult{
		"Command go",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
