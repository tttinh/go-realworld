package domain

import "context"

type ArticleRepo interface {
	GetDetail(ctx context.Context, viewerID int, slug string) (ArticleDetail, error)
	Get(ctx context.Context, slug string) (Article, error)
	Insert(ctx context.Context, a Article) (Article, error)
	Update(ctx context.Context, a Article) (Article, error)
	Delete(ctx context.Context, slug string) error
}
