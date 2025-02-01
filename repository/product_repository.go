package repository

import (
	"database/sql"
	"go-rest-api/model"
	"log"
)

type ProductRepository struct {
	connection *sql.DB
}

func NewProductRepository(connection *sql.DB) ProductRepository {
	return ProductRepository{
		connection: connection,
	}
}

func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM products"
	rows, err := pr.connection.Query(query)

	if err != nil {
		log.Fatal(err)
		return []model.Product{}, err
	}

	var productList []model.Product
	var productObj model.Product

	for rows.Next() {
		err = rows.Scan(
			&productObj.ID,
			&productObj.Name,
			&productObj.Price)

		if err != nil {
			log.Fatal(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	err = rows.Close()

	if err != nil {
		log.Fatal(err)
		return productList, err
	}

	return productList, nil
}
