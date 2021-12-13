package internal

import (
	"errors"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/palacio_francisco/11_testing2/TurnoTarde/pkg/store"
	"github.com/stretchr/testify/assert"
)

var transactons string =  `[{
	"id": 2,
	"codigo": "24safdsadfasdf385",
	"moneda": "Peso Colombiano",
	"monto": "$8228845654645678",
	"emisor": "Luis",
	"receptor": "Perez",
	"fecha": "01/01/2001"
   },
   {
	"id": 3,
	"codigo": "11673-417",
	"moneda": "Franc",
	"monto": "$2.76",
	"emisor": "minstone2",
	"receptor": "sinnott2",
	"fecha": "1/4/2021"
   }]`


func TestUpdate(t *testing.T){
	transacionTest := Transaction{
		Codigo: "New001",
		Moneda: "ARS",
		Monto: "850",
		Emisor: "Luis",
		Receptor: "Ppepe",
		Fecha: "13/12/2021",
	}

	mock := store.Mock{Data: []byte("[]"),IsStoreRead: false}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService:= NewService(myRepo)

	newTransaction,err := myService.Store(transacionTest.Codigo,transacionTest.Moneda,transacionTest.Monto,
										transacionTest.Emisor,transacionTest.Receptor,transacionTest.Fecha)

	assert.Nil(t,err)
	assert.True(t,mock.IsStoreRead)
	assert.Equal(t,transacionTest.Codigo,newTransaction.Codigo)
	assert.Equal(t,transacionTest.Emisor,newTransaction.Emisor)
	assert.Equal(t,transacionTest.Fecha,newTransaction.Fecha)
	assert.Equal(t,transacionTest.Moneda,newTransaction.Moneda)
	assert.Equal(t,transacionTest.Monto,newTransaction.Monto)
	assert.Equal(t,transacionTest.Receptor,newTransaction.Receptor)

}


func TestUpdateError(t *testing.T){
	transacionTest := Transaction{
		Codigo: "New001",
		Moneda: "ARS",
		Monto: "850",
		Emisor: "Luis",
		Receptor: "Ppepe",
		Fecha: "13/12/2021",
	}
	erCreated:= errors.New("Error al updtear transaction")
	mock := store.Mock{Data: []byte("[]"),Err: erCreated}
	typeFileMock := store.FileStore{Mock: &mock}
	myRepo := NewRepository(&typeFileMock)
	myService:= NewService(myRepo)

	newTransaction,err := myService.Store(transacionTest.Codigo,transacionTest.Moneda,transacionTest.Monto,
										transacionTest.Emisor,transacionTest.Receptor,transacionTest.Fecha)

	assert.NotNil(t,err,erCreated)
	assert.Equal(t,Transaction{},newTransaction)

}