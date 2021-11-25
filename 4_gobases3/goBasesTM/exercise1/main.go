package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

func main() {
	producto := Producto{1122, 300.5, 5}
	producto1 := Producto{1123, 303.3, 10}
	producto2 := Producto{1124, 333.9, 15}

	/* 	var lista []Producto
	   	lista = append(lista, producto, producto1, producto2) */

	guardarArchivo(producto, producto1, producto2)

	fmt.Println("\n", producto, "\n", producto1, "\n", producto2)

}

type Producto struct {
	IdProd   int64   `json:"ID"`
	Precio   float64 `json:"PRECIO"`
	Cantidad int64   `json:"CANTIDAD"`
}

func guardarArchivo(prod ...Producto) {

	formProd, err := json.Marshal(prod)

	if err != nil {
		fmt.Println("Error json: ", err)
	}

	prodFile := os.WriteFile("./fileE.txt", formProd, 0644)

	if prodFile != nil {
		fmt.Println("Error", err)
	}

	//chequear esto de abajo

	data, err := os.ReadFile("./fileE.txt")
	if err == nil {
		file := string(data)
		fmt.Println(file)
	} else {
		fmt.Println("El archivo no existe...")
	}

	// Unmarshal JSON data
	var d []Producto
	err = json.Unmarshal([]byte(data), &d)
	if err != nil {
		fmt.Println(err)
	}
	// Create a csv file
	f, err := os.Create("./fileE.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// Write Unmarshaled json data to CSV file
	w := csv.NewWriter(f)
	for _, obj := range d {
		var record []string
		record = append(record, strconv.FormatInt(obj.IdProd, 10))
		record = append(record, strconv.FormatFloat(obj.Precio, 'f', -1, 32))
		record = append(record, strconv.FormatInt(obj.Cantidad, 10))
		w.Write(record)
	}
	w.Flush()

}
