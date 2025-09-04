// maps - type of data structure
// unordered collection of key value pairs

package main

import "fmt"

func main(){

	fmt.Println("Maps in Golang!");

	languages := make(map[string]string);

    languages["js"] = "JavaScript"
    languages["RB"] = "Ruby"
    languages["py"] = "Python"

	fmt.Println("List of all languages are: ", languages);
	fmt.Println("js is short for: ", languages["js"]);

	delete(languages, "RB")
	fmt.Println("List of languages now: ", languages);

	for key, value := range languages {
		fmt.Printf("For key %v , value is %v\n", key, value)
	}
}

