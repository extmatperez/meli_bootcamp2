package internal

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/18_storage2/afternoon_activities/Exercise_1_2_3/internal/models"
)

type Service_sql interface {
	Store(first_name, last_name, email string, age, height int, active bool, date string) (models.Users, error)
	Get_one_user(id int) models.Users
	Get_by_name(name string) ([]models.Users, error)
	Get_all_users() ([]models.Users, error)
	Get_full_data() ([]models.Users, error)
	Get_one_with_context(ctx context.Context, id int) (models.Users, error)
	Update_user(users models.Users) (models.Users, error)
	Delete_user(id int) error
}

type service_sql struct {
	repository Repository_sql
}

func New_service_sql(repo Repository_sql) Service_sql {
	return &service_sql{repository: repo}
}

func (ser *service_sql) Store(first_name, last_name, email string, age, height int, active bool, date string) (models.Users, error) {
	new_users := models.Users{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
		Age:       age,
		Height:    height,
		Active:    active,
		Date:      date,
	}
	user_created, err := ser.repository.Store(new_users)
	if err != nil {
		return models.Users{}, err
	}
	return user_created, nil
}

func (ser *service_sql) Get_one_user(id int) models.Users {
	return ser.repository.Get_one_user(id)
}

func (ser *service_sql) Get_by_name(name string) ([]models.Users, error) {
	return ser.repository.Get_by_name(name)
}

func (ser *service_sql) Get_all_users() ([]models.Users, error) {
	return ser.repository.Get_all_users()
}

func (ser *service_sql) Get_full_data() ([]models.Users, error) {
	return ser.repository.Get_full_data()
}

func (ser *service_sql) Get_one_with_context(ctx context.Context, id int) (models.Users, error) {
	return ser.repository.Get_one_with_context(ctx, id)
}

func (ser *service_sql) Update_user(users models.Users) (models.Users, error) {
	return ser.repository.Update_user(users)
}

func (ser *service_sql) Delete_user(id int) error {
	return ser.repository.Delete_user(id)
}
