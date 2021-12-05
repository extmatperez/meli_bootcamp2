package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c1 := make(chan float64)
	c2 := make(chan float64)
	c3 := make(chan float64)

	variable1 := rand.Perm(100)
	variable2 := rand.Perm(1000)
	variable3 := rand.Perm(10000)

	go insercion(variable1, variable2, variable3, c1)
	go burbuja(variable1, variable2, variable3, c2)
	go seleccion(variable1, variable2, variable3, c3)

	fmt.Printf("\nEl timepo consumido por el algoritmo de insercion fue: %.5v segundos", <-c1)
	fmt.Printf("\nEl timepo consumido por el algoritmo de Burbuja fue: %.5v segundos", <-c2)
	fmt.Printf("\nEl timepo consumido por el algoritmo de Seleccion fue: %.5v segundos\n", <-c3)

}

func insercion(sliceCien []int, sliceMil []int, sliceDiezm []int, c chan float64) {
	var auxiliar int
	ini := time.Now()
	for i := 1; i < len(sliceCien); i++ {
		auxiliar = sliceCien[i]
		for j := i - 1; j >= 0 && sliceCien[j] > auxiliar; j-- {
			sliceCien[j+1] = sliceCien[j]
			sliceCien[j] = auxiliar
		}
	}

	for i := 1; i < len(sliceMil); i++ {
		auxiliar = sliceMil[i]
		for j := i - 1; j >= 0 && sliceMil[j] > auxiliar; j-- {
			sliceMil[j+1] = sliceMil[j]
			sliceMil[j] = auxiliar
		}
	}

	for i := 1; i < len(sliceDiezm); i++ {
		auxiliar = sliceDiezm[i]
		for j := i - 1; j >= 0 && sliceDiezm[j] > auxiliar; j-- {
			sliceDiezm[j+1] = sliceDiezm[j]
			sliceDiezm[j] = auxiliar
		}
	}

	fin := time.Now()
	tiempo := fin.Sub(ini)
	c <- float64(tiempo.Seconds())
}

func burbuja(sliceCien []int, sliceMil []int, sliceDiezm []int, c chan float64) {
	var auxiliar int
	ini := time.Now()
	for i := 0; i < len(sliceCien); i++ {
		for j := 0; j < len(sliceCien); j++ {
			if sliceCien[i] > sliceCien[j] {
				auxiliar = sliceCien[i]
				sliceCien[i] = sliceCien[j]
				sliceCien[j] = auxiliar
			}
		}
	}

	for i := 0; i < len(sliceMil); i++ {
		for j := 0; j < len(sliceMil); j++ {
			if sliceMil[i] > sliceMil[j] {
				auxiliar = sliceMil[i]
				sliceMil[i] = sliceMil[j]
				sliceMil[j] = auxiliar
			}
		}
	}

	for i := 0; i < len(sliceDiezm); i++ {
		for j := 0; j < len(sliceDiezm); j++ {
			if sliceDiezm[i] > sliceDiezm[j] {
				auxiliar = sliceDiezm[i]
				sliceDiezm[i] = sliceDiezm[j]
				sliceDiezm[j] = auxiliar
			}
		}
	}

	fin := time.Now()
	tiempo := fin.Sub(ini)
	c <- float64(tiempo.Seconds())
}

func seleccion(sliceCien []int, sliceMil []int, sliceDiezm []int, c chan float64) {
	ini := time.Now()
	for i := 0; i < len(sliceCien); i++ {
		minimo_encontrado, posicion_minimo := sliceCien[i], i
		valor_original := sliceCien[i]
		// encontrar minimo en parte desordenada
		for j := i + 1; j < len(sliceCien); j++ {
			valor_comparacion := sliceCien[j]
			if valor_comparacion < minimo_encontrado {
				minimo_encontrado, posicion_minimo = valor_comparacion, j
			}
		}
		if minimo_encontrado != valor_original {
			// intercambio posiciones con primer desordenado
			sliceCien[i], sliceCien[posicion_minimo] = minimo_encontrado, valor_original
		}
	}

	for i := 0; i < len(sliceMil); i++ {
		minimo_encontrado, posicion_minimo := sliceMil[i], i
		valor_original := sliceMil[i]
		// encontrar minimo en parte desordenada
		for j := i + 1; j < len(sliceMil); j++ {
			valor_comparacion := sliceMil[j]
			if valor_comparacion < minimo_encontrado {
				minimo_encontrado, posicion_minimo = valor_comparacion, j
			}
		}
		if minimo_encontrado != valor_original {
			// intercambio posiciones con primer desordenado
			sliceMil[i], sliceMil[posicion_minimo] = minimo_encontrado, valor_original
		}
	}

	for i := 0; i < len(sliceDiezm); i++ {
		minimo_encontrado, posicion_minimo := sliceDiezm[i], i
		valor_original := sliceDiezm[i]
		// encontrar minimo en parte desordenada
		for j := i + 1; j < len(sliceDiezm); j++ {
			valor_comparacion := sliceDiezm[j]
			if valor_comparacion < minimo_encontrado {
				minimo_encontrado, posicion_minimo = valor_comparacion, j
			}
		}
		if minimo_encontrado != valor_original {
			// intercambio posiciones con primer desordenado
			sliceDiezm[i], sliceDiezm[posicion_minimo] = minimo_encontrado, valor_original
		}
	}

	fin := time.Now()
	tiempo := fin.Sub(ini)
	c <- float64(tiempo.Seconds())
}
