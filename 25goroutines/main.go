package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

func main() {
	// go greeter("Hello")
	// greeter("World")

	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://keep.google.com",
	}

	var waitGroup sync.WaitGroup

	for _, website := range websiteList {
		waitGroup.Add(1)
		var statusCode int
		go func(currentWebsite string) {
			statusCode = getStatusCode(currentWebsite, &waitGroup)
			fmt.Println("Status code from ", currentWebsite, " ", statusCode)
		}(website)
	}
	waitGroup.Wait()
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string, waitGroup *sync.WaitGroup) int {
	response, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer waitGroup.Done()
	return response.StatusCode
}
