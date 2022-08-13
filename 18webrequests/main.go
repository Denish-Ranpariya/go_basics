package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://api.chucknorris.io/jokes/random"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	if response.Status == "200 OK" {
		dataBytes, err := ioutil.ReadAll(response.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(dataBytes))
	}

	defer response.Body.Close()
}
