package calculadora

import "errors"

func Dividir(num, dem int) (int, error) {
	if dem == 0 {
		return 0, errors.New("El denominador no puede ser 0")
	}
	return num / dem, nil
}
