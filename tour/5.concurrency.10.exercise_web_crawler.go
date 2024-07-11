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

type SafeMap struct {
	mutex      sync.Mutex
	string_map map[string]string
}

func (safe_map *SafeMap) SetUnsetKey(key string, val string) {
	safe_map.mutex.Lock()
	_, ok := safe_map.string_map[key]
	if !ok {
		safe_map.string_map[key] = val
	}
	safe_map.mutex.Unlock()
}

func (safe_map *SafeMap) KeyExists(key string) bool {
	safe_map.mutex.Lock()
	defer safe_map.mutex.Unlock()
	_, ok := safe_map.string_map[key]
	return ok

}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) map[string]string {
	var safe_map = SafeMap{string_map: make(map[string]string)}
	var wait_group sync.WaitGroup
	var crawler func(url string, depth int, fetcher Fetcher)

	crawler = func(url string, depth int, fetcher Fetcher) {
		defer wait_group.Done()

		if depth <= 0 {
			return
		}
		if safe_map.KeyExists(url) {
			return
		}

		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		safe_map.SetUnsetKey(url, body)
		for _, u := range urls {
			wait_group.Add(1)
			go crawler(u, depth-1, fetcher)
		}
		return
	}

	wait_group.Add(1)
	go crawler(url, depth, fetcher)
	wait_group.Wait()

	return safe_map.string_map
}

func main() {
	ret := Crawl("https://golang.org/", 4, fetcher)

	for url, body := range ret {
		fmt.Printf("found: %s %q\n", url, body)
	}
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
