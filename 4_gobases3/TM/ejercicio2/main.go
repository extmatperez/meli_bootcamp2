package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("./ejercicio1/productos.csv")
	r := csv.NewReader(strings.NewReader(string(data)))
	rows, _ := r.ReadAll()
	fmt.Println(rows)
	formatter(rows)
}

func formatter(rows [][]string) {
	fmt.Printf("%-10v %10v %10v\n", "Id", "Precio", "Cantidad")
	total := 0.00
	for _, row := range rows {
		formattedRows := strings.Split(row[0], ";")
		fmt.Printf("%-10s %10s %10s\n", formattedRows[0], formattedRows[1], formattedRows[2])
		parsedFloat, err := strconv.ParseFloat(formattedRows[1], 64)
		if err != nil {
			fmt.Println(err)
		} else {
			total += parsedFloat
		}
	}
	fmt.Printf("%21v\n", total)
}
