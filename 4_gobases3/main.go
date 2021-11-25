package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Setenv("NAME", "valor")
	variable := os.Getenv("NAME")

	if err == nil {
		fmt.Println(variable)
	} else {
		fmt.Println("Error ", err)
	}

	variable2, ok := os.LookupEnv("NAME")
	fmt.Println(variable2, ok)
	variable2, ok = os.LookupEnv("NAME2")
	fmt.Println(variable2, ok)

	data, err := os.ReadFile("./main.go")
	fmt.Println(string(data))

}
