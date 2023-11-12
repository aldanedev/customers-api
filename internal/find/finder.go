package find

import customers "compartamos/customers/internal"

type CustomerFinder struct {
	customers customers.CustomerRepository
}

func NewCustomerFinder(customers customers.CustomerRepository) CustomerFinder {
	return CustomerFinder{customers: customers}
}

func (f *CustomerFinder) FindCustomer(dni string) (*customers.Customer, error) {
	dniVo, err := customers.NewCustomerDNI(dni)
	if err != nil {
		return nil, err
	}

	return f.customers.FindByDNI(dniVo)
}
