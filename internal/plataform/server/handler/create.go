package handler

import (
	"compartamos/customers/internal/creating"

	"github.com/gofiber/fiber/v2"
)

type createRequest struct {
	DNI       string `json:"dni"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	CityID    string `json:"city_id"`
}

func CreateHandler(customerCreator creating.CustomerCreator) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var request createRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		err := customerCreator.CreateCustomer(
			request.DNI,
			request.FirstName,
			request.LastName,
			request.Phone,
			request.Email,
			request.CityID,
		)

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Customer created successfully",
		})

	}
}
