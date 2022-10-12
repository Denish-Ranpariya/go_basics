package main

import (
	"fmt"
	"sync"
)

func main() {
	myChan := make(chan int, 5)

	wg := &sync.WaitGroup{}

	// fmt.Println(<-myChan)
	// myChan <- 5
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()

		value, isChannelOpen := <-myChan
		fmt.Println(value, " ", isChannelOpen)
		newValue, open := <-myChan
		fmt.Println(newValue, " ", open)
	}(myChan, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		myChan <- 5
		close(myChan)
	}(myChan, wg)

	wg.Wait()
}
