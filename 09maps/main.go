package main

import (
	"fmt"
)

func main() {
	fmt.Println("Say hello to maps in GO")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Languages : ", languages)

	fmt.Println("JS stands for : ", languages["JS"])

	delete(languages, "RB")

	fmt.Println("Languages : ", languages)

	//iterating through map in GO

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}

}
