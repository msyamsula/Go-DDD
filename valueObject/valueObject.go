package valueObject

import (
	"time"
	"github.com/google/uuid"
)

type Transaction struct {
	ID uuid.UUID
	amount int
	from uuid.UUID
	to uuid.UUID
	createdAt time.Time

}