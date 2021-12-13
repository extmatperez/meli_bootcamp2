package internal

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type StrubStore struct{}

var transacciones string = `[
	{"transaction_code": "12345", "coin": "USD", "amount": 300.00, "emitor": "Juan", "receptor": "Enrique", "transaction_date": "12/12/21"}, 
	{"transaction_code": "54321", "coin": "Euro", "amount": 150.00, "emitor": "Miguel", "receptor": "Luis", "transaction_date": "13/12/21"}
]`

func (s *StrubStore) Read(data interface{}) error {
	return json.Unmarshal([]byte(transacciones), &data)
}

func (s *StrubStore) Write(data interface{}) error {
	return nil
}
func TestGetAll(t *testing.T) {
	store := &StrubStore{}
	repo := NewRepository(store)
	var expected []Transaction
	json.Unmarshal([]byte(transacciones), &expected)

	trans, err := repo.GetAll()

	assert.Equal(t, expected, trans)

	assert.Nil(t, err)
}
