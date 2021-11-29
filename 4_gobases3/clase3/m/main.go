package main

import (
	"fmt"
	"os"
	"strings"
)

type product struct {
	Id       int
	Price    float64
	Quantity int
}

func writeCSV() {
	products := []product{{Id: 1, Price: 10.00, Quantity: 1}, {Id: 2, Price: 22.00, Quantity: 2}, {Id: 3, Price: 33.33, Quantity: 3}}

	productsText := ""

	for _, p := range products {
		productsText += fmt.Sprintf("%d;%.2f;%d\n", p.Id, p.Price, p.Quantity)
	}

	bytes := []byte(productsText)

	err := os.WriteFile("./ej1.txt", bytes, 0644)

	if err != nil {
		fmt.Println("error")
	}
}

func readCSV() {
	data, err := os.ReadFile("./ej1.txt")

	if err != nil {
		return
	}

	lines := string(data)

	rows := strings.Split(lines, "\n")
	fmt.Println("Id\tPrecio\tCantidad")
	for _, row := range rows[:len(rows)-1] {
		fields := strings.Split(row, ";")
		newLine := ""
		for _, field := range fields {
			newLine += fmt.Sprintf("%v\t", field)
		}
		fmt.Println(newLine)
	}

}

func main() {
	writeCSV()
	readCSV()
}
