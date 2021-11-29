package main

import (
	"fmt"
	"os"
)

func openFile(file string) {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}

		//Println execution ended can be here too
	}()

	_, err := os.ReadFile(file)
	if err != nil {
		panic("file not found or broken")
	}
}

func main() {

	openFile("../customers.txt")

	fmt.Println("Execution ended")
}
