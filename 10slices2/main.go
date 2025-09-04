package main

import "fmt"

func main() {

	courses := []string{"python", "reactjs", "java", "golang"}
	fmt.Println(courses);

	indexToRemove := 2 ;
	// +1 tells Go to start the second slice after the element we want to remove.
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)
    fmt.Println(courses);

}