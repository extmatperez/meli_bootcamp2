package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//La funcion siempre debe comenzar con la palabra Test
func TestSumar(t *testing.T) {
	//Arrange
	num1 := 3
	num2 := 10
	expectedResult := 13

	//Act
	result := Sumar(num1, num2)

	//Assert
	if result != expectedResult {
		t.Errorf("func Sumar failed, result: %v, expected: %v", result, expectedResult)
	}
}

//Corremos "go test" en el directorio donde esta el test
//Alternativa "go test -v"

func TestSumarConTestify(t *testing.T) {
	//Arrange
	num1 := 3
	num2 := 10
	expectedResult := 13

	//Act
	result := Sumar(num1, num2)

	//Assert con Testify
	assert.Equal(t, expectedResult, result, "Sumar is not working properly")
}