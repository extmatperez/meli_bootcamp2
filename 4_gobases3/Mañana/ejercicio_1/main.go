package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type producto struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

func main() {
	p1 := producto{
		Id:       99,
		Precio:   30012.00,
		Cantidad: 1,
	}

	p2 := producto{
		Id:       100,
		Precio:   55000.00,
		Cantidad: 4,
	}
	//fmt.Printf("ID \t\tPrecio \t\tCantidad\n")
	var lista []producto
	//text := fmt.Sprint(p1.id, "\t\t", p1.precio, "\t\t", p1.cantidad)
	lista = append(lista, p1)
	lista = append(lista, p2)
	p1formateado, err := json.Marshal(lista)
	// fmt.Println(lista)
	// fmt.Println(p1formateado)
	err = os.WriteFile("./myFile.txt", p1formateado, 0644)

	if err != nil {
		fmt.Println("No se pudo escribir")
	}

}
