package internal

import (
	"errors"
	"log"
	"strings"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/models"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/db"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/store"
)

type CustomerRepository interface {
	ImportAllCustomers() error
	StoreCustomer(customer models.Customer) (models.Customer, error)
	UpdateCustomer(customer models.Customer) (models.Customer, error)
}

type repository_customer struct {
	arr store.SaveFile
}

func NewCustomerRepository(arr store.SaveFile) CustomerRepository {
	return &repository_customer{arr}
}

func (r *repository_customer) ImportAllCustomers() error {
	customers_string, err := r.arr.ReadLines("/Users/rovega/Documents/GitHub/meli_bootcamp2/hackaton/cmd/server/data/customers.txt")
	if err != nil {
		return err
	}

	for _, customer := range customers_string {
		only_customer := strings.Split(customer, "#$%#")
		id := only_customer[0]
		last_name := only_customer[1]
		first_name := only_customer[2]
		condition := only_customer[3]

		db := db.StorageDB
		query := "INSERT INTO Customer(id,`lastName`, `firstName`, `condition`) VALUES (?,?,?,?)"
		stmt, err := db.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(id, last_name, first_name, condition)

		if err != nil {
			return errors.New("No se pudo guardar elemento en BD.")
		}
	}
	return nil
}

func (r *repository_customer) StoreCustomer(customer models.Customer) (models.Customer, error) {
	db := db.StorageDB
	query := "INSERT INTO Customer(lastName, firstName, condition) VALUES (?,?,?)"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(customer.LastName, customer.FirstName, customer.Condition)
	if err != nil {
		return models.Customer{}, err
	}

	idCreado, _ := result.LastInsertId()
	customer.Id = int(idCreado)
	return customer, nil
}

func (r *repository_customer) UpdateCustomer(customer models.Customer) (models.Customer, error) {
	db := db.StorageDB
	query := "UPDATE Customer SET lastName = ?, firstName = ?, condition = ? WHERE id = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	result, err := stmt.Exec(customer.LastName, customer.FirstName, customer.Condition, customer.Id)
	if err != nil {
		return models.Customer{}, err
	}
	updatedRows, _ := result.RowsAffected()
	if updatedRows == 0 {
		return models.Customer{}, errors.New("No se encontro al customer.")
	}
	return customer, nil
}
