package main

import (
	"fmt"
	"strings"
)

func main() {
	var palabra string = "Talleres"

	fmt.Println(len(palabra))

	for i := 0; i < len(palabra); i++ {

		array := strings.Split(palabra, "")
		fmt.Printf("\n%v", array[i])
	}
}
