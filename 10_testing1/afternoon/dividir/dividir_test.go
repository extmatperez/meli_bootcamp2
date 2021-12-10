package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir1Testify(t *testing.T) {
	num1 := 3
	num2 := 0

	_, err := Dividir(num1, num2)

	assert.NotNil(t, err, "cannot be 0")
	//assert.Fail(t, "arg2", err)
}

func TestDividir2Testify(t *testing.T) {
	num1 := 3
	num2 := 3
	expected := 1

	result, _ := Dividir(num1, num2)

	assert.Equal(t, expected, result, "Not equal")
}
