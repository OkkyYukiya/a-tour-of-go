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

type SafeCache struct {
	mu      sync.Mutex
	visited map[string]bool
}

// Visit はURLを訪問済みとして記録、すでに訪問済みかどうか判定
func (c *SafeCache) Visit(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	// すでに訪問済み
	if c.visited[url] {
		return false 
	}
	c.visited[url] = true
	// 初めての訪問
	return true 
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, cache *SafeCache, wg *sync.WaitGroup) {
	defer wg.Done()
	
	if depth <= 0 {
		return
	}
	
	// すでに訪問済みのURLはスキップ
	if !cache.Visit(url) {
		return
	}
	
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	
	fmt.Printf("found: %s %q\n", url, body)
	
	// 見つかったURLを並列でクロールする
	for _, u := range urls {
		wg.Add(1)
		// goをつけてgoroutineに
		go Crawl(u, depth-1, fetcher, cache, wg)
	}
}

func main() {
	cache := &SafeCache{visited: make(map[string]bool)}
	var wg sync.WaitGroup
	
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, cache, &wg)
	
	wg.Wait()
	fmt.Println("\n--- Crawling finished! ---")
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