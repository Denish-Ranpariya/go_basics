package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Say hello to switch case in GO")
	rand.Seed(time.Now().UnixNano())

	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice is : ", diceNumber)

	//in case 3 and case 4 fallthrough is present
	//so in case 3 it will go inside case 3,4 and 5
	//in case 4 it will go inside case 4,5
	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("Move ahead 2 steps")
	case 3:
		fmt.Println("Move ahead 3 steps")
		fallthrough
	case 4:
		fmt.Println("Move ahead 4 steps")
		fallthrough
	case 5:
		fmt.Println("Move ahead 5 steps")
	case 6:
		fmt.Println("Move ahead 6 steps and roll the dice again")
	default:
		fmt.Println("What was that!")
	}

}
