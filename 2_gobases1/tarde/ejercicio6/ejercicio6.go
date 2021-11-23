package main

import "fmt"

var employees = map[string]int{
	"Benjamin": 20, "Nahuel": 26,
	"Brenda": 19, "Darío": 44,
	"Pedro": 30,
}

func queEdadTiene(nombre string) {
	for key, edad := range employees {
		if key == nombre {
			fmt.Printf("La edad de %v es %v \n", nombre, edad)
		}
		//fmt.Println("El empleado no existe")
	}
}

func mayores(edad int) {
	var empleadosMayores []string
	for key, value := range employees {
		if value > edad {
			empleadosMayores = append(empleadosMayores, key)
		}
	}
	fmt.Printf("Los mayores a %v son %v \n", edad, empleadosMayores)
}

func agregarEmpleado(nombre string, edad int) {
	employees[nombre] = edad
	fmt.Printf("Los empleados son %v ,%v\n", len(employees), employees)
}

func borrarEmpleados(nombre string) {
	delete(employees, nombre)
	fmt.Printf("Los empleados ahora son %v, %v\n", len(employees), employees)
}

func main() {
	queEdadTiene("Benjamin")
	mayores(21)
	agregarEmpleado("Federico", 25)
	borrarEmpleados("Pedro")
}
