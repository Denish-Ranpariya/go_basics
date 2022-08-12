package main

import "fmt"

func main() {
	fmt.Println("Say hello to strutctures in GO")

	//there is no inheritance in GO; No super or parent

	var denish User = User{"Denish", "thisisdenish@go.dev", true, 21}

	fmt.Println(denish)

	fmt.Printf("Denish details are : %+v\n", denish) //to get more details use %+v
	fmt.Printf("Denish email is : %v", denish.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
