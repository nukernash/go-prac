package main

import (
	"fmt"
	crawler "go-prac/src/concurrency/threadpool/web_crawler/crawler"
	"strings"
	"sync"
)

func main() {
	result := crawlAndGetUniqueUrls("https://google.com")
	fmt.Println("\nFinal Unique URLs:")
	for _, r := range result {
		fmt.Println(r)
	}
}

func crawlAndGetUniqueUrls(seed string) []string {
	const workers = 3
	urlChan := make(chan string, 10)
	resultsChan := make(chan []string, 10)
	done := make(chan string)

	var wg sync.WaitGroup
	visited := make(map[string]bool)

	// Seed URL
	visited[seed] = true
	wg.Add(1)
	urlChan <- seed // Start with the seed URL

	// Start worker goroutines
	for i := 0; i < workers; i++ {
		go handle(i, urlChan, resultsChan, &wg)
	}

	// Submit goroutine
	go func() {
		for {
			select {
			case urls := <-resultsChan:
				fmt.Printf("processing urls : %+v", urls)
				for _, curr := range urls {
					if _, exist := visited[curr]; !exist {
						visited[curr] = true
						wg.Add(1)
						urlChan <- curr // Add new URLs to the channel for crawling
					}
				}
			case <-done:
				fmt.Println("Done. CLosing channels")
				// Close channels after all processing is done
				close(urlChan)
				close(resultsChan)
				return
			default:
				fmt.Println("Next")
			}
		}
	}()

	wg.Wait() // Wait for all workers to finish
	fmt.Println("All tasks done!")
	done <- "done"

	var collection []string
	for k := range visited {
		collection = append(collection, k)
	}
	return collection
}

func handle(id int, url <-chan string, results chan<- []string, wg *sync.WaitGroup) {
	for u := range url {
		fmt.Printf("Crawling %s \n", u)
		next := crawl(u)
		var collect []string
		for _, n := range next {
			if isSameDomain(n, u) {
				collect = append(collect, n)
			}
		}
		results <- collect
		fmt.Printf("Crawled %s : %+v\n", u, collect)
		wg.Done()
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
	c := crawler.Crawler{}
	c.Init()
	return c.Crawl()
}
