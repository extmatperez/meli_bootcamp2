package internal

import (
	"testing"

	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/18_storage2/morning_activities/Exercise_1_2_3/internal/models"
	"github.com/stretchr/testify/assert"
)

func Test_store_service_sql(t *testing.T) {
	new_users := models.Users{
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "email",
		Age:       33,
		Height:    98,
		Active:    true,
		Date:      "27/12/2021",
	}
	repo := New_repository_sql()
	service := New_service_sql(repo)

	user_created, _ := service.Store(
		new_users.FirstName,
		new_users.LastName,
		new_users.Email,
		new_users.Age,
		new_users.Height,
		new_users.Active,
		new_users.Date,
	)
	assert.Equal(t, new_users.FirstName, user_created.FirstName)
	assert.Equal(t, new_users.LastName, user_created.LastName)

}

func Test_Get_one_user_service_sql(t *testing.T) {
	new_users := models.Users{
		FirstName: "first_name",
		LastName:  "last_name",
		Email:     "email",
		Age:       33,
		Height:    98,
		Active:    true,
		Date:      "27/12/2021",
	}

	repo := New_repository_sql()
	service := New_service_sql(repo)

	user_loaded := service.Get_one_user(1)

	assert.Equal(t, new_users.FirstName, user_loaded.FirstName)
	assert.Equal(t, new_users.LastName, user_loaded.LastName)
}

func Test_Get_by_name_service_sql(t *testing.T) {
	new_users := "Viviana"
	/* models.Users{
		FirstName: "Viviana",
		LastName:  "last_name",
		Email:     "email",
		Age:       33,
		Height:    98,
		Active:    true,
		Date:      "27/12/2021",
	} */

	repo := New_repository_sql()
	service := New_service_sql(repo)

	user_loaded, err := service.Get_by_name(new_users)
	assert.Nil(t, err)
	assert.True(t, len(user_loaded) > 0)

	/* assert.Equal(t, new_users.FirstName, user_loaded) */
}

func Test_get_all_users_service_sql(t *testing.T) {
	repo := New_repository_sql()
	service := New_service_sql(repo)

	all_users_db, err := service.Get_all_users()

	assert.Nil(t, err)
	assert.True(t, len(all_users_db) >= 0)
}

func Test_update_service_sql_ok(t *testing.T) {
	user_updated := models.Users{
		ID:        5,
		FirstName: "Viviana",
		LastName:  "Valera",
		Age:       27,
	}

	repo := New_repository_sql()
	service := New_service_sql(repo)
	// Esto lo hacemos con el fin de que al actualizar no me cree problemas en el siguiente test manteniendo los datos (1)
	last_user := service.Get_one_user(user_updated.ID)

	user_loaded, _ := service.Update(user_updated)

	assert.Equal(t, user_updated.FirstName, user_loaded.FirstName)
	assert.Equal(t, user_updated.LastName, user_loaded.LastName)

	// Esto lo hacemos con el fin de que al actualizar no me cree problemas en el siguiente test manteniendo los datos (2)
	_, err := service.Update(last_user)
	assert.Nil(t, err)
}

func Test_update_failed(t *testing.T) {
	user_updated := models.Users{
		ID:        20,
		FirstName: "Viviana",
		LastName:  "Valera",
		Age:       27,
	}

	repo := New_repository_sql()
	service := New_service_sql(repo)

	_, err := service.Update(user_updated)

	assert.Equal(t, "user not found", err.Error())
}
