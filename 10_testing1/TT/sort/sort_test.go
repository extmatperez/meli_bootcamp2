package sort

// Se importa el package testing
import "testing"

func TestSort(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)

	nums := []int{5, 4, 3, 2, 7, 1, 9, 11}
	resultadoEsperado := []int{1, 2, 3, 4, 5, 7, 9, 11}
	// Se ejecuta el test
	resultado := Sort(nums)
	// Se validan los resultados
	for i := range resultadoEsperado {
		if resultado[i] != resultadoEsperado[i] {
			t.Errorf("Funcion Sort() arrojo el resultado = %v, pero el esperado es  %v", resultado, resultadoEsperado)
		}
	}

}
