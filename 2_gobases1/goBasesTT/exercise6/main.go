package main

import "fmt"

func main() {
	var empleadoBusqueda string
	var empleadoNuevo string
	var edadEmpleadoNuevo int
	var despedirEmpleado string
	fmt.Println("Ejercicio 6")
	fmt.Println("Ingresa el nombre del empleado que queres averiguar la edad")
	fmt.Scanf("%s", &empleadoBusqueda)
	fmt.Println("Ingresa el nombre del empleado nuevo que va a ingresar")
	fmt.Scanf("%s", &empleadoNuevo)
	fmt.Println("Ingresa la edad del empleado nuevo que va a ingresar")
	fmt.Scanf("%v", &edadEmpleadoNuevo)
	fmt.Println("Ingresa el nombre del empleado nuevo que vas a despedir")
	fmt.Scanf("%s", &despedirEmpleado)
	fmt.Println(empleados(empleadoBusqueda, empleadoNuevo, edadEmpleadoNuevo, despedirEmpleado))

}

func empleados(empleadoBusqueda string, empleadoNuevo string, edadEmpleadoNuevo int, despedirEmpleado string) (int, int, map[string]int) {
	var employees = map[string]int{"benjamin": 20, "nahuel": 26, "brenda": 19, "dario": 44, "pedro": 30}

	var resultado1 int = employees[empleadoBusqueda]
	var mayores21 int

	employees[empleadoNuevo] = edadEmpleadoNuevo
	delete(employees, despedirEmpleado)

	for _, mayor := range employees {
		if mayor > 21 {
			mayores21 += 1
		}
	}

	return resultado1, mayores21, employees

}
