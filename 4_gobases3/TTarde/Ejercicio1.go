package main

import "fmt"

type Usuario struct{
	Nombre string
	Apellido string
	Edad int
	Correo string
	Contraseña string
}

func (u *Usuario) cambiar_nombre(nuevo string){
	u.Nombre = nuevo
}

func (u *Usuario) cambiar_edad(nuevo int){
	u.Edad = nuevo
}

func (u *Usuario) cambiar_correo(nuevo string){
	u.Correo = nuevo
}

func (u *Usuario) cambiar_contraseña(nuevo string){
	u.Contraseña = nuevo
}

func main(){
fmt.Println("Terminado")
}