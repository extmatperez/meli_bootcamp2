package main

import "fmt"

func main() {
	nombre := "Andres"
	apellido := "Ghione"
	saludar(nombre, apellido)
}

func saludar(nombre, apellido string) {
	fmt.Println("Hola", nombre, apellido, "bienvenido !")
}
