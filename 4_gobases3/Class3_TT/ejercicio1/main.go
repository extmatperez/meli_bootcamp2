package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	u1 := usuario{}

	u1.setNombre("ariel", "lol")
	u1.setEdad(23)
	u1.setCorreo("arielol@meli.com")
	u1.setContraseña("1234567890")

	usuarios := []usuario{u1}
	uFormated, err := json.Marshal(usuarios)

	if err != nil {
		fmt.Println("Ha ocurrido un error")
	} else {
		fmt.Println(string(uFormated))
	}
}

type usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (u *usuario) setNombre(nombre string, apellido string) {
	u.Nombre = nombre
	u.Apellido = apellido
}

func (u *usuario) setEdad(edad int) {
	u.Edad = edad
}

func (u *usuario) setContraseña(contraseña string) {
	u.Contraseña = contraseña
}

func (u *usuario) setCorreo(correo string) {
	u.Correo = correo
}
