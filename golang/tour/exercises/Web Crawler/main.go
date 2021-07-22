package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type CachedUrl struct {
	body    string
	visited bool
}

type Counter struct {
	v   map[string]*CachedUrl
	mux sync.Mutex
	wg  sync.WaitGroup
}

func (c *Counter) addUrl(url, body string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	_, ok := c.v[url]
	if ok {
		return
	}
	c.v[url] = &CachedUrl{}
	c.v[url].body = body
	c.v[url].visited = true
}

func (c *Counter) checkVisited(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	_, ok := c.v[url]
	if ok {
		return c.v[url].visited
	}
	return ok
}

var counter *Counter = &Counter{v: make(map[string]*CachedUrl)}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer counter.wg.Done()
	if depth <= 0 {
		return
	}
	if counter.checkVisited(url) {
		fmt.Printf("found in cache, skipping: %s %q\n", url, counter.v[url].body)
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	counter.addUrl(url, body)
	for _, u := range urls {
		counter.wg.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	counter.wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher)
	counter.wg.Wait()
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
