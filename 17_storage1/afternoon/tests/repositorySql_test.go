package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"

	users "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/internal/users"
	models "github.com/extmatperez/meli_bootcamp2/tree/archuby_federico/17_storage1/afternoon/pkg/models"
)

func testStoreSql(t *testing.T) {
	repo := users.NewRepositorySql()
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

func TestGetOneSql(t *testing.T) {
	repo := users.NewRepositorySql()

	userObtained, err := repo.GetOne(1)
	assert.Equal(t, 1, userObtained.ID)
	assert.Equal(t, "Archuby", userObtained.LastName)
	assert.Nil(t, err)
}

func TestGetAllSql(t *testing.T) {
	repo := users.NewRepositorySql()

	users, err := repo.GetAll()
	assert.LessOrEqual(t, 6, len(users))
	assert.Nil(t, err)
}
