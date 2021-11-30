package main

import "errors"
import "fmt"

func promedio( calificaciones ...float64) (float64, error) {
	var promedio float64
	var cantidad_notas int
	var sumatoria float64
if cantidad_notas ==0{
	return 0, errors.New("El divisor no puede ser cero.")
}
	cantidad_notas = len(calificaciones)
	for _, value := range calificaciones {
		sumatoria += value
		
	}
	

	promedio = sumatoria / float64(cantidad_notas)
	return promedio, nil
}

func main() {

	promedio,err
}
