package main

import (
	"errors"
	"fmt"
	"net/http"
)

type requestResult struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string) // make init => empty map
	c := make(chan requestResult)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _, url := range urls {
		go hitUrl(url, c)
	}

	for i := 0; i < len(urls); i++ {
		result := <-c
		results[result.url] = result.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

var errRequestFailed = errors.New("Request failed")

// 보내기만 가능하고 받을 순 없는 채널을 설정할 수도 있다.
// func hitUrl(url string, c chan<- result) {
func hitUrl(url string, c chan requestResult) {
	// fmt.Println(<-c) // 채널로부터 받을 수도 있다.
	resp, err := http.Get(url)
	status := "OK"
	if err != nil || resp.StatusCode >= 400 {
		status = "FAILED"
	}

	c <- requestResult{url: url, status: status}
}
