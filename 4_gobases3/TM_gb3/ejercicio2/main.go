package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	readProducts, err := os.ReadFile("../ejercicio/products_stock.txt")

	if err == nil {
		stringInfo := string(readProducts)
		info := strings.Replace(stringInfo, ";", "\t", -1)

		fmt.Print(info)

	}

}
