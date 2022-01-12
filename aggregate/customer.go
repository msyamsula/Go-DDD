package aggregate

import (
	"errors"

	"github.com/google/uuid"
	"github.com/msyamsula/Go-DDD/entity"
	"github.com/msyamsula/Go-DDD/valueObject"
)

var (
	ErrInvalidPerson = errors.New("Invalid Person")
)

type Customer struct {
	person      *entity.Person
	items       []*entity.Item
	transations []valueObject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}

	items := []*entity.Item{}

	transactions := []valueObject.Transaction{}

	customer := Customer{
		person:      person,
		items:       items,
		transations: transactions,
	}

	return customer, nil
}
