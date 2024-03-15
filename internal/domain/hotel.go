package domain

import "time"

type Hotel struct {
	ID        string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
