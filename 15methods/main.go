package main

import "fmt"

func main() {
	//there is no inheritance in GO; No super or parent

	var denish User = User{"Denish", "thisisdenish@go.dev", true, 21}

	fmt.Println(denish)

	fmt.Printf("Denish details are : %+v\n", denish) //to get more details use %+v
	fmt.Printf("Denish email is : %v\n", denish.Email)
	fmt.Println("Status : ", denish.GetStatus())
	denish.SetStatus(false)
	fmt.Println("Status : ", denish.GetStatus())

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (user User) GetStatus() bool {
	return user.Status
}

func (user *User) SetStatus(status bool) {
	user.Status = status
}
