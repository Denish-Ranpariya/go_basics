package main

import "fmt"

//first capital letter denotes that it is a public variable
const LoginToken string = "Change me if you can!"

func main() {
	var username string = "denish"
	fmt.Println(username)
	fmt.Printf("variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallIntValue uint8 = 255
	fmt.Println(smallIntValue)
	fmt.Printf("Variable is of type: %T \n", smallIntValue)

	var smallFloatValue float32 = 255.32454544343435
	fmt.Println(smallFloatValue)
	fmt.Printf("Variable is of type: %T \n", smallFloatValue)

	//default values and some aliases
	var intValue int
	fmt.Println(intValue)
	fmt.Printf("Variable is of type: %T \n", intValue)

	var stringValue string
	fmt.Println(stringValue)
	fmt.Printf("Variable is of type: %T \n", stringValue)

	//implicit type
	var website = "saymyname.com"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//no var style
	//this method is allowed only inside a method
	name := "denish"
	fmt.Println(name)
	fmt.Printf("Variable is of type: %T \n", name)

	fmt.Println("Login token: ", LoginToken)
}
