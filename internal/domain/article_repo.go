package domain

import "context"

type ArticleRepo interface {
	FindBySlug(ctx context.Context, viewerID int, slug string) (Article, error)
	Insert(ctx context.Context, a Article) (Article, error)
	Update(ctx context.Context, a Article) (Article, error)
	Delete(ctx context.Context, slug string) error
}
