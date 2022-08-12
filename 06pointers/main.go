package main

import "fmt"

func main() {
	fmt.Println("Say hello to pointers!")

	// var ptr *int

	// fmt.Println("Value of var ptr : ", ptr) // <nil>
	// fmt.Printf("Type of var ptr : %T", ptr) // *int

	myNumber := 23 //suppose myNumber is stored at memory address X

	var ptr = &myNumber //& is used to refer a variable (and reference of a variable is always a memory address!)

	fmt.Println("Pointer : ", ptr)
	// ptr refers to memory address, so it will be X

	fmt.Println("Value of pointer : ", *ptr)
	//*ptr will give value stored at memory address ptr, so it will be 23 (stored at memory address X)

	*ptr = *ptr + 2
	//this operation will take place on value stored at memory address X

	fmt.Println("New value is : ", myNumber)
	//as myNumber is also refering to the memory address X the new value of myNumber will be 23 + 2 = 25

	//that's the actual use of the pointers, the operations will be performed on actual variables not on the copy of varibles
}
