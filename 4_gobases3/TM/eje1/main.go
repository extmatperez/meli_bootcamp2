package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Productos struct {
	Id       int     `json:"id"`
	Precio   float64 `json:"precio"`
	Cantidad int     `json:"cantidad"`
}

func escribirArhivos(lista []Productos) {

	listaFormateada, err := json.Marshal(lista)
	err = os.WriteFile("./GOBasesTM.txt", listaFormateada, 0644)
	if err != nil {
		fmt.Println("No se pudo escribir")
	} else {
		fmt.Println("Grabo")
	}
}

func main() {

	producto1 := Productos{1, 550.25, 20}
	producto2 := Productos{2, 2550.25, 10}
	producto3 := Productos{3, 800.00, 2}
	producto4 := Productos{4, 15000.00, 100}
	producto5 := Productos{5, 300.00, 6}

	var listaProductos []Productos

	listaProductos = append(listaProductos, producto1)
	listaProductos = append(listaProductos, producto2)
	listaProductos = append(listaProductos, producto3)
	listaProductos = append(listaProductos, producto4)
	listaProductos = append(listaProductos, producto5)

	escribirArhivos(listaProductos)

}
