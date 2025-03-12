package domain

import "time"

type Comment struct {
	Author    Author
	ID        int
	ArticleID int
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
