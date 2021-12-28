package internal

import "github.com/extmatperez/meli_bootcamp2/tree/montenegro_edgar/18_storage2/morning_activities/Exercise_1_2_3/internal/models"

type Service_sql interface {
	Store(first_name, last_name, email string, age, height int, active bool, date string) (models.Users, error)
	Get_one_user(id int) models.Users
	Get_by_name(name string) ([]models.Users, error)
	Update(users models.Users) (models.Users, error)
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

func (ser *service_sql) Update(users models.Users) (models.Users, error) {
	return ser.repository.Update(users)
}
