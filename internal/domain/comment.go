package domain

import "time"

type Comment struct {
	ID        int
	AuthorID  int
	ArticleID int
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
