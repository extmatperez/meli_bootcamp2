package tdd

// Paso 1: Elegir la funcionalidad a desarrollar y analizar los casos de prueba.
// Paso 2: Escribir un unit test para el caso de prueba.
// Paso 3: Escribir el mínimo código posible para que el Test pase.
// Paso 4: Ejecutar todos los unit tests. Todos deberían pasar.
// Paso 5: Refactor, mejorar el código.
// Paso 6: Repetir paso 2 y 5 para cada caso de prueba.

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	var total, a, b int
	a = 0
	b = 1

	for i := 2; i <= n; i++ {
		total = a + b
		a = b
		b = total
	}
	return total
	//return fibonacci(n-1) + fibonacci(n-2)
}
