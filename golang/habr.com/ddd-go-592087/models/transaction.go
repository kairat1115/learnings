package models

import (
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	// all values lowercase since they are immutable
	amount    int
	from      uuid.UUID
	to        uuid.UUID
	createdAt time.Time
}
