package main

import "fmt"

func main() {
	defer fmt.Println("World")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")
	myDefer()
	//Hello 4 3 2 1 0 two one World
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

//in normal flow of program lines getting executed in sequence
//but when defer is used it wont execute that particular line of code but go to next line and that lines of code will get executed in LIFO(last in first out) manner at the end of the program
//can use stack for visualization
