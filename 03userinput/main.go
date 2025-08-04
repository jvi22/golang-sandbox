package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){

	welcome := "Welcome to user input!"
	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please rate our Pizza between 1 & 5: ")

   
	input, _ := reader.ReadString('\n')
    fmt.Println("Thanks for rating", input);	


}