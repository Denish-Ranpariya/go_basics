package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Say hello to slices!")

	var fruitList = []string{"Apple", "Tomato", "Peach"}

	fmt.Printf("Type of fruitList is : %T \n", fruitList) //[]string

	fruitList = append(fruitList, "Banana")

	fmt.Println("Fruit list : ", fruitList)

	result1 := fruitList[1:] //slicing from index 1 to end
	fmt.Println("Fruit list : ", result1)

	result2 := fruitList[1:3] //1 inclusive 3 exclusive
	fmt.Println("Fruit list : ", result2)

	result3 := fruitList[:2] //slicing from start to 1 (because 2 is exclusive)
	fmt.Println("Fruit list : ", result3)

	highScores := make([]int, 4)

	highScores[0] = 113
	highScores[1] = 2445
	highScores[2] = 53232
	highScores[3] = 443

	// highScores[4] = 113
	// above line of code will give error of array index out of bound because in the make method initial size is 4 and memory allocation is already happened

	fmt.Println("Highscores : ", highScores)

	highScores = append(highScores, 124, 1113, 3112, 323)
	// above line of code won't give the error because append will do memory allocation again

	fmt.Println("Highscores : ", highScores)

	sort.Ints(highScores)

	fmt.Println("Highscores in ascending order : ", highScores)

}
