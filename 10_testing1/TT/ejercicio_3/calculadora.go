package calculadora

import "fmt"

func Dividir(num1 float64, num2 float64) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("División por cero")
	}
	return (num1 / num2), nil
}
