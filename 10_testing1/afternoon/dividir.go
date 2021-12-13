package calculadora

import "fmt"

func Dividir(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("El denominador no puede ser 0")
	}

	return a / b, nil
}
