package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Mango", "Cherry"}
	fmt.Printf("Type of fruitlist %T \n", fruitList);

	fruitList = append(fruitList, "Banana", "Peach")
	fmt.Println(fruitList);

	fruitList = append(fruitList[1:3]);
	fmt.Println(fruitList);


    highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	// even tho default allocation of memory is 4
	//using 'append' method reallocates entire memory

	highScores = append(highScores, 555, 666, 321)
	sort.Ints(highScores);

	fmt.Println(highScores);
	fmt.Println(highScores);

	fmt.Println(sort.IntsAreSorted(highScores));
}