package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{
			"ReactJS Bootcamp",
			299,
			"lco.in",
			"abc123",
			[]string{"web-dev", "js"},
		},
		{
			"MERN Bootcamp",
			299,
			"lco.in",
			"qwerty123",
			nil,
		},
		{
			"Angular Bootcamp",
			299,
			"lco.in",
			"bcd123",
			[]string{"web-dev", "js", "google"},
		},
	}

	//packing above data into json
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", finalJson)

}

func DecodeJson() {

	//data from web request will always be in []byte
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"price": 299,
		"website": "lco.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		err := json.Unmarshal(jsonDataFromWeb, &lcoCourse)

		if err != nil {
			panic(err)
		}

		fmt.Println(lcoCourse)
		fmt.Printf("Type : %T", lcoCourse)

	} else {
		fmt.Println("Invalid Json format!")
	}

	//some cases where you just want to add data to key value pairs

	var myOnlineData map[string]interface{} // it is guaranteed that key is always a string but value can be any data type

	err := json.Unmarshal(jsonDataFromWeb, &myOnlineData)

	if err != nil {
		panic(err)
	}

	fmt.Println("Map : ", myOnlineData)

}
