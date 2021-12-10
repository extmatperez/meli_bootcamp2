package calculadora

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRestar(t *testing.T) {

	num1 := 5
	num2 := 3
	resutlEsp := 2

	resultado := Restar(num1, num2)

	assert.Equal(t, resutlEsp, resultado, "la resta esta mal hecha")
}
