package dividir

// Se importa el package testing
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiv(t *testing.T) {
	// Se inicializan los datos a usar en el test (input/output)

	num1 := 10
	num2 := 2
	// Se ejecuta el test
	resultado, err := Dividir(num1, num2)
	assert.NoError(t, err)
	// Se validan los resultados
	assert.NotNil(t, resultado)
	/*
			Cambiar el método para que no sólo retorne un entero sino también un error. Incorporar una validación en la que si el denominador es igual a 0,  retornar un error cuyo mensaje sea “El denominador no puede ser 0”. Diseñar un test unitario que valide el error cuando se invoca con 0 en el denominador.
		Dentro de la carpeta go-testing crear un archivo dividir.go con la función a probar.
		Dentro de la carpeta go-testing crear un archivo dividir_test.go con el test diseñado.
	*/
}
