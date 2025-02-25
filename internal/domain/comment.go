package domain

import "time"

type Comment struct {
	ID        int
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
