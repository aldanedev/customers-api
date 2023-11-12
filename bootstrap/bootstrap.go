package bootstrap

import (
	"compartamos/customers/internal/creating"
	"compartamos/customers/internal/list"
	"compartamos/customers/internal/plataform/server"
	"compartamos/customers/internal/plataform/storage/postgres"
	"compartamos/customers/internal/updating"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Run() error {
	db, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		return err
	}

	customers := postgres.NewCustomerRepository(db)
	cities := postgres.NewCityRepository(db)
	customerCreator := creating.NewCustomerCreator(customers, cities)
	customerLister := list.NewCustomerLister(customers)
	customerUpdater := updating.NewCustomerUpdater(customers, cities)
	cityLister := list.NewCityLister(cities)

	server := server.New("0.0.0.0", 3000, customerCreator, customerLister, customerUpdater, cityLister)

	return server.Run()
}
