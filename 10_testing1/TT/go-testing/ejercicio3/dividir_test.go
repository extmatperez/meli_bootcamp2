package dividir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDividir(t *testing.T) {
	_, err := Dividir(10, 0)

	assert.Nil(t, err)
}
