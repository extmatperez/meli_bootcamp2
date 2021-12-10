package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestDivision(t *testing.T) {
	_, err := Dividir(5, 0)

	assert.Nil(t, err)
}