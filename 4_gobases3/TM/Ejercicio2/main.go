package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("./newArchive.csv")

	if err != nil {
		fmt.Println(err)
	}

	newStr := strings.Replace(string(data), ";", "\t", -1)

	fmt.Println(newStr)
}
