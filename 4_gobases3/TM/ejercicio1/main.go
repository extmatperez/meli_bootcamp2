package main

import (
	"fmt"
	"os"
)

type Producto struct {
	Id       int     `json:id`
	Precio   float64 `json:precio`
	Cantidad int     `json:cantidad`
}

func main() {
	prod1 := Producto{123, 100.00, 2}
	csvProd1 := fmt.Sprintf("%v;%v;%v\n", prod1.Id, prod1.Precio, prod1.Cantidad)
	err2 := os.WriteFile("./myFile.csv", []byte(csvProd1), 0644)
	fmt.Println(err2)
}
