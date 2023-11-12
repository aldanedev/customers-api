package list

import customers "compartamos/customers/internal"

type CustomerLister struct {
	customerRepository customers.CustomerRepository
}

func NewCustomerLister(customerRepository customers.CustomerRepository) CustomerLister {
	return CustomerLister{customerRepository: customerRepository}
}

func (cl CustomerLister) ListCustomers() ([]*customers.Customer, error) {
	return cl.customerRepository.FindAll()
}

type CityLister struct {
	cityRepository customers.CityRepository
}

func NewCityLister(cityRepository customers.CityRepository) CityLister {
	return CityLister{cityRepository: cityRepository}
}

func (cl CityLister) ListCities() ([]*customers.City, error) {
	return cl.cityRepository.FindAll()
}
