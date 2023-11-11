package bootstrap

import (
	"compartamos/customers/internal/creating"
	"compartamos/customers/internal/plataform/server"
	"compartamos/customers/internal/plataform/storage/postgres"
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
	creatingService := creating.NewCustomerService(customers, cities)

	server := server.New("0.0.0.0", 3000, creatingService)

	return server.Run()
}
