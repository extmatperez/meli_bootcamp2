package main

import "fmt"

func main() {

	var word string = "elefante"
	chars := []rune(word)
	letters := make([]string, 8)
	for i := 0; i < len(chars); i++ {
		letters[i] = string(chars[i])
	}
	fmt.Println(len(letters))
	for _, letter := range letters {
		fmt.Println(letter)
	}

}
