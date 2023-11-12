package handler

import (
	"compartamos/customers/internal/find"

	"github.com/gofiber/fiber/v2"
)

func FindHandler(customerFinder find.CustomerFinder) fiber.Handler {
	return func(c *fiber.Ctx) error {
		dni := c.Params("dni")
		customer, err := customerFinder.FindCustomer(dni)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
		city := cityResponse{
			ID:   customer.City().ID().String(),
			Name: customer.City().Name().String(),
		}
		response := customerResponse{
			DNI:       customer.DNI().String(),
			FirstName: customer.FirstName().String(),
			LastName:  customer.LastName().String(),
			Phone:     customer.Phone().String(),
			Email:     customer.Email().String(),
			City:      city,
		}
		return c.JSON(response)
	}
}
