package product

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/willy-r/ecom-example/types"
	"github.com/willy-r/ecom-example/utils"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetProducts() ([]types.Product, error) {
	rows, err := s.db.Query("SELECT * FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make([]types.Product, 0)
	for rows.Next() {
		p, err := utils.ScanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) CreateProduct(p types.Product) error {
	_, err := s.db.Exec("INSERT INTO products (name, description, image, price, quantity) VALUES (?, ?, ?, ?, ?)", p.Name, p.Description, p.Image, p.Price, p.Quantity)
	if err != nil {
		return err
	}
	return nil
}

func (s *Store) GetProductsByIds(ids []int) ([]types.Product, error) {
	placeHolders := strings.Repeat(",?", len(ids)-1)
	query := fmt.Sprintf("SELECT * FROM products WHERE id IN (?%s)", placeHolders)

	args := make([]interface{}, len(ids))
	for i, id := range ids {
		args[i] = id
	}

	rows, err := s.db.Query(query, args...)
	if err != nil {
		return nil, err
	}

	products := []types.Product{}
	for rows.Next() {
		p, err := utils.ScanRowsIntoProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, *p)
	}

	return products, nil
}

func (s *Store) UpdateProduct(p types.Product) error {
	_, err := s.db.Exec("UPDATE products SET name = ?, description = ?, image = ?, price = ?, quantity = ? WHERE id = ?", p.Name, p.Description, p.Image, p.Price, p.Quantity, p.ID)
	if err != nil {
		return err
	}
	return nil
}
