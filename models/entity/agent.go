package entity

import (
	"time"
)

// Ini buat tabel di database buat di migrate.
type Agent struct {
	Id        uint   `json:"id" gorm:"primarykey"`
	Name      string `json:"name"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
