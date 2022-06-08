package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, site := range sites {
		checkSite(site)
	}
}

func checkSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println(site, " is down")
		return
	}

	fmt.Println(resp.Status)
}
