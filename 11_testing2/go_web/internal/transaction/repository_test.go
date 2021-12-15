package internal

import (
	"encoding/json"
	"errors"
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

	updatedTransaction2, errTransaction := service.Update(77, newTransaction.Transaction_Code,
		newTransaction.Coin, newTransaction.Emitor, newTransaction.Receptor,
		newTransaction.Transaction_Date, newTransaction.Amount)

	errorTest := errors.New("La transacción 77 no existe")

	assert.Equal(t, newTransaction.Coin, updatedTransaction.Coin)

	assert.Equal(t, errorTest, errTransaction)

	assert.Nil(t, err)

	assert.Equal(t, Transaction{}, updatedTransaction2)

}

func TestDeleteMockStore(t *testing.T) {
	transByte, _ := json.Marshal(transaccionesParaMock)
	mock := store.Mock{Data: transByte}
	filestore := store.FileStore{Mock: &mock}
	repo := NewRepository(&filestore)
	service := NewService(repo)

	err := service.Delete(1)

	assert.Nil(t, err, "Se borro exitosamente")
}

func TestDeleteErrMockStore(t *testing.T) {
	transByte, _ := json.Marshal(transaccionesParaMock)
	mock := store.Mock{Data: transByte}
	filestore := store.FileStore{Mock: &mock}
	repo := NewRepository(&filestore)
	service := NewService(repo)

	err := service.Delete(3)

	errorTest := errors.New("La transacción 3 no existe")

	assert.Equal(t, errorTest, err)

}

func TestGetAllMockStore(t *testing.T) {
	transByte, _ := json.Marshal(transaccionesParaMock)
	mock := store.Mock{Data: transByte}
	filestore := store.FileStore{Mock: &mock}
	repo := NewRepository(&filestore)
	service := NewService(repo)

	transactionsGetAll, err := service.GetAll()

	assert.Equal(t, transaccionesParaMock, transactionsGetAll)
	assert.Nil(t, err)

}

func TestUpdateReceptorMockStore(t *testing.T) {
	transByte, _ := json.Marshal(transaccionesParaMock)
	mock := store.Mock{Data: transByte}
	filestore := store.FileStore{Mock: &mock}
	repo := NewRepository(&filestore)
	service := NewService(repo)

	var transactionTest Transaction = Transaction{
		ID:               1,
		Transaction_Code: "12345",
		Coin:             "USD",
		Amount:           300.00,
		Emitor:           "Juan",
		Receptor:         "Mariano",
		Transaction_Date: "12/12/21",
	}

	transactionReceptorUpdated, err := service.UpdateReceptor(1, "Mariano")

	assert.Equal(t, transactionTest, transactionReceptorUpdated)
	assert.Nil(t, err)

}
func TestUpdateReceptorErrorMockStore(t *testing.T) {
	transByte, _ := json.Marshal(transaccionesParaMock)
	mock := store.Mock{Data: transByte}
	filestore := store.FileStore{Mock: &mock}
	repo := NewRepository(&filestore)
	service := NewService(repo)

	transactionReceptorUpdated, err := service.UpdateReceptor(3, "Mariano")

	errorTest := errors.New("La transacción 3 no existe")

	assert.Equal(t, errorTest, err)
	assert.Equal(t, Transaction{}, transactionReceptorUpdated)

}

func TestStoreMockStore(t *testing.T) {
	transByte, _ := json.Marshal(transaccionesParaMock)
	mock := store.Mock{Data: transByte}
	filestore := store.FileStore{Mock: &mock}
	repo := NewRepository(&filestore)
	service := NewService(repo)

	var transactionTest Transaction = Transaction{
		ID:               3,
		Transaction_Code: "90909",
		Coin:             "Real",
		Amount:           300.00,
		Emitor:           "Peter",
		Receptor:         "Parker",
		Transaction_Date: "12/12/21",
	}

	transactionStored, err := service.Store(transactionTest.Transaction_Code,
		transactionTest.Coin, transactionTest.Emitor, transactionTest.Receptor,
		transactionTest.Transaction_Date, transactionTest.Amount)

	assert.Equal(t, transactionTest, transactionStored)
	assert.Nil(t, err)

}
