package repository

import (
	"database/sql"
	"errors"
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

// Create --------------------------------------------------------------------------------------------------------------

// CreateProduct :
// Creates a new product inside the database based on the model received as parameter.
func (pr *ProductRepository) CreateProduct(product model.Product) (int, error) {
	query, err := pr.connection.Prepare(
		"INSERT INTO products (product_name, price) VALUES ($1, $2) RETURNING id")

	if err != nil {
		log.Fatalln(err)
		return -1, err
	}

	var id int
	err = query.QueryRow(product.Name, product.Price).Scan(&id)

	if err != nil {
		log.Fatalln(err)
		return -1, err
	}

	err = query.Close()

	if err != nil {
		log.Fatalln(err)
		return id, err
	}

	return id, nil
}

// Read ----------------------------------------------------------------------------------------------------------------

// GetProducts :
// Returns all products inside the database.
func (pr *ProductRepository) GetProducts() ([]model.Product, error) {
	query := "SELECT * FROM products"
	rows, err := pr.connection.Query(query)

	if err != nil {
		log.Fatalln(err)
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
			log.Fatalln(err)
			return []model.Product{}, err
		}

		productList = append(productList, productObj)
	}

	err = rows.Close()

	if err != nil {
		log.Fatalln(err)
		return productList, err
	}

	return productList, nil
}

// GetProductByID :
// Returns by id a product inside the database.
func (pr *ProductRepository) GetProductByID(id int) (*model.Product, error) {
	query, err := pr.connection.Prepare("SELECT * FROM products WHERE id = $1")

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	var productObj model.Product
	err = query.QueryRow(id).Scan(&productObj.ID, &productObj.Name, &productObj.Price)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		} else {
			return nil, err
		}
	}

	return &productObj, nil
}

// Update --------------------------------------------------------------------------------------------------------------

// UpdateProductByID :
// Updates a product inside the database that has the same id as the product parameter
func (pr *ProductRepository) UpdateProductByID(product model.Product) error {
	_, err := pr.GetProductByID(product.ID)

	if err != nil {
		return err
	}

	query, err := pr.connection.Prepare("UPDATE products SET product_name = $1, price = $2 WHERE id = $3")

	if err != nil {
		return err
	}

	err = query.QueryRow(product.Name, product.Price, product.ID).Scan(nil, nil, nil)

	if err != nil {
		return err
	}

	return nil
}
