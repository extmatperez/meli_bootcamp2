package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
)

func GenerateID() int { //tarea 1
	return rand.Intn(50000)
}

func CheckExist(ID int) (res bool, err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	data, err := LeerArchivo()
	if err != nil {
		res = false
		panic("error: el archivo indicado no fue encontrado o está dañado")
	} else {
		for _, customer := range data {
			if customer.Legajo == ID {
				res = true
				return
			}
		}
		res = false
		return
	}
}

func checkValues(legajo int, nombre, apellido string, dni int, telefono int, domicilio string) (bool, error) {
	valid := legajo != 0 && nombre != "" && apellido != "" && dni != 0 && telefono != 0 && domicilio != ""
	var err error
	if valid {
		err = nil
	} else {
		err = fmt.Errorf("Todos los datos son obligatorios")
	}
	return valid, err
}

func registerCustomer(customer *Customer) error {
	data, err := LeerArchivo()
	if err != nil {
		data = make([]Customer, 0)
	}
	data = append(data, *customer)
	dataJson, err := json.Marshal(data)
	if err != nil {
	}
	os.WriteFile(os.Getenv("ARCHIVO"), dataJson, 0644)
	return err
}
