package main

import (
	"sync"
	"testing"
)

func TestCrawler(t *testing.T) {
	var fetcher = fakeFetcher{
		"http://golang.org/": &fakeResult{
			"The Go Programming Language",
			[]string{
				"http://golang.org/pkg/",
				"http://golang.org/cmd/",
			},
		},
		"http://golang.org/pkg/": &fakeResult{
			"Packages",
			[]string{
				"http://golang.org/",
				"http://golang.org/cmd/",
				"http://golang.org/pkg/fmt/",
				"http://golang.org/pkg/os/",
			},
		},
		"http://golang.org/pkg/fmt/": &fakeResult{
			"Package fmt",
			[]string{
				"http://golang.org/",
				"http://golang.org/pkg/",
			},
		},
		"http://golang.org/pkg/os/": &fakeResult{
			"Package os",
			[]string{
				"http://golang.org/",
				"http://golang.org/pkg/",
			},
		},
	}

	var wg sync.WaitGroup
	wg.Add(1)
	Crawl("http://golang.org/", 4, fetcher, make(map[string]bool), &wg)
	wg.Wait()
}
