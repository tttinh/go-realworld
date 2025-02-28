package httpport

import (
	"time"

	"github.com/tinhtt/go-realworld/internal/domain"
)

type createCommentReq struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}

type commentRes struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Author    struct {
		Username  string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"author"`
}

func (res *commentRes) fromEntity(c domain.Comment) {
	res.ID = c.ID
	res.Body = c.Body
	res.CreatedAt = c.CreatedAt
	res.UpdatedAt = c.UpdatedAt
}

type commentsRes struct {
	Comments []commentRes `json:"comments"`
}

func (res *commentsRes) fromEntity(comments []domain.Comment) {
	for _, c := range comments {
		res.Comments = append(res.Comments, commentRes{
			ID:        c.ID,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}
}
