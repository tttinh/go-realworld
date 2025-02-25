package domain

import "context"

type CommentRepo interface {
	FindByArticleId(ctx context.Context, id int) ([]Comment, error)
	Insert(ctx context.Context, slug string, c Comment) (Comment, error)
	Delete(ctx context.Context, id int) error
}
