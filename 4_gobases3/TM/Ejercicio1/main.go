package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Producto struct {
	ID       int
	PRECIO   float64
	CANTIDAD int
}

func main() {
	/*var prodList []Producto

	p1 := Producto{1, 20.0, 4}
	p2 := Producto{2, 57.4, 5}
	p3 := Producto{3, 17.5, 3}

	prodList = append(prodList, p1)
	prodList = append(prodList, p2)
	prodList = append(prodList, p3)

	listaCodificada, err := json.Marshal(prodList)*/

	list := [][]string{
		{"ID", "PRECIO", "CANT"},
		{"1", "157.99", "20"},
		{"2", "25.33", "100"},
		{"3", "115", "500"},
	}
	csvFile, err := os.Create("../Files/ProdList.csv")
	if err == nil {
		csvwriter := csv.NewWriter(csvFile)
		for _, empRow := range list {
			csvwriter.Write(empRow)
		}
		csvwriter.Flush()
		csvFile.Close()
	} else {
		fmt.Print("Error al abrir archivo")
	}
}
