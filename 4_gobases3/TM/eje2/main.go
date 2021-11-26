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

func leerArchivo() {
	data, err := os.ReadFile("../eje1/GOBasesTM.txt")
	if err == nil {
		var pListaProductos []Productos
		json.Unmarshal(data, &pListaProductos)
		imprimirLista(pListaProductos)
	} else {
		fmt.Println("No se pudo leer")
	}
}

func imprimirLista(lista []Productos) {
	fmt.Println()
	fmt.Println("- - - - - - - - - - - - - - - ")
	fmt.Printf("%-10v", "ID")
	fmt.Printf("%10v", "Precio")
	fmt.Printf("%10v", "Cantidad")
	fmt.Println()
	sumaTotal := 0.00
	for _, val := range lista {
		fmt.Printf("%-10v", val.Id)
		fmt.Printf("%10v", val.Precio)
		fmt.Printf("%10v", val.Cantidad)
		fmt.Println()
		sumaTotal = sumaTotal + val.Precio
	}
	fmt.Printf("%20v", sumaTotal)
	fmt.Println()
}

func main() {

	leerArchivo()

}
