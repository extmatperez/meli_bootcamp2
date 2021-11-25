package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type producto struct {
	id       int
	precio   float64
	cantidad int
}

func main() {
	p1 := producto{
		id:       111223,
		precio:   30012.00,
		cantidad: 1,
	}

	p2 := producto{
		id:       444321,
		precio:   1000000.00,
		cantidad: 4,
	}
	fmt.Printf("ID \t\tPrecio \t\tCantidad\n")
	var lista []producto
	//text := fmt.Sprint(p1.id, "\t\t", p1.precio, "\t\t", p1.cantidad)
	lista = append(lista, p1)
	lista = append(lista, p2)
	p1formateado, err := json.Marshal(lista)

	err = os.WriteFile("./archivo/myFile.txt", p1formateado, 0644)

	if err != nil {
		fmt.Println("No se pudo escribir")
	}

}
