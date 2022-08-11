package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of GO")

	presentTime := time.Now()

	//have to use below date/time/day to format the time :(
	//Basic full date  "Mon Jan 2 15:04:05 MST 2006" gives "Wed Feb 25 11:06:39 PST 2015"
	//visit https://pkg.go.dev/time#Time.Format for more info
	fmt.Println(presentTime.Format("02-01-2006 Monday 15:04:05"))

	createdDate := time.Date(2000, time.December, 15, 0, 0, 0, 0, time.UTC)
	fmt.Println(createdDate.Format("02-01-2006 Monday 15:04:05 MST"))
}
