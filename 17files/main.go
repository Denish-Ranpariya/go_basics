package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello to files in GO")

	content := "This needs to go in a file. -Unknown,2022"

	file, err := os.Create("./demo.txt")

	checkNilError(err)

	length, err := file.WriteString(content)

	checkNilError(err)

	fmt.Println("Length : ", length)
	defer file.Close()

	result := readFile("./demo.txt")
	fmt.Println(result)
}

func readFile(fileName string) string {
	dataBytes, err := ioutil.ReadFile(fileName)
	checkNilError(err)
	return string(dataBytes)
}

func checkNilError(err error) {
	if err != nil {
		panic(err) //panic will shutdown the execution of program
	}
}
