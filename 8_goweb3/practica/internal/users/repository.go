package internal

import (
	"fmt"
	"time"
	"github.com/extmatperez/meli_bootcamp2/8_goweb3/practica/pkg/store"
)

type Users struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Age          int       `json:"age"`
	Height       float64   `json:"height"`
	Active       bool      `json:"active"`
	CreationDate time.Time `json:"creation_date"`
}

var (
	us []Users
)

type Repository interface {
	GetAll() ([]Users, error)
	Store(id int, firstName string, lastName string, age int) (Users, error)
	Update(id int, firstName string, lastName string, age int) (Users, error)
	UpdateName(id int, firstName string) (Users, error)
	Delete(id int) error
	LastId() (int, error)
}
type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Users, error) {
	err := repo.db.Read(&Users{})
	if err != nil {
		return nil, err
	}
	return Users{}, nil
}

func (repo *repository) Store(id int, fisrtName string, lastName string, email string, age int, height float64,
	active bool, creationDate time.Time) (Users, error) {
	repo.db.Read(&Users{})

	user := Users{
		id,
		fisrtName,
		lastName,
		email,
		age,
		height,
		active,
		creationDate,
	}

	us = append(us, user)
	err := repo.db.Write(us)

	if err != nil {
		return Users{}, err
	}

	return per, nil
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&us)

	if err != nil {
		return 0, err
	}

	if len(us) == 0 {
		return 0, nil
	}

	return us[len(us)-1].ID, nil
}

func (repo *repository) Update(id int, firstName string, lastName string, email string, age int, height float64,
	active bool, creationDate time.Time) (Users, error) {
	err := repo.db.Read(&us)
	if err != nil {
		return Users{}, err
	}

	user := Users{id, firstName, lastName, email, age, height, active, creationDate}
	for i, v := range us {
		if v.ID == id {
			us[i] = user
			err := repo.db.Write(us)
			if err != nil {
				return Users{}, err
			}
			return user, nil
		}
	}
	return Users{}, fmt.Errorf("La persona %d no existe", id)

}
func (repo *repository) UpdateName(id int, name string) (Users, error) {
	for i, v := range us {
		if v.ID == id {
			us[i].FirstName = name
			return us[i], nil
		}
	}
	return Users{}, fmt.Errorf("La persona %d no existe", id)

}

func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&us)
	if err != nil {
		return err
	}

	index := 0
	for i, v := range us {
		if v.ID == id {
			index = i
			us = append(us[:index], us[index+1:]...)
			err := repo.db.Write(us)

			return err
		}
	}
	return fmt.Errorf("La persona %d no existe", id)

}
