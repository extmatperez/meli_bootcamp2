package main

type Usuario struct {
	nombre     string
	apellido   string
	edad       int
	correo     string
	contraseña string
}

func(u *usuario) cambiarNombre(nombreNuevo, apellidoNuevo string){
	u.nombre = nombreNuevo
	u.apellido = apellidoNuevo
}

func main() {
	nuevoUsuario := Usuario(nombre: "Pepe", apellido:"Perez")

	fmt.Println(nuevoUsuario)

	nuevoUsuario.cambiarNombre("Dario", "Gonzalez")

	fmt.Println(nuevoUsuario)
}
