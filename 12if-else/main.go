package main

import "fmt"

func main() {

	fmt.Println("Learning If-else in Golang!");

	loginCount := 50;
	var result string;

	if loginCount <= 50 {
		result = "Regular User";
	} else {
		result = "Watch out";
	}
	fmt.Println(result);
    
//even odd

	if 8%2 == 0 {
		fmt.Println("number is even.");
	} else {
		fmt.Println("number is odd.");
	}

// we can declare and write conditions at the same line

    num := 3; if num < 10 {
		fmt.Println("num less than 10.");
	} else {
		fmt.Println("num is equals or greater than 10.");
	}

}