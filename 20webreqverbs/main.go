package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Hello!")
	// PerfomGetRequest("http://localhost:8000/get")
	// PerformPostJsonRequest("http://localhost:8000/post")
	PerformPostFormRequest("http://localhost:8000/postform")
}

func PerfomGetRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code : " + response.Status)

	fmt.Println("Content length : ", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Println("Content : ", string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Content : ", responseString.String())
	fmt.Println("Byte count : ", byteCount)

}

func PerformPostJsonRequest(url string) {
	//fale json payload
	requestBody := strings.NewReader(`{
		"coursename" : "let's go with go lang",
		"price" : 0
	}
	`)
	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code : " + response.Status)

	fmt.Println("Content length : ", response.ContentLength)

	content, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	// fmt.Println("Content : ", string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Content : ", responseString.String())
	fmt.Println("Byte count : ", byteCount)
}

func PerformPostFormRequest(targetUrl string) {
	data := url.Values{}

	data.Add("firstname", "denish")
	data.Add("lastname", "ranpariya")
	data.Add("email", "thisisdenish@gmail.com")

	response, err := http.PostForm(targetUrl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	// fmt.Println("Content : ", string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("Content : ", responseString.String())
	fmt.Println("Byte count : ", byteCount)
}
