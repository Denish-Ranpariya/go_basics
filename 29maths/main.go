package main

import (

	// "math/rand"

	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {

	// rand.Seed(time.Now().UnixNano())
	// randomNumber := rand.Intn(5)

	// fmt.Println("Random number :", randomNumber)

	//random number from "crypto/rand"
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(5))

	if err != nil {
		panic(err)
	}

	fmt.Println("Random number :", randomNumber)
}
