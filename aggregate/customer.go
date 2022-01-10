package aggregate

import (
	"github.com/msyamsula/Go-DDD/valueObject"
	"github.com/msyamsula/Go-DDD/entity"
)

type Customer struct {
	person *entity.Person
	items []*entity.Item
	transations []valueObject.Transaction
}