package main

import (
	"testing"
)

func TestCrawler(t *testing.T) {
	Crawl("http://golang.org/", 4, fetcher)
}
