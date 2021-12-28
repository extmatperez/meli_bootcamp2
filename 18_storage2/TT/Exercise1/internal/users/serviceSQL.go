package internal

import (
	"context"

	"github.com/extmatperez/meli_bootcamp2/18_storage2/TT/Exercise1/internal/models"
)

type ServiceSQL interface {
	Store(first_name, last_name, email string, age, height int, active bool, cration_date string) (models.User, error)
	GetOne(id int) models.User
	Update(user models.User) (models.User, error)
	GetAll() ([]models.User, error)
	Delete(id int) error
	GetOneWithContext(ctx context.Context, id int) (models.User, error)
	GetFullData() ([]models.User, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (ser *serviceSQL) Store(first_name, last_name, email string, age, height int, active bool, cration_date string) (models.User, error) {
	newUsers := models.User{FirstName: first_name, LastName: last_name, Email: email, Age: age, Height: height, Active: active, CrationDate: cration_date}
	userCreated, err := ser.repository.Store(newUsers)

	if err != nil {
		return models.User{}, err
	}
	return userCreated, nil
}

func (ser *serviceSQL) GetOne(id int) models.User {
	return ser.repository.GetOne(id)
}

func (ser *serviceSQL) Update(user models.User) (models.User, error) {
	return ser.repository.Update(user)
}

func (ser *serviceSQL) GetAll() ([]models.User, error) {
	return ser.repository.GetAll()
}

func (ser *serviceSQL) Delete(id int) error {
	return ser.repository.Delete(id)
}

func (ser *serviceSQL) GetOneWithContext(ctx context.Context, id int) (models.User, error) {
	return ser.repository.GetOneWithContext(ctx, id)
}

func (ser *serviceSQL) GetFullData() ([]models.User, error) {
	return ser.repository.GetFullData()
}
