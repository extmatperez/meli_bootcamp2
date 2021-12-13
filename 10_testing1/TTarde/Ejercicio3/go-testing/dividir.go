package dividir

import "errors"

func Dividir(num, den int) (int, error){
	if den == 0 {
		return 0, errors.New("El denominador no puede ser cero")
	}

	return num/den, nil
}