package updating

import customers "compartamos/customers/internal"

type CustomerUpdater struct {
	customerRepository customers.CustomerRepository
	cityRepository     customers.CityRepository
}

func NewCustomerUpdater(customerRepository customers.CustomerRepository, cityRepository customers.CityRepository) CustomerUpdater {
	return CustomerUpdater{customerRepository: customerRepository, cityRepository: cityRepository}
}

func (cu CustomerUpdater) UpdateCustomer(dni string, firstName string, lastName string, phone string, email string, cityID string) error {
	dniVo, err := customers.NewCustomerDNI(dni)
	if err != nil {
		return err
	}

	customer, err := cu.customerRepository.FindByDNI(dniVo)
	if err != nil {
		return err
	}

	cityIDVo, err := customers.NewCityID(cityID)

	if err != nil {
		return err
	}

	city, err := cu.cityRepository.FindByID(*cityIDVo)
	if err != nil {
		return err
	}

	err = customer.UpdateCustomer(firstName, lastName, phone, email, city)

	if err != nil {
		return err
	}

	return cu.customerRepository.Save(customer)
}
