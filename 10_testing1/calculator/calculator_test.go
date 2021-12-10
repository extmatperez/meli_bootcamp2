package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubtract(t *testing.T) {
	num1 := 10
	num2 := 5
	expectedResult := num1 - num2
	actualResult := Subtract(num1, num2)

	assert.Equal(t, expectedResult, actualResult, "must be equal")
}

func TestDivision(t *testing.T) {
	num1 := 10
	num2 := 5
	expectedResult := num1 / num2
	actualResult, err := Divide(num1, num2)
	assert.Equal(t, expectedResult, actualResult, "must be equal")
	assert.Nil(t, err)
}

func TestDivisionByZero(t *testing.T) {
	num1 := 10
	num2 := 0

	actualResult, err := Divide(num1, num2)
	assert.Equal(t, 0, actualResult)
	assert.Error(t, err)
}
