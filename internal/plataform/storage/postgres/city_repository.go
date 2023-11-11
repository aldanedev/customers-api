package postgres

import (
	customers "compartamos/customers/internal"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var sqlSelectCities = `
  SELECT id, name FROM cities
`

var sqlSelectCityByID = sqlSelectCities + `
  WHERE id = $1
`

type CityRepository struct {
	db *pgxpool.Pool
}

func NewCityRepository(db *pgxpool.Pool) customers.CityRepository {
	return &CityRepository{db: db}
}

func (r *CityRepository) FindByID(id customers.CityID) (*customers.City, error) {
	row, err := r.db.Query(context.Background(), sqlSelectCityByID, id.String())
	if err != nil {
		return nil, err
	}
	defer row.Close()
	row.Next()

	return r.scanRow(row)
}

func (r *CityRepository) FindAll() ([]*customers.City, error) {
	rows, err := r.db.Query(context.Background(), sqlSelectCities)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var cities []*customers.City
	for rows.Next() {
		city, err := r.scanRow(rows)
		if err != nil {
			return nil, err
		}
		cities = append(cities, city)
	}

	return cities, nil
}

func (r *CityRepository) scanRow(row pgx.Rows) (*customers.City, error) {
	var id, name string
	err := row.Scan(&id, &name)
	if err != nil {
		return nil, err
	}

	return customers.NewCity(id, name)
}
