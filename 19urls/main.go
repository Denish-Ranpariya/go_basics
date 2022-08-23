package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://github.com:3000/learn?coursename=reactjs&paymentid=kfjdfdkkbk"

func main() {
	fmt.Println("Hello")
	fmt.Println(myUrl)

	//parsing
	result, error := url.Parse(myUrl)

	if error != nil {
		panic(error)
	}

	fmt.Println(result.Scheme)   //https
	fmt.Println(result.Host)     // https
	fmt.Println(result.Path)     //github.com:3000
	fmt.Println(result.Port())   // /learn
	fmt.Println(result.RawQuery) //coursename=reactjs&paymentid=kfjdfdkkbk

	queryParams := result.Query()
	fmt.Println(queryParams) // map

	for _, value := range queryParams {
		fmt.Println("Queryparam: ", value)
	}

	//construct URL
	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "github.com",
		Path:     "/user",
		RawQuery: "userid=ngflgsfgn",
	}

	anotherURL := partsOfURL.String()
	fmt.Println(anotherURL)

}
