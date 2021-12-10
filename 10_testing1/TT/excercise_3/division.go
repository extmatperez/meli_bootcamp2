package division

import "errors"

func Dividir(num1, num2 float64) (float64, error) {
	if num2 == 0.0 {
		return 0, errors.New("NO SE PUEDE DIVIDIR POR 0")
	}
	return num1 / num2, nil
}
