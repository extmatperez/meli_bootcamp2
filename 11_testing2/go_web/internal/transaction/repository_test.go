package internal

import (
	"encoding/json"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/arevalo_ivan/11_testing2/go_web/pkg/store"
	"github.com/stretchr/testify/assert"
)

var transaccionesParaMock []Transaction = []Transaction{
	{ID: 1, Transaction_Code: "12345", Coin: "USD", Amount: 300.00, Emitor: "Juan", Receptor: "Enrique", Transaction_Date: "12/12/21"},
	{ID: 2, Transaction_Code: "54321", Coin: "Euro", Amount: 150.00, Emitor: "Miguel", Receptor: "Luis", Transaction_Date: "13/12/21"},
}

func TestUpdateMockStores(t *testing.T) {
	var newTransaction Transaction = Transaction{
		Transaction_Code: "777",
		Coin:             "Peso",
		Amount:           199.00,
		Emitor:           "Rogelio",
		Receptor:         "Funes Mori",
		Transaction_Date: "12/12/21",
	}
	transByte, _ := json.Marshal(transaccionesParaMock)

	mock := store.Mock{Data: transByte}
	filestore := store.FileStore{Mock: &mock}
	repo := NewRepository(&filestore)
	service := NewService(repo)

	updatedTransaction, err := service.Update(1, newTransaction.Transaction_Code,
		newTransaction.Coin, newTransaction.Emitor, newTransaction.Receptor,
		newTransaction.Transaction_Date, newTransaction.Amount)

	errTransaction, _ := service.Update(77, newTransaction.Transaction_Code,
		newTransaction.Coin, newTransaction.Emitor, newTransaction.Receptor,
		newTransaction.Transaction_Date, newTransaction.Amount)

	assert.Equal(t, newTransaction.Coin, updatedTransaction.Coin)

	assert.Equal(t, Transaction{}, errTransaction)

	assert.Nil(t, err)

}

func TestDeleteMockStore(t *testing.T) {
	transByte, _ := json.Marshal(transaccionesParaMock)
	mock := store.Mock{Data: transByte}
	filestore := store.FileStore{Mock: &mock}
	repo := NewRepository(&filestore)
	service := NewService(repo)

	err := service.Delete(1)

	assert.Nil(t, err, "Se borro exitosamente")

	erasedTransactions, _ := service.GetAll()
	assert.Equal(t, len(transaccionesParaMock)-1, len(erasedTransactions), "Deberian tener el mismo length")
}

// transactionTest := Transaction{1, "12345", "USD", 300.00, "Juan", "After Update", "12/12/21"}

// transUpdated, err := repo.UpdateReceptor(1, "After Update")

// assert.Equal(t, myMock.BeforeUpdate, transUpdated)
// assert.Nil(t, err)

// assert.Equal(t, true, changed)
