package main

import "fmt"

func main() {
	var word string = "MELI"

	lenght := len(word)

	fmt.Println("Number of letters:", lenght)

	for _, c := range word {
		fmt.Println(string(c))
	}

}
