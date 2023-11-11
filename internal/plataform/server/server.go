package server

import (
	"compartamos/customers/internal/creating"
	"compartamos/customers/internal/plataform/server/handler"

	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Server struct {
	httpAddr        string
	engine          *fiber.App
	creatingService creating.CustomerService
}

func New(host string, port uint, creatingService creating.CustomerService) Server {
	srv := Server{
		httpAddr:        fmt.Sprintf("%s:%d", host, port),
		engine:          fiber.New(),
		creatingService: creatingService,
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
	s.engine.Post("/customers", handler.CreateHandler(s.creatingService))
}
