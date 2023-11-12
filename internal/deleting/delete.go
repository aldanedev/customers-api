package deleting

import customers "compartamos/customers/internal"

type CustomerDeleter struct {
	customerRepository customers.CustomerRepository
}

func NewCustomerDeleter(customerRepository customers.CustomerRepository) CustomerDeleter {
	return CustomerDeleter{customerRepository: customerRepository}
}

func (cd CustomerDeleter) DeleteCustomer(id string) error {
	dniVo, err := customers.NewCustomerDNI(id)
	if err != nil {
		return err
	}

	return cd.customerRepository.Delete(dniVo)
}
