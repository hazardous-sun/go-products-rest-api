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

func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	query, err := pr.connection.Prepare(
		"INSERT INTO products (product_name, price) VALUES ($1, $2) RETURNING id")

	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	var id int
	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		log.Fatal(err)
		return -1, err
	}

	err = query.Close()

	if err != nil {
		log.Fatal(err)
		return id, err
	}

	return id, nil
}
