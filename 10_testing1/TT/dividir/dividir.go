package dividir

import "fmt"

// Función que recibe dos enteros (numerador y denominador) y retorna la división resultante
func Dividir(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0")
	}
	return (num / den), nil
}
