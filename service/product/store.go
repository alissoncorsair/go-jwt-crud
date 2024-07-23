package product

import (
	"database/sql"

	"github.com/alissoncorsair/goapi/types"
	"github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) GetProductsByID(ids []int) ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * from products WHERE id = ANY($1)", pq.Array(ids))

	if err != nil {
		return nil, err
	}

	products := make([]*types.Product, 0)

	for rows.Next() {
		product, err := ScanRowIntoProduct(rows)

		if err != nil {
			return nil, err
		}

		products = append(products, product)
	}

	return products, nil
}

func (s *Store) GetProductByID(id int) (*types.Product, error) {
	row, err := s.db.Query("SELECT * from products WHERE id = $1", id)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	var product *types.Product

	if row.Next() {
		product, err = ScanRowIntoProduct(row)

		if err != nil {
			return nil, err
		}
	}

	return product, nil
}

func (s *Store) GetProducts() ([]*types.Product, error) {
	rows, err := s.db.Query("SELECT * from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := make([]*types.Product, 0)
	for rows.Next() {
		p, err := ScanRowIntoProduct(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func (s *Store) CreateProduct(product types.CreateProductPayload) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES ($1, $2, $3, $4, $5)", product.Name, product.Description, product.Image, product.Price, product.Quantity)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) UpdateProduct(product types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = $1, price = $2, image = $3, description = $4, quantity = $5 WHERE id = $6", product.Name, product.Price, product.Image, product.Description, product.Quantity, product.ID)

	if err != nil {
		return err
	}

	return nil
}

func ScanRowIntoProduct(rows *sql.Rows) (*types.Product, error) {
	product := &types.Product{}
	err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Image, &product.Price, &product.Quantity, &product.CreatedAt)

	if err != nil {
		return nil, err
	}

	return product, nil
}
