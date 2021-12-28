package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/internal/users"
	models "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/pkg/models"

	basicDb "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/pkg/database"
)

func testStoreSql(t *testing.T) {
	db := basicDb.StorageDB
	repo := users.NewRepositorySql(db)
	user := models.User{
		Name:     "Federico",
		LastName: "Archuby",
		Email:    "fede@hola.com",
		Age:      32,
		Height:   1.72,
		Active:   true,
		Created:  "2021-10-21",
	}

	userCreated, err := repo.Store(user)
	assert.NotEqual(t, 0, userCreated.ID)
	assert.Nil(t, err)
}

func testGetOneSql(t *testing.T) {
	db := basicDb.StorageDB
	repo := users.NewRepositorySql(db)

	userObtained, err := repo.GetOne(1)
	assert.Equal(t, 1, userObtained.ID)
	assert.Equal(t, "Archuby", userObtained.LastName)
	assert.Nil(t, err)
}

func testGetAllSql(t *testing.T) {
	db := basicDb.StorageDB
	repo := users.NewRepositorySql(db)

	users, err := repo.GetAll()
	assert.LessOrEqual(t, 6, len(users))
	assert.Nil(t, err)
}

func TestStoreGetOneTrx(t *testing.T) {
	db, err := basicDb.InitDb()
	assert.NoError(t, err)
	repo := users.NewRepositorySql(db)

	user := models.User{
		Name:     "Federico",
		LastName: "Archuby",
		Email:    "fede@hola.com",
		Age:      32,
		Height:   1.72,
		Active:   true,
		Created:  "2021-10-21",
	}
	userCreated, err := repo.Store(user)
	assert.NoError(t, err)

	var userObtained models.User
	userObtained, err = repo.GetOne(userCreated.ID)

	assert.NoError(t, err)
	assert.Equal(t, userObtained.Name, userCreated.Name)
	assert.Equal(t, userObtained.ID, userCreated.ID)
}

func TestUpdateDeleteTrx(t *testing.T) {
	db, err := basicDb.InitDb()
	assert.NoError(t, err)
	repo := users.NewRepositorySql(db)

	userInitial := models.User{
		ID:       1,
		Name:     "Juan Carlos",
		LastName: "Perez",
		Email:    "fede@hola.com",
		Age:      32,
		Height:   1.72,
		Active:   true,
		Created:  "2021-10-21",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	var userObtained, userEdited models.User
	userObtained, err = repo.GetOne(userInitial.ID)

	userEdited, err = repo.Update(userInitial, ctx)

	userObtained, err = repo.GetOne(userInitial.ID)

	assert.NoError(t, err)
	assert.Equal(t, userObtained.Name, userEdited.Name)
	assert.Equal(t, userObtained.ID, userEdited.ID)

	var deleted bool
	deleted, err = repo.Delete(userInitial.ID)
	assert.NoError(t, err)
	assert.Equal(t, true, deleted)
}
