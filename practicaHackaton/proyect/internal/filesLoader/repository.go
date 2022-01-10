package loader

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/extmatperez/meli_bootcamp2/practicaHackaton/proyect/pkg/db"
)

type Repository interface {
	CustomersLoader() error
	InvoicersLoader() error
	ProductsLoader() error
	SalesLoader() error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

var (
	insertCustomer = `
		INSERT INTO customers (id, last_name, first_name, conditionn)
		VALUES 
	`
	insertInvoicer = `
		INSERT INTO invoicers (id, date_time, id_customer, total)
		VALUES
	`
	insertProduct = `
		INSERT INTO products (id, description, price)
		VALUES
	`
	insertSale = `
		INSERT INTO sales (id, id_invoice, id_product, quantity)
		VALUES
	`
)

func loader(path, insert, paramsCantity string) error {

	db := db.StorageDB

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var datafinal []interface{}

	for scanner.Scan() {
		insert += paramsCantity
		line := scanner.Text()
		data := strings.Split(line, "#$%#")

		for i := 0; i < len(data); i++ {
			datafinal = append(datafinal, data[i])
		}

		if line[len(line)-4:] == "#$%#" {
			datafinal[len(datafinal)-1] = 0
		}
	}
	insert = insert[:len(insert)-1]
	stmt, err := db.Prepare(insert)
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(datafinal...)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) CustomersLoader() error {
	return loader("../../../HackthonGo/datos/customers.txt", insertCustomer, "(?, ?, ?, ?),")
}

func (r *repository) InvoicersLoader() error {
	return loader("../../../HackthonGo/datos/invoices.txt", insertInvoicer, "(?, ?, ?, ?),")
}

func (r *repository) ProductsLoader() error {
	return loader("../../../HackthonGo/datos/products.txt", insertProduct, "(?, ?, ?),")
}

func (r *repository) SalesLoader() error {
	return loader("../../../HackthonGo/datos/sales.txt", insertSale, "(?, ?, ?, ?),")
}
