package main

import (
	"fmt"

	urlchecker "github.com/olcw78/yt/learn-go-nomadcoders/url-checker"
)

func main() {
	var results = make(map[string]string)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://www.soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://www.academy.nomadcoders.co/",
	}

	for _, url := range urls {
		res := "OK"
		if err := urlchecker.HitURL(url); err != nil {
			res = "NO"
		}
		results[url] = res
	}

	for url, result := range results {
		fmt.Println(url, result)
	}
}
