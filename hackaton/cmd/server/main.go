package main

import (
	"fmt"
	"log"

	customers "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/customers"
	invoices "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/invoices"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/store"
)

func main() {
	arr_costumer := store.NewSave(store.FileTypeSave, "/Users/rovega/Documents/GitHub/meli_bootcamp2/hackaton/cmd/server/data/customers.txt")
	repository_costumer := customers.NewCustomerRepository(arr_costumer)
	service_costumer := customers.NewCustomerService(repository_costumer)

	err := service_costumer.ImportAllCustomers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se pasaron los customers, chequear DB.")

	arr_invoices := store.NewSave(store.FileTypeSave, "/Users/rovega/Documents/GitHub/meli_bootcamp2/hackaton/cmd/server/data/invoices.txt")
	repository_invoice := invoices.NewInvoiceRepository(arr_invoices)
	service_invoice := invoices.NewInvoiceService(repository_invoice)

	err = service_invoice.ImportAllInvoices()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se pasaron los invoices, chequear DB.")
}
