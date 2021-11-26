package main

import (
	"bufio"
	"fmt"
	"os"
)

var string_buffer = bufio.NewReader(os.Stdin)

type Usuario struct {
	Nombre     string `json:"nombre"`
	Apellido   string `json:"apellido"`
	Edad       int    `json:"edad"`
	Correo     string `json:"correo"`
	Contrasena string `json:"contrasena"`
}

func (u *Usuario) cambiar_nombre() {
	var name string
	var surname string
	fmt.Println("Ingrese el nuevo nombre: ")
	name, _ = string_buffer.ReadString('\n')
	fmt.Println("Ingrese el nuevo apellido: ")
	surname, _ = string_buffer.ReadString('\n')
	u.Nombre = name
	u.Apellido = surname
}

func (u *Usuario) cambiar_edad() {
	var age int
	fmt.Println("Ingrese la nueva edad: ")
	fmt.Scanf("%d", &age)
	u.Edad = age
}

func (u *Usuario) cambiar_correo() {
	var email string
	fmt.Println("Ingrese el nuevo correo: ")
	email, _ = string_buffer.ReadString('\n')
	u.Correo = email
}

func (u *Usuario) cambiar_contrasena() {
	var password string
	fmt.Println("Ingrese la nueva contraseña: ")
	password, _ = string_buffer.ReadString('\n')
	u.Contrasena = password
}

func main() {
	u := Usuario{"Rodrigo", "Vega Gimenez", 25, "rvega389@gmail.com", "p455w0rd"}
	fmt.Println(u)
	u.cambiar_nombre()
	u.cambiar_edad()
	u.cambiar_correo()
	u.cambiar_contrasena()
	fmt.Println(u)
}
