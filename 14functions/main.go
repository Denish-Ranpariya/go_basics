package main

import "fmt"

func main() {
	fmt.Println("Functions in GO")
	greetings()
	result := add(3, 2)
	fmt.Println("Result : ", result)
	result = addMultipleValues(1, 2, 3, 4, 5)
	fmt.Println("Result : ", result)
	sum, product := addAndMultiplyMultipleValues(1, 2, 3, 4)
	fmt.Printf("Sum is %v, Product is %v\n", sum, product)
}

func greetings() {
	fmt.Println("Welcome to the world of Golang")
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func addMultipleValues(values ...int) int {
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

func addAndMultiplyMultipleValues(values ...int) (int, int) {
	total := 0
	product := 1
	for _, value := range values {
		total += value
		product *= value
	}
	return total, product
}
