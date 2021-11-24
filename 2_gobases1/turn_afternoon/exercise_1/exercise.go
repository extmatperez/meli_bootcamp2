package main

import "fmt"

func main() {
	var word string = "MELI"

	lenght := len(word)

	fmt.Println("Number of letters:", lenght)

	for _, c := range word {
		fmt.Println(string(c))
	}

	for i := 0; 1 < len(word); i++ {
		fmt.Println(string(word[i]))
	}
}
