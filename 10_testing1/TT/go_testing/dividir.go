package calculadora

import "fmt"

func Divide(num1, den int) (interface{}, error) {
	if den == 0 {
		return nil, fmt.Errorf("El denominador no puede ser 0")
	}
	return num1 / den, nil
}
