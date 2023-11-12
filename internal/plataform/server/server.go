package server

import (
	"compartamos/customers/internal/creating"
	"compartamos/customers/internal/deleting"
	"compartamos/customers/internal/find"
	"compartamos/customers/internal/list"
	"compartamos/customers/internal/plataform/server/handler"
	"compartamos/customers/internal/updating"

	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	httpAddr        string
	engine          *fiber.App
	customerCreator creating.CustomerCreator
	customerLister  list.CustomerLister
	customerUpdater updating.CustomerUpdater
	cityLister      list.CityLister
	customerDeleter deleting.CustomerDeleter
	customerFinder  find.CustomerFinder
}

func New(
	host string,
	port uint,
	customerCreator creating.CustomerCreator,
	customerLister list.CustomerLister,
	customerUpdater updating.CustomerUpdater,
	cityLister list.CityLister,
	customerDeleter deleting.CustomerDeleter,
	customerFinder find.CustomerFinder,
) Server {

	srv := Server{
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		engine:          fiber.New(),
		customerCreator: customerCreator,
		customerLister:  customerLister,
		customerUpdater: customerUpdater,
		cityLister:      cityLister,
		customerDeleter: customerDeleter,
		customerFinder:  customerFinder,
	}

	srv.RegisterRoutes()

	return srv
}

func (s *Server) Run() error {
	log.Printf("Server running on %s", s.httpAddr)

	s.engine.Use(recover.New())

	return s.engine.Listen(s.httpAddr)
}

func (s *Server) RegisterRoutes() {
	s.engine.Post("/customers", handler.CreateHandler(s.customerCreator))
	s.engine.Get("/customers", handler.ListHandler(s.customerLister))
	s.engine.Put("/customers/:dni", handler.UpdateHandler(s.customerUpdater))
	s.engine.Get("/cities", handler.ListCitiesHandler(s.cityLister))
	s.engine.Delete("/customers/:id", handler.DeleteHandler(s.customerDeleter))
	s.engine.Get("/customers/:dni", handler.FindHandler(s.customerFinder))
}
