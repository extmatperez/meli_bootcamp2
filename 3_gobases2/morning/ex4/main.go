package main

import (
	"fmt"
	"errors"
)

// Calcular estadísticas
const (
	minimo = "minimo"
	promedio = "promedio"
	maximo = "maximo"
)

func minFunc(values ...int) (int, error) {
	if len(values) == 0 {
        return 0, errors.New("error in minFunc Function")
    }

    min := values[0]
    for _, v := range values {
            if (v < min) {
                min = v
            }
    }
    return min, nil
}

func maxFunc(values ...int) (int, error) {
	if len(values) == 0 {
        return 0, errors.New("error in maxFunc Function")
    }

    max := values[0]
    for _, v := range values {
            if (v > max) {
                max = v
            }
    }
    return max, nil
}

func avgFunc(values ...int) (int, error) {
	if len(values) == 0 {
        return 0, errors.New("error in avgFunc Function")
    }

    sum := 0
    for _, v := range values {
		sum += v
    }
    return sum/int(len(values)), nil
}

func operacion(operador string)func(vals ...int) (int, error) {
	switch operador {
	case "minimo":
		return minFunc
	case "promedio":
		return avgFunc
	case "maximo":
		return maxFunc
	}
	return nil
}

func main() {
	/* //minimo
	op := operacion(minimo)
	res, err := op(10,20,30) */

	/* //maximo 
	op := operacion(maximo)
	res, err := op(10,20,30,35,55) */

	//promedio
	op := operacion(promedio)
	res, err := op(10,20,30,35,56)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}