package entity

import (
	"time"
)

type Order struct {
	ID        string
	Status    string
	CreatedAt time.Time
}
