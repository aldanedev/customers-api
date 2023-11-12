package postgres

import (
	customers "compartamos/customers/internal"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var sqlSelectCustomers = `
  SELECT c.dni, c.first_name, c.last_name, c.phone, c.email, ci.id, ci.name
  FROM customers c
  INNER JOIN cities ci ON ci.id = c.city_id
`

var sqlInsertCustomer = `
  INSERT INTO customers (dni, first_name, last_name, phone, email, city_id)
  VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (dni)
  DO UPDATE SET first_name = $2, last_name = $3, phone = $4, email = $5, city_id = $6
`

var sqlExistsCustomer = `
  SELECT EXISTS(SELECT 1 FROM customers WHERE dni = $1)
`

var sqlDeleteCustomer = `
  DELETE FROM customers WHERE dni = $1
`

var sqlSelectCutomerByDNI = sqlSelectCustomers + " WHERE c.dni = $1"

type CustomerRepository struct {
	db *pgxpool.Pool
}

func (r *CustomerRepository) Delete(dni *customers.CustomerDNI) error {
	_, err := r.db.Exec(context.Background(), sqlDeleteCustomer, dni.String())

	return err
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

func (r *CustomerRepository) FindAll() ([]*customers.Customer, error) {
	rows, err := r.db.Query(context.Background(), sqlSelectCustomers)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var customersList []*customers.Customer
	for rows.Next() {
		customer, err := r.scanRow(rows)
		if err != nil {
			return nil, err
		}
		customersList = append(customersList, customer)
	}
	return customersList, nil
}

func (r *CustomerRepository) FindByDNI(dni *customers.CustomerDNI) (*customers.Customer, error) {
	rows, err := r.db.Query(context.Background(), sqlSelectCutomerByDNI, dni.String())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if !rows.Next() {
		return nil, errors.New("customer not found")
	}

	return r.scanRow(rows)
}

func (r *CustomerRepository) scanRow(rows pgx.Rows) (*customers.Customer, error) {
	var dni, firstName, lastName, phone, email, cityID, cityName string
	err := rows.Scan(&dni, &firstName, &lastName, &phone, &email, &cityID, &cityName)
	if err != nil {
		return nil, err
	}

	customer, err := customers.NewCustomer(dni, firstName, lastName, phone, email)
	if err != nil {
		return nil, err
	}
	city, err := customers.NewCity(cityID, cityName)
	if err != nil {
		return nil, err
	}

	customer.AddCity(city)

	return customer, nil
}
