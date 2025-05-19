package model

import "time"

type Order struct {
	ID        string    `gorm:"column:order_id;primaryKey"`
	Status    string    `gorm:"column:status"`
	CreatedAt time.Time `gorm:"column:created_at"`
}
