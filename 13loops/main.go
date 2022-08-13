package main

import "fmt"

func main() {
	fmt.Println("Say hello to loops in GO")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("For index %v, value is %v\n", index, day)
	// }

	randomValue := 1

	for randomValue < 10 {
		if randomValue == 8 {
			goto here
		}

		if randomValue == 5 {
			randomValue++ //need to increment the value before going for next iteration otherwise program will stuck in infinite loop
			continue
		}
		fmt.Println(randomValue)
		randomValue++
	}

here:
	fmt.Println("I am Here")
}
