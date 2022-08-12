package main

import (
	"fmt"
)

func main() {
	fmt.Println("Say hello to if else in GO")

	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Something else"
	} else {
		result = "PARKOUR"
	}

	fmt.Println(result)

	//initializing before checking for conditions
	//this syntax is helpful when data is coming from web request
	if number := 3; number > 10 {
		fmt.Println("Number is greater than 10.")
	} else {
		fmt.Println("Number is not greater than 10.")
	}
}
