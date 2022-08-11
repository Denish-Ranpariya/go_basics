package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza shop")
	fmt.Println("Please rate our pizza between 1 to 5 :")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("ERROR : ", err)
	} else {
		fmt.Println("Rating :", rating)
	}

}
