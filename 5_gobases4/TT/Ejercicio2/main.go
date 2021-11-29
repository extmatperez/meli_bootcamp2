package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Cliente struct {
	Legajo           int
	Nombre, Apellido string
	DNI              int
	Telefono         int
	Domicilio        string
}

func generarLegajo() (legajo int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return r1.Intn(100) * 123
	// return 0
}

func validarUsuario(legajo int) {

	defer func() {
		recover()
		fmt.Println("error: el archivo indicado no fue encontrado o está dañado")
	}()

	_, err := os.Open("custom.txt")

	if err != nil {
		panic("error: El archivo indicado no fue encontrado o está dañado")
	}
}

func validarDatosUsuario() error {
	if cliente.Legajo == 0 || cliente.DNI == 0 || cliente.Telefono == 0 {

	}

}

func main() {
	//clientes := []Cliente{}
	newLegajo := generarLegajo()
	if newLegajo == 0 {
		panic("Legajo inválido")
	} else {
		fmt.Println(newLegajo)
	}
	c1 := Cliente{newLegajo, "Nicolas", "Aponte", 134394, 2342343, "sdfsd"}
	validarUsuario(c1.Legajo)

	validarDatosUsuario(c1)

}
