package internal

import "github.com/extmatperez/meli_bootcamp2/17_storage1/TT/Exercise1/internal/models"

type ServiceSQL interface {
	Store(first_name, last_name, email string, age, height int, active bool, cration_date string) (models.User, error)
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
