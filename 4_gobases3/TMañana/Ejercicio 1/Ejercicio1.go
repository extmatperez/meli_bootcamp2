package main

import (
	"fmt"
	"os"
)


type Producto struct {
	Id int
	Precio float64
	Cantidad int
}



func main() {

	p1 := Producto{Id: 1, Precio:255.99, Cantidad:25}
	p2 := Producto{Id: 2, Precio:201.99, Cantidad:15}

	list := make([]Producto,0)
	list = append(list, p1, p2)

	var str string
	
	for i := 0; i < len(list); i++ {
	str += fmt.Sprintf("%v %v %v;\n", list[i].Id, list[i].Precio, list[i].Cantidad)
	}

	_ = os.WriteFile("Productos.txt", []byte(str), 0644)
}