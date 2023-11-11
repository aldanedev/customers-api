package creating

import (
	customers "compartamos/customers/internal"
	"errors"
)

type CustomerService struct {
	customerRepository customers.CustomerRepository
	cityRepository     customers.CityRepository
}

func NewCustomerService(customerRepository customers.CustomerRepository, cityRepository customers.CityRepository) CustomerService {
	return CustomerService{customerRepository: customerRepository, cityRepository: cityRepository}
}

func (cs CustomerService) CreateCustomer(dni, firstName, lastName, phone, email, cityID string) error {

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
