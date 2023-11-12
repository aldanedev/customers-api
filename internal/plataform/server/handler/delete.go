package handler

import (
	"compartamos/customers/internal/deleting"

	"github.com/gofiber/fiber/v2"
)

func DeleteHandler(customerDeleter deleting.CustomerDeleter) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		err := customerDeleter.DeleteCustomer(id)
		if err != nil {
			return err
		}
		return c.SendStatus(fiber.StatusNoContent)
	}
}
