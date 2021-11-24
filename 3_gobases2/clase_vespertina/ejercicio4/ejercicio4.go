package main

func main() {

}

func obtenerFuncion(operacion string) func() {
	var minimo = func(notas []float64) float64 {
		aux := 0.0
		for _, val := range notas {
			if val < aux || aux == 0 {
				aux = val
			}
		}
		return aux
	}

	var maximo = func(notas []float64) float64 {
		aux := 0.0
		for _, val := range notas {
			if val > aux || aux == 0 {
				aux = val
			}
		}
		return aux
	}

	var promedio = func(notas []float64) float64 {
		total := 0.0
		for _, val := range notas {
			total += val
		}
		return total / float64(len(notas))
	}

	switch operacion {
	case "minimo":
		return minimo
	case "maximo":
		return maximo
	case "promedio":
		return promedio
	default:
		return nil
	}
}
