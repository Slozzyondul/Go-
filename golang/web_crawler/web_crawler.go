package webcrawler1

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// SafeMap is a concurrency-safe map for keeping track of fetched URLs.
type SafeMap struct {
	mu      sync.Mutex
	fetched map[string]bool
}

func (s *SafeMap) Set(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.fetched[url] = true
}

func (s *SafeMap) Exists(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.fetched[url]
}

var fetchedURLs = SafeMap{fetched: make(map[string]bool)}

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 {
		return
	}
	if fetchedURLs.Exists(url) {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fetchedURLs.Set(url)
	fmt.Printf("found: %s %q\n", url, body)
	var innerWg sync.WaitGroup
	for _, u := range urls {
		innerWg.Add(1)
		go Crawl(u, depth-1, fetcher, &innerWg)
	}
	innerWg.Wait()
}

func WebCrawler() {
	var wg sync.WaitGroup
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, &wg)
	wg.Wait()
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
}
