package tdd

import "testing"

// Paso 1: Elegir la funcionalidad a desarrollar y analizar los casos de prueba.
// Paso 2: Escribir un unit test para el caso de prueba.
// Paso 3: Escribir el mínimo código posible para que el Test pase.
// Paso 4: Ejecutar todos los unit tests. Todos deberían pasar.
// Paso 5: Refactor, mejorar el código.
// Paso 6: Repetir paso 2 y 5 para cada caso de prueba.

func TestFibonacci(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{{0, 0}, {1, 1}, {2, 1}, {3, 2}, {4, 3}, {5, 5}}
	for i, d := range tests {
		got := fibonacci(d.arg)
		if got != d.want {
			t.Errorf("Test[%d]: factorial(%d) returned %d, want %d",
				i, d.arg, got, d.want)
		}
	}
}
