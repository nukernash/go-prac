package Crawler

import "math/rand"

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
