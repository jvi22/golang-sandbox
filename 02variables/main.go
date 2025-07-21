package main

import "fmt"

func main() {
	var username string = "jaanvi"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.455445674541389
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//default values check
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n ", anotherVariable)

	//no var or short var declaration
	userName := "jaanvi"
	fmt.Println(userName)
	fmt.Printf("Variable is of type: %T \n ", userName)

	noOfUsers := 99
	fmt.Println(noOfUsers)
    fmt.Printf("Variable is of type: %T \n ", noOfUsers)
	
}