package main

import (
	"fmt"
	"os"
	"encoding/json"
)


type Productos struct {
	Id int
	Precio float64
	Cantidad float64
}

func main() {
	var lista []Productos
	data, _:= os.ReadFile("Productos.txt")
	json.Unmarshal(data, &lista)

	var total float64

	fmt.Printf("%-10v %10v %10v", "ID", "Precio", "Cantidad")
	for i := 0; i < len(lista); i++ {
	fmt.Printf("\n%-10v %10.5v %10v", lista[i].Id, lista[i].Precio*lista[i].Cantidad, lista[i].Cantidad)
	total +=  lista[i].Precio*lista[i].Cantidad
	}
	fmt.Printf("\n%21v", total)
}