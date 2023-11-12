package handler

import (
	"compartamos/customers/internal/updating"

	"github.com/gofiber/fiber/v2"
)

type updateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CityID    string `json:"city_id"`
}

func UpdateHandler(cutomerUpdate updating.CustomerUpdater) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var request updateRequest
		if err := c.BodyParser(&request); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		dni := c.Params("dni")

		err := cutomerUpdate.UpdateCustomer(
			dni,
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
			"message": "Customer updated successfully",
		})
	}
}
