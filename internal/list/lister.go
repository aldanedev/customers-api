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
