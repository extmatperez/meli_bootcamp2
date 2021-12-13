package main

import "fmt"

func main() {
	primer_usuario := Usuario{"a", "a", 30, "a@a.com", "asd1"}
	segundo_usuario := Usuario{"b", "b", 30, "b@b.com", "asd1"}
	tercer_usuario := Usuario{"c", "c", 30, "c@c.com", "asd1"}
	cuarto_usuario := Usuario{"d", "d", 30, "d@d.com", "asd1"}
	quinto_usuario := Usuario{"e", "e", 30, "e@e.com", "asd1"}

	fmt.Printf("%v\n", primer_usuario)
	fmt.Printf("%v\n", segundo_usuario)
	fmt.Printf("%v\n", tercer_usuario)
	fmt.Printf("%v\n", cuarto_usuario)
	fmt.Printf("%v\n\n", quinto_usuario)

	cambiar_nombre("nuevo a", &primer_usuario)
	cambiar_edad(35, &segundo_usuario)
	cambiar_correo("z@z.com", &tercer_usuario)
	cambiar_pass("asd2", &cuarto_usuario)
	cambiar_nombre("nuevo e", &quinto_usuario)

	fmt.Printf("%v\n", primer_usuario)
	fmt.Printf("%v\n", segundo_usuario)
	fmt.Printf("%v\n", tercer_usuario)
	fmt.Printf("%v\n", cuarto_usuario)
	fmt.Printf("%v\n", quinto_usuario)
}

type Usuario struct {
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Edad     int    `json:"edad"`
	Correo   string `json:"correo"`
	Pass     string `json:"pass"`
}

func cambiar_nombre(nombre string, u *Usuario) {
	u.Nombre = nombre
}

func cambiar_edad(edad int, u *Usuario) {
	u.Edad = edad
}

func cambiar_correo(correo string, u *Usuario) {
	u.Correo = correo
}

func cambiar_pass(pass string, u *Usuario) {
	u.Pass = pass
}
