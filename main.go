package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, site := range sites {
		go checkSite(site, c)
	}

	for site := range c {
		go func(s string) {
			time.Sleep(3 * time.Second)
			checkSite(s, c)
		}(site)
	}
}

func checkSite(site string, c chan string) {
	_, err := http.Get(site)

	if err != nil {
		fmt.Println(site, "is down")
		c <- site
		return
	}

	fmt.Println(site, "is up")
	c <- site
}
