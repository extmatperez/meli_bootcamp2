package internal

import (
	"context"
	"testing"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestStoreDynamo(t *testing.T) {
	db, err := InitDynamo()
	assert.Nil(t, err)

	repo := NewDynamoRepository(db, "Transactions")
	transactionDynamo := models.TransactionDynamo{
		ID:              "3",
		TransactionCode: "1234-3232",
		Currency:        "$",
		Amount:          25.44,
		Receiver:        "Rivera",
		Sender:          "Soto",
		TransactionDate: "14/05/2020",
	}
	ctx := context.Background()
	err = repo.Store(ctx, &transactionDynamo)
	assert.Nil(t, err)

}

func TestGetDynamo(t *testing.T) {

	db, err := InitDynamo()
	assert.Nil(t, err)
	repo := NewDynamoRepository(db, "Transactions")
	personaId := "2"
	expectedTransactionDynamo := models.TransactionDynamo{
		ID:              "2",
		TransactionCode: "1234-3232",
		Currency:        "$",
		Amount:          25.44,
		Receiver:        "Rivera",
		Sender:          "Soto",
		TransactionDate: "14/05/2020",
	}

	ctx := context.Background()

	transactionCreated, err := repo.GetOne(ctx, personaId)
	assert.Nil(t, err)
	assert.Equal(t, expectedTransactionDynamo.Sender, transactionCreated.Sender)

}

func TestDeleteDynamo(t *testing.T) {

	db, err := InitDynamo()
	assert.Nil(t, err)
	repo := NewDynamoRepository(db, "Transactions")
	transactionId := "2"

	ctx := context.Background()

	err = repo.Delete(ctx, transactionId)
	assert.Nil(t, err)

}
