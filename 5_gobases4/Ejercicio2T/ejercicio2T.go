package main

import (
	"fmt"
	"os"
)

type Customer struct {
	legajo    int
	nombre    string
	DNI       string
	telefono  int
	domicilio string
}

func readFile() {

}
func main() {
	//var cliente1 Customer

	defer func() {
		fmt.Println("holi")
		er := recover()
		if er != nil {
			fmt.Println(er)
		}
	}()

	_, err := os.Open("./archivo/customers.txt")
	if err != nil {
		panic(err)
	}

	//fmt.Print("Escriba el legajo del cliente: ")
	//fmt.Scan(&cliente1.legajo)

	//fmt.Println(cliente1)

	fmt.Println("termino la ejecucion")
}
