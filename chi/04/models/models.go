package models


type Car struct {
	ID     string
	Brand  string
	Models []string
}

type CarRepo struct {
	Cars []Car
}

func (c *CarRepo) GetAllCars() ([]string, error) {
	var list []string
	for _, car := range c.Cars {
		list = append(list, car.Brand)
	}
	return list, nil
}

