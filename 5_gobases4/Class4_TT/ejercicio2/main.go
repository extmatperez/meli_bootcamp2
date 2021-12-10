package main

import (
	"errors"
	"fmt"
	"os"
)

var lastID int

type cliente struct {
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       string
	Telefono  string
	Domicilio string
}

func setId(ultimoId int) (int, error) {
	return ultimoId + 1, nil
}

func main() {

	defer func() {
		fmt.Printf("Fin de la ejecucion")
	}()

	legajo, err := setId(0)
	if err != nil {
		panic("error en la generación del legajo")
	}
	c1 := cliente{
		Legajo: legajo,
	}

	validarExCliente()

	err = validarDatoCliente(c1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(c1)

}

func validarExCliente() error {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	data, err := os.ReadFile("./customers.txt")
	if err != nil {
		panic("error: el archivo indicado no fue encontrado o está dañado")
	} else {
		fmt.Println(data)
	}
	return nil
}

func validarDatoCliente(c cliente) error {
	if c.Nombre == "" ||
		c.Apellido == "" ||
		c.DNI == "" ||
		c.Telefono == "" ||
		c.Domicilio == "" {
		err := errors.New("Falta alguno de los datos del cliente")
		return err
	}

	return nil
}
