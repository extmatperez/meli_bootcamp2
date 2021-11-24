package main

import "fmt"

func main() {
	empleados := crearEmpleados()
	for _, empleado := range empleados {
		fmt.Printf("Categoria: %v", empleado.Categoria)
		sueldo := calcularSalario(empleado.HorasTrabajadas, empleado.Categoria)
		fmt.Printf("Al empleado %v le corresponde un sueldo de %v \n", empleado.Nombre, sueldo)
	}
}

type empleado struct {
	Nombre          string
	HorasTrabajadas uint32
	Categoria       string
}

func crearEmpleados() []empleado {
	empleado1 := empleado{"Ramiro Gonzalez", 15, "A"}
	empleado2 := empleado{"Francisco Mitre", 58, "B"}
	empleado3 := empleado{"Andres Otero", 50, "B"}
	empleado4 := empleado{"Maria Gimenez", 38, "C"}
	empleado5 := empleado{"Natalia Antonaz", 25, "A"}
	empleado6 := empleado{"Virginia Tojo", 14, "C"}
	empleado7 := empleado{"Tomas Montero", 10, "B"}
	empleado8 := empleado{"Nahuel Bengochea", 1, "A"}
	var empleados = []empleado{empleado1, empleado2, empleado3, empleado4, empleado5, empleado6, empleado7, empleado8}
	return empleados
}

func calcularSalario(minutosTrabajados uint32, categoria string) uint32 {
	horasTrabajadas := minutosTrabajados / 60
	return calcularSalarioPorCategoria(horasTrabajadas, categoria)
}

func calcularSalarioPorCategoria(horasTrabajadas uint32, categoria string) uint32 {
	const categoriaA = 3000
	const categoriaB = 1500
	const categoriaC = 1000

	switch categoria {
	case "A":
		sueldo := categoriaA * horasTrabajadas
		return sueldo + ((sueldo * 50) / 100)
	case "B":
		sueldo := categoriaB * horasTrabajadas
		return sueldo + ((sueldo * 20) / 100)
	case "C":
		return categoriaC * horasTrabajadas
	}
	return 0
}
