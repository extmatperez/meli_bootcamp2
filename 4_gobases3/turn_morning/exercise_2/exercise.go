package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	prod1 := Product{111223, 30012.00, 1}
	prod2 := Product{444321, 1000000.00, 4}
	prod3 := Product{434321, 50.50, 1}

	prodSlice := []Product{prod1, prod2, prod3}

	newCompany := Company{prodSlice}
	fmt.Println(newCompany)

	data, err := os.ReadFile("../exercise_1/products.csv")

	fmt.Println(data)
	if err != nil {
		fmt.Println("Error while reading file")
	}

	newData := string(data)

	newData = newData + "; 4030062.50 ;"

	newStr := strings.Replace(newData, ";", "\t", -1)

	fmt.Println(newStr)
}

type Company struct {
	Products []Product
}

type Product struct {
	ID       int
	Price    float64
	Quantity int
}
