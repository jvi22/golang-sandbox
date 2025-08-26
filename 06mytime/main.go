package main

import "fmt"
import "time"

func main() {

	fmt.Println("Welcome to Time Study of Golang");

	presentTime := time.Now()
	
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2020, time.September, 22, 12, 30, 0, 0, time.UTC);
	
	fmt.Println(createdDate.Format("01-02-2006 Monday"));
}  	