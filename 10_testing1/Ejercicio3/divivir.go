package dividir

import "errors"

func Dividir(n1, n2 float64) (float64, error) {
	if n2 == 0 {
		return 0, errors.New("el denomiador no puede ser 0")
	}
	return n1 / n2, nil
}
