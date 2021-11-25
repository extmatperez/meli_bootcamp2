package main

import (
	"encoding/json"
	"os"
)


type Producto struct {
	Id int
	Precio float64
	Cantidad float64
}



func main() {

	p1 := Producto{Id: 1, Precio:255.99, Cantidad:25}
	p2 := Producto{Id: 2, Precio:201.99, Cantidad:15}

	list := make([]Producto,0)
	list = append(list, p1, p2)

	u, _ := json.Marshal(list)

	_ = os.WriteFile("Productos.txt", u, 0667)
}