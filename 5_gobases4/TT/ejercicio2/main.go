package main

import (
	"fmt"
	"os"
)

type Customer struct {
	Legajo    int    `json:"legajo"`
	Nombre    string `json:"nombre"`
	Apellido  string `json:"apellido"`
	DNI       int    `json:"dni"`
	Telefono  int    `json:"telefono"`
	Domicilio string `json:"domicilio"`
}

func main() {
	defer func() {
		fmt.Println("No han quedado archivos abiertos")
	}()

	defer func() {
		fmt.Printf("Se detectaron varios errores en tiempo de ejecución\n")
	}()

	os.Setenv("ARCHIVO", "customers.json")

	LeerArchivo()
	cust := Customer{123, "aaa", "bbb", 12, 232132, "Calle falsa 123"}
	exists, err := CheckExist(cust.Legajo)
	if err != nil || exists {
		fmt.Println(err)
	} else {
		registerCustomer(&cust)
	}

	fmt.Printf("Fin de la ejecución\n")
}
