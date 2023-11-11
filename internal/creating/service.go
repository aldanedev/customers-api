package creating

import (
	customers "compartamos/customers/internal"
	"errors"
)

type CustomerCreator struct {
	customerRepository customers.CustomerRepository
	cityRepository     customers.CityRepository
}

func NewCustomerCreator(customerRepository customers.CustomerRepository, cityRepository customers.CityRepository) CustomerCreator {
	return CustomerCreator{customerRepository: customerRepository, cityRepository: cityRepository}
}

func (cs CustomerCreator) CreateCustomer(dni, firstName, lastName, phone, email, cityID string) error {

	dniVo, err := customers.NewCustomerDNI(dni)
	if err != nil {
		return err
	}

	exists, err := cs.customerRepository.Exists(dniVo)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("Customer already exists")
	}

	cityIDVo, err := customers.NewCityID(cityID)
	if err != nil {
		return err
	}

	city, err := cs.cityRepository.FindByID(*cityIDVo)
	if err != nil {
		return err
	}

	customer, err := customers.NewCustomer(dni, firstName, lastName, phone, email)
	if err != nil {
		return err
	}

	customer.AddCity(city)

	return cs.customerRepository.Save(customer)
}
