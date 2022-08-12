package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in GO")

	//required to initialize length while declaring list as per language specs
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Watermelon"

	fmt.Println("Fruit list : ", fruitList)
	//[Apple Mango  Watermelon]
	//notice that extra white space!

	fmt.Println("Length of the list : ", len(fruitList))

	//initialization while declaration
	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Veg list : ", vegList)
	
}
