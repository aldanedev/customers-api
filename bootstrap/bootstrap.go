package bootstrap

import (
	"compartamos/customers/internal/creating"
	"compartamos/customers/internal/deleting"
	"compartamos/customers/internal/find"
	"compartamos/customers/internal/list"
	"compartamos/customers/internal/plataform/server"
	"compartamos/customers/internal/plataform/storage/postgres"
	"compartamos/customers/internal/updating"
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Run() error {
	db, err := pgxpool.New(context.Background(), "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable")

	if err != nil {
		return err
	}

	if db.Ping(context.Background()) != nil {
		return errors.New("database is not available")
	}

	customers := postgres.NewCustomerRepository(db)
	cities := postgres.NewCityRepository(db)
	customerCreator := creating.NewCustomerCreator(customers, cities)
	customerLister := list.NewCustomerLister(customers)
	customerUpdater := updating.NewCustomerUpdater(customers, cities)
	cityLister := list.NewCityLister(cities)
	customerDeleter := deleting.NewCustomerDeleter(customers)
	customerFinder := find.NewCustomerFinder(customers)

	server := server.New("0.0.0.0", 3000, customerCreator, customerLister, customerUpdater, cityLister, customerDeleter, customerFinder)

	return server.Run()
}
