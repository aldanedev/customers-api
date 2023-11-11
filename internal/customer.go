package customers

import (
	"errors"
	"net/mail"
	"strconv"
)

type CustomerDNI struct {
	value string
}

func NewCustomerDNI(value string) (*CustomerDNI, error) {
	if len(value) != 8 {
		return nil, errors.New("DNI must have 8 digits")
	}
	if _, err := strconv.Atoi(value); err != nil {
		return nil, errors.New("DNI must be a number")
	}

	return &CustomerDNI{value: value}, nil
}

func (dni *CustomerDNI) String() string {
	return dni.value
}

type CustomerFirstName struct {
	value string
}

func NewCustomerFirstName(value string) (*CustomerFirstName, error) {
	if len(value) == 0 {
		return nil, errors.New("First name can't be empty")
	}
	return &CustomerFirstName{value: value}, nil
}

func (firstName *CustomerFirstName) String() string {
	return firstName.value
}

type CustomerLastName struct {
	value string
}

func NewCustomerLastName(value string) (*CustomerLastName, error) {
	if len(value) == 0 {
		return nil, errors.New("Last name can't be empty")
	}
	return &CustomerLastName{value: value}, nil
}

func (lastName *CustomerLastName) String() string {
	return lastName.value
}

type CustomerPhone struct {
	value string
}

func NewCustomerPhone(value string) (*CustomerPhone, error) {
	if len(value) != 9 {
		return nil, errors.New("Phone must have 9 digits")
	}
	if _, err := strconv.Atoi(value); err != nil {
		return nil, errors.New("Phone must be a number")
	}
	return &CustomerPhone{value: value}, nil
}

func (phone *CustomerPhone) String() string {
	return phone.value
}

type CustomerEmail struct {
	value string
}

func NewCustomerEmail(value string) (*CustomerEmail, error) {
	if len(value) == 0 {
		return nil, errors.New("Email can't be empty")
	}

	if _, err := mail.ParseAddress(value); err != nil {
		return nil, errors.New("Email is not valid" + err.Error())
	}

	return &CustomerEmail{value: value}, nil
}

func (email *CustomerEmail) String() string {
	return email.value
}

type Customer struct {
	dni       *CustomerDNI
	firstName *CustomerFirstName
	lastName  *CustomerLastName
	email     *CustomerEmail
	phone     *CustomerPhone
	city      *City
}

func (c *Customer) DNI() *CustomerDNI {
	return c.dni
}

func (c *Customer) FirstName() *CustomerFirstName {
	return c.firstName
}

func (c *Customer) LastName() *CustomerLastName {
	return c.lastName
}

func (c *Customer) Email() *CustomerEmail {
	return c.email
}

func (c *Customer) Phone() *CustomerPhone {
	return c.phone
}

func (c *Customer) City() *City {
	return c.city
}

func (c *Customer) AddCity(city *City) {
	c.city = city
}

func NewCustomer(dni string, firstName string, lastName string, phone string, email string) (*Customer, error) {
	customerDniVo, err := NewCustomerDNI(dni)
	if err != nil {
		return nil, err
	}
	customerFirstNameVo, err := NewCustomerFirstName(firstName)
	if err != nil {
		return nil, err
	}
	customerLastNameVo, err := NewCustomerLastName(lastName)
	if err != nil {
		return nil, err
	}
	customerEmailVo, err := NewCustomerEmail(email)
	if err != nil {
		return nil, err
	}
	customerPhoneVo, err := NewCustomerPhone(phone)
	if err != nil {
		return nil, err
	}
	return &Customer{
		dni:       customerDniVo,
		firstName: customerFirstNameVo,
		lastName:  customerLastNameVo,
		email:     customerEmailVo,
		phone:     customerPhoneVo,
	}, nil
}

type CustomerRepository interface {
	Save(*Customer) error
	Exists(*CustomerDNI) (bool, error)
	FindAll() ([]*Customer, error)
}
