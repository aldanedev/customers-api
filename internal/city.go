package customers

import "errors"

type CityID struct {
	value string
}

func NewCityID(value string) (*CityID, error) {
	if len(value) == 0 {
		return nil, errors.New("City ID can't be empty")
	}
	return &CityID{value: value}, nil
}

func (id *CityID) String() string {
	return id.value
}

type CityName struct {
	value string
}

func NewCityName(value string) (*CityName, error) {
	if len(value) == 0 {
		return nil, errors.New("City name can't be empty")
	}
	return &CityName{value: value}, nil
}

func (name *CityName) String() string {
	return name.value
}

type City struct {
	id   *CityID
	name *CityName
}

func (city *City) ID() *CityID {
	return city.id
}

func (city *City) Name() *CityName {
	return city.name
}

func NewCity(id string, name string) (*City, error) {
	cityIDVo, err := NewCityID(id)
	if err != nil {
		return nil, err
	}
	cityNameVo, err := NewCityName(name)
	if err != nil {
		return nil, err
	}

	return &City{
		id:   cityIDVo,
		name: cityNameVo,
	}, nil
}

type CityRepository interface {
	FindByID(id CityID) (*City, error)
	FindAll() ([]*City, error)
}
