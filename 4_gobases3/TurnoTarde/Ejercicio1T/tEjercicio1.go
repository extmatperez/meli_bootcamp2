package main
import "fmt"

type Alumno struct {
	nombre string
	edad int64
	correo	string
	contrasena string

} 

func detalleAlumno(alumno Alumno){
	fmt.Println(alumno)

}

func main() {
	detalleAlumno(Alumno{nombre: "francisco",edad: 20,correo:"panchoinca10@gmail.com",contrasena: "skjdfhasjkfhask"})

}

