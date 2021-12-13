package divide

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivideSuccess(t *testing.T) {
	numer := 100
	denom := 2
	expectResult := 50

	results, _ := Divide(numer, denom)

	assert.Equal(t, expectResult, results, "Divide is success")
}

func TestDivideError(t *testing.T) {
	numer := 100
	denom := 0

	_, err := Divide(numer, denom)

	assert.NotNil(t, err)
}
