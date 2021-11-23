package ej1

import "fmt"

func Ej1(word string) (int, []string) {
	fmt.Println(len(word))

	var letters []string

	for _, el := range word {
		fmt.Printf("%s,", string(el))
		letters = append(letters, string(el))
	}

	fmt.Println(letters)
	return len(word), letters
}
