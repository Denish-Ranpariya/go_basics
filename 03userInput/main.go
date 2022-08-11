package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("abcd....")

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Name ? :")

	//comma ok || comma err syntax

	input, _ := reader.ReadString('\n')

	fmt.Println("Name : " + input)
	fmt.Printf("Type of varible is: %T", input)
}
