package main

import "fmt"

func main() {
	notas := leerNotas()
	funcion := obtenerFuncion("maximo")
	fmt.Printf("El resultado es: %v \n", funcion(notas))
}

func leerNotas() []float64 {
	var notas []float64
	fmt.Println("Para salir escriba -1 \n")
	for {
		nota := capturarNotas()
		if nota != -1 {
			notas = append(notas, nota)
		} else {
			break
		}
		fmt.Printf("\n")
	}
	return notas
}

func capturarNotas() float64 {
	var nota float64
	fmt.Println("Ingrese la nota:")
	fmt.Scanf("%f", &nota)
	return nota
}

func obtenerFuncion(operacion string) func(notas []float64) float64 {
	switch operacion {
	case "minimo":
		return func(notas []float64) float64 {
			aux := 0.0
			for _, val := range notas {
				if val < aux || aux == 0 {
					aux = val
				}
			}
			return aux
		}
	case "maximo":
		return func(notas []float64) float64 {
			aux := 0.0
			for _, val := range notas {
				if val > aux || aux == 0 {
					aux = val
				}
			}
			return aux
		}
	case "promedio":
		return func(notas []float64) float64 {
			total := 0.0
			for _, val := range notas {
				total += val
			}
			return total / float64(len(notas))
		}
	default:
		return nil
	}
}
