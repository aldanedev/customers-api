package handler

import (
	"compartamos/customers/internal/list"

	"github.com/gofiber/fiber/v2"
)

type cityResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type customerResponse struct {
	DNI       string       `json:"dni"`
	FirstName string       `json:"first_name"`
	LastName  string       `json:"last_name"`
	Phone     string       `json:"phone"`
	Email     string       `json:"email"`
	City      cityResponse `json:"city"`
}

func ListHandler(customerLister list.CustomerLister) fiber.Handler {
	return func(c *fiber.Ctx) error {

		customers, err := customerLister.ListCustomers()
		if err != nil {
			return err
		}

		var response []customerResponse
		for _, customer := range customers {

			city := cityResponse{
				ID:   customer.City().ID().String(),
				Name: customer.City().Name().String(),
			}

			response = append(response, customerResponse{
				DNI:       customer.DNI().String(),
				FirstName: customer.FirstName().String(),
				LastName:  customer.LastName().String(),
				Phone:     customer.Phone().String(),
				Email:     customer.Email().String(),
				City:      city,
			})
		}

		return c.JSON(response)
	}
}

type citiesResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func ListCitiesHandler(cityLister list.CityLister) fiber.Handler {
	return func(c *fiber.Ctx) error {
		cities, err := cityLister.ListCities()

		if err != nil {
			return c.JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		var response []citiesResponse
		for _, city := range cities {
			response = append(response, citiesResponse{
				ID:   city.ID().String(),
				Name: city.Name().String(),
			})
		}

		return c.JSON(response)
	}
}
