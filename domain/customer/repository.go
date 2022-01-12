package customer

import (
	"github.com/google/uuid"
	"github.com/msyamsula/Go-DDD/aggregate"
)

type CustomerRepo interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
}
