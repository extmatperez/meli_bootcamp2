package main

import (
	"fmt"
	"log"

	customers "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/customers"
	invoices "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/invoices"
	products "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/products"
	sales "github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/internal/sales"

	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/db"
	"github.com/extmatperez/meli_bootcamp2/tree/vega_rodrigo/hackaton/pkg/store"
)

func main() {
	arr_costumer := store.NewSave(store.FileTypeSave)
	db_customer := db.StorageDB
	repository_costumer := customers.NewCustomerRepository(arr_costumer, db_customer)
	service_costumer := customers.NewCustomerService(repository_costumer)

	err := service_costumer.ImportAllCustomers()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se pasaron los customers, chequear DB.")

	arr_invoices := store.NewSave(store.FileTypeSave)
	repository_invoice := invoices.NewInvoiceRepository(arr_invoices)
	service_invoice := invoices.NewInvoiceService(repository_invoice)

	err = service_invoice.ImportAllInvoices()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se pasaron los invoices, chequear DB.")

	arr_products := store.NewSave(store.FileTypeSave)
	repository_product := products.NewProductRepository(arr_products)
	service_product := products.NewProductService(repository_product)

	err = service_product.ImportAllProducts()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se pasaron los products, chequear DB.")

	arr_sales := store.NewSave(store.FileTypeSave)
	repository_sale := sales.NewSaleRepository(arr_sales)
	service_sale := sales.NewSaleService(repository_sale)

	err = service_sale.ImportAllSales()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se pasaron los sales, chequear DB.")

	fmt.Println("Ahora se modificaran los totales de las facturas.")
	err = service_invoice.UpdateTotalsOfInvoices()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Se modificaron los totales de las facturas, chequear DB.")
}
