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
	filas, _ := r.ReadAll()
	fmt.Println(filas)
	formatter(filas)
}

func formatter(filas [][]string) {
	fmt.Printf("%-10v %10v %10v\n", "Id", "Precio", "Cantidad")
	total := 0.00
	for _, fila := range filas {
		formattedRows := strings.Split(fila[0], ";")
		fmt.Printf("%-10s %10s %10s\n", formattedRows[0], formattedRows[1], formattedRows[2])
		parsedFloat1, err := strconv.ParseFloat(formattedRows[1], 64)
		parsedFloat2, err2 := strconv.ParseFloat(formattedRows[2], 64)
		if err != nil || err2 != nil {
			fmt.Println(err, err2)
		} else {
			total += (parsedFloat1 * parsedFloat2)
		}
	}
	fmt.Printf("%21v\n", total)
}
