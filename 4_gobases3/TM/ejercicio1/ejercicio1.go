package main

/*
Una empresa que se encarga de vender productos de limpieza necesita:
Implementar una funcionalidad para guardar un archivo de texto, con la información de productos comprados, separados por punto y coma (csv).
Debe tener el id del producto, precio y la cantidad.
Estos valores pueden ser hardcodeados o escritos en duro en una variable.
*/

import (
	"encoding/csv"
	"fmt"
	"os"
)

// type Producto struct {
// 	ID     int     `json:"id"`
// 	PRECIO float64 `json:"price"`
// 	CANT   int     `json:"cant"`
// }

func main() {
	// var prodList []Producto
	// p1 := Producto{1, 157.99, 20}
	// p2 := Producto{2, 25.33, 100}
	// p3 := Producto{3, 115, 500}

	// prodList = append(prodList, p1)
	// prodList = append(prodList, p2)
	// prodList = append(prodList, p3)

	// listaCodificada, err := json.Marshal(prodList)

	list := [][]string{
		{"ID", "PRECIO", "CANT"},
		{"1", "157.99", "20"},
		{"2", "25.33", "100"},
		{"3", "115", "500"},
		{"4", "215", "50"},
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
