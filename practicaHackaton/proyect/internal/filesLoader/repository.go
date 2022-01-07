package loader

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/pkg/db"
)

type Repository interface {
	CustomersLoader() error
	// InvoicersLoader() error
	// ProductsLoader() error
	// SalesLoader() error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

const (
	insertCustomer = `
		INSERT INTO customers (id, last_name, first_name, conditionn)
		VALUES ( ?, ?, ?, ? )
	`
)

func (r *repository) CustomersLoader() error {

	db := db.StorageDB

	file, err := os.Open("../../../HackthonGo/datos/customers.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		linea := scanner.Text()
		var registro string
		var contador int
		var id int
		var last_name string
		var first_name string
		var condition string

		for i := 0; i < len(linea); i++ {

			registro += string(linea[i])

			if linea[i] == '#' && linea[i+1] == '$' && linea[i+2] == '%' && linea[i+3] == '#' {

				registroDepurado := registro[0 : len(registro)-1]

				if contador == 0 {
					id, err = strconv.Atoi(registroDepurado)
					if err != nil {
						log.Fatal(err)
					}
					registro = ""
					i += 4
				} else if contador == 1 {
					last_name = registroDepurado
					registro = ""
					i += 4
				} else if contador == 2 {
					first_name = registroDepurado
					registro = ""
					i += 4
				}
				contador++
			}
		}
		condition = registro

		stmt, err := db.Prepare(insertCustomer)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		_, err = stmt.Exec(id, last_name, first_name, condition)
		if err != nil {
			return err
		}
	}

	return nil
}
