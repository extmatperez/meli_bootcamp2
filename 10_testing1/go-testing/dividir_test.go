package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	//Arrange
	num1 := 4
	num2 := 2
	num3 := 0

	expectedResult := 2

	//Act
	result1, err1 := Dividir(num1, num2)
	result2, err2 := Dividir(num1, num3)

	//Assert
	assert.Nil(t, err1)
	assert.Equal(t, expectedResult, result1)

	assert.NotNil(t, err2)
	assert.Equal(t, 0, result2)
}
