package api

import (
	"time"

	"github.com/tinhtt/go-realworld/internal/entity"
)

type CreateCommentReq struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}
type CommentRes struct {
	Id        int       `json:"id"`
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

type ListCommentsRes struct {
	Comments []CommentRes `json:"comments"`
}

func commentFromEntity(c entity.Comment) CommentRes {
	return CommentRes{
		Id:        c.Id,
		Body:      c.Body,
		CreatedAt: c.CreatedAt,
		UpdatedAt: c.UpdatedAt,
	}
}

func commentsFromEntity(comments []entity.Comment) ListCommentsRes {
	var res ListCommentsRes
	for _, c := range comments {
		res.Comments = append(res.Comments, CommentRes{
			Id:        c.Id,
			Body:      c.Body,
			CreatedAt: c.CreatedAt,
			UpdatedAt: c.UpdatedAt,
		})
	}

	return res
}
