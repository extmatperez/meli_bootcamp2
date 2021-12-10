package calculator

import "fmt"

// Función que recibe dos enteros y retorna la resta o diferencia resultante
func Divide(num, den int) (int, error) {
	if den == 0 {
		return 0, fmt.Errorf("denominator cannot be 0")
	}

	return num / den, nil
}
