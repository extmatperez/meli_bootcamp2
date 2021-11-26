package main

// FALTA EL PUNTO 1

import (
	"errors"
	"fmt"
	"os"
)

type Cliente struct {
	Legajo    int
	Nombre    string
	Apellido  string
	DNI       int
	Numero    int
	Domicilio string
}

// func generarId() (int, error) {
// 	return rand.Intn(9999)
// }

func verificarExistenciaCliente(cliente Cliente) {
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
	// return err
}

func verificarDatosCliente(cliente Cliente) error {
	if cliente.Apellido == "" ||
		cliente.Nombre == "" ||
		cliente.Legajo == 0 ||
		cliente.DNI == 0 ||
		cliente.Numero == 0 ||
		cliente.Domicilio == "" {
		err := errors.New("alguno de los valores del cliente no fue ingresado")
		return err
	}

	return nil
}

func main() {
	//Declaración de funciones defer para mensajes finales
	defer func() {
		fmt.Println("No han quedado archivos abiertos")
	}()
	defer func() {
		fmt.Println("Se detectaron varios errores en tiempo de ejecución")
	}()
	defer func() {
		fmt.Println("Fin de la ejecución")
	}()
	//------------------------------------------------------------

	cliente := Cliente{Legajo: 123456,
		Nombre:    "Facundo",
		Apellido:  "Bouza",
		DNI:       41332191,
		Numero:    12345678,
		Domicilio: ""}

	verificarExistenciaCliente(cliente)

	err := verificarDatosCliente(cliente)
	if err != nil {
		fmt.Println(err)
	}

}
