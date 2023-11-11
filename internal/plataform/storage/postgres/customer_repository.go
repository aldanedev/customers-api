package postgres

import (
	customers "compartamos/customers/internal"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

var sqlInsertCustomer = `
  INSERT INTO customers (dni, first_name, last_name, phone, email, city_id)
  VALUES ($1, $2, $3, $4, $5, $6)
`

var sqlExistsCustomer = `
  SELECT EXISTS(SELECT 1 FROM customers WHERE dni = $1)
`

type CustomerRepository struct {
	db *pgxpool.Pool
}

func NewCustomerRepository(db *pgxpool.Pool) customers.CustomerRepository {
	return &CustomerRepository{db: db}
}

func (r *CustomerRepository) Save(customer *customers.Customer) error {
	_, err := r.db.Exec(context.Background(),
		sqlInsertCustomer,
		customer.DNI().String(),
		customer.FirstName().String(),
		customer.LastName().String(),
		customer.Phone().String(),
		customer.Email().String(),
		customer.City().ID().String(),
	)

	if err != nil {
		return err
	}

	return nil
}

func (r *CustomerRepository) Exists(dni *customers.CustomerDNI) (bool, error) {
	var exists bool
	err := r.db.QueryRow(context.Background(), sqlExistsCustomer, dni.String()).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}
