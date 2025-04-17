package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"
)

type Crawler struct {
	urls []string
}

func (c *Crawler) Init() *Crawler {
	c.urls = []string{
		"https://google.com/a",
		"https://google.com/b",
		"https://google.com/b#notnew",
		"https://google.com/c",
		"https://yahoo.com/b",
		"https://yahoo.com/c",
		"https://google1.com/a",
		"http://google.com/a#same",
		"mail://google@googl.com",
		"https://google.com/new",
		"http://google.com/anothernew",
		"https://google.com/d",
		"https://google.com/e",
		"https://google.com/f",
	}
	return c
}

func (c *Crawler) Crawl() []string {
	if len(c.urls) == 0 {
		return nil
	}

	count := rand.Intn(len(c.urls)) + 1

	shuffled := append([]string{}, c.urls...)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	return shuffled[:count]
}

func main() {
	// ✅ Fixed the invalid URL (you had "https//:")
	result := crawlAndGetUniqueUrls("https://google.com")
	fmt.Println("\nFinal Unique URLs:")
	for _, r := range result {
		fmt.Println(r)
	}
}

func crawlAndGetUniqueUrls(seed string) []string {
	const workers = 3
	urlChan := make(chan string)
	resultsChan := make(chan []string)
	var wg sync.WaitGroup
	visited := make(map[string]bool)
	var mu sync.Mutex

	// Use a sync.Once to ensure channels are closed only once.
	var once sync.Once

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		go handle(i, urlChan, resultsChan, &wg)
	}

	// Submit goroutine
	go func() {
		for urls := range resultsChan {
			for _, curr := range urls {
				mu.Lock()
				if !visited[curr] {
					visited[curr] = true
					wg.Add(1)
					urlChan <- curr // Add new URLs to the channel for crawling
				}
				mu.Unlock()
			}
		}

		// Close channels after all processing is done
		once.Do(func() {
			close(urlChan)
			close(resultsChan)
		})
	}()

	// Seed URL
	mu.Lock()
	visited[seed] = true
	mu.Unlock()
	wg.Add(1)
	urlChan <- seed // Start with the seed URL

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All tasks done!")

	var collection []string
	for k := range visited {
		collection = append(collection, k)
	}
	return collection
}

func handle(id int, url <-chan string, results chan<- []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for u := range url {
		next := crawl(u)
		var collect []string
		for _, n := range next {
			if isSameDomain(n, u) {
				collect = append(collect, n)
			}
		}
		results <- collect
	}
}

func isSameDomain(i string, j string) bool {
	domainI := getDomain(i)
	domainJ := getDomain(j)
	return domainI != "" && domainJ != "" && domainI == domainJ
}

func getDomain(s string) string {
	if strings.Contains(s, "http") {
		parts := strings.Split(s, "/")
		if len(parts) >= 3 {
			return parts[2]
		}
	}
	return ""
}

func crawl(s string) []string {
	c := Crawler{}
	c.Init()
	return c.Crawl()
}
