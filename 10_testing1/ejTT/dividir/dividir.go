package dividir

import "errors"

func Dividir(n1, n2 int) (float64, error) {
	if n2 == 0 {
		return 0, errors.New("the division can not be between 0")
	}
	return float64(n1) / float64(n2), nil
}
