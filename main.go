package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://go.dev",
		"https://golang.org",
	}
	for _, link := range links {
		checkStatus(link)
	}
}

func checkStatus(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link + "might be down :(")
		return
	}
	fmt.Println(link, "is up :)")
}
