package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.ReadFile("./archivo/customers.txt")
	fmt.Println(file)
	fmt.Println(err)

}
