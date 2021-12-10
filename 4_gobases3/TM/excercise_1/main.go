package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Product struct {
	ID       string
	Precio   float64
	Cantidad int
}

func main() {
	producto1 := Product{"1", 2000.00, 3}
	producto2 := Product{"2", 5000.00, 8}
	var listaProductos []Product
	listaProductos = append(listaProductos, producto1)
	listaProductos = append(listaProductos, producto2)
	//err := os.WriteFile("/Users/apachon/bootcamp/meli_bootcamp2/4_gobases3/TM/excercise_1", listaProductos, 0644)
	archivo, err := os.Create("archivo.csv")
	if err != nil {
		fmt.Println(err)
	}
	archivo.Close()

	writer := csv.NewWriter(archivo)
	writer.Comma = ';'
	defer writer.Flush()
	//for _, data := range listaProductos {
	//_ =writer.Write(data)
	//}

}
