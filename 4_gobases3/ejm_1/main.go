package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

type Producto struct {
	id string
	price float64
	quantity int
}

var productsList []Producto

func refillList() {
	productsList = []Producto{}
	productsList = append(productsList, Producto{"1", 100, 10})
	productsList = append(productsList, Producto{"2", 200, 20})
	productsList = append(productsList, Producto{"3", 300, 30})
	productsList = append(productsList, Producto{"4", 400, 40})
	productsList = append(productsList, Producto{"5", 500, 50})
}

func generateListCSVFile() {
	refillList()
	os.Remove("products.csv")
	err := os.WriteFile("products.csv", []byte(""), 0644)
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("products.csv", os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	for _, product := range productsList {
		err := writer.Write([]string{product.id, strconv.FormatFloat(product.price, 'f', 2, 64), strconv.Itoa(product.quantity)})
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

}

func main() {
	generateListCSVFile()
}