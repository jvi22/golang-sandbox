package main
import "fmt"

func main(){

fmt.Println("Welcome to a class on Pointers: ");
      
   myNumber := 23 

   var ptr = &myNumber 

   fmt.Println("value of actual pointer is: ", ptr);
   fmt.Println("value of actual pointer is: ", *ptr);

   *ptr = *ptr + 2 
   fmt.Println("New value is: ", myNumber);
   
}

// ptr mai koi change hoga mtlb, woh jis value ko point
// kar rha hai, usme bhi chnages honge.
// jaise *ptr mai changes hue, to actual value 
// mai bhi chnages hue. 
