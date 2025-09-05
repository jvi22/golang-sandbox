package main

import "fmt"

func main() {

	fmt.Println("Structs in golang.");

	jaanvi := User{"Jaanvi", "jvigo@gmai.com", true, 19}
	fmt.Println(jaanvi);
	fmt.Printf("jaanvi's details are: %+v", jaanvi);
	//no inheritance in Go
	// no super or parent
}

type User struct {
	Name   string
	Email  string
	Status bool 
	Age    int
	
}