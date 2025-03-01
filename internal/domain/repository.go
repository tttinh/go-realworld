package domain

import "context"

type UserRepo interface {
	Get(ctx context.Context, id int) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Insert(ctx context.Context, u User) (User, error)
	Update(ctx context.Context, u User) (User, error)
}

type ArticleRepo interface {
	Get(ctx context.Context, slug string) (Article, error)
	Insert(ctx context.Context, a Article) (Article, error)
	Update(ctx context.Context, a Article) (Article, error)
	Delete(ctx context.Context, id int) error

	GetDetail(ctx context.Context, viewerID int, slug string) (ArticleDetail, error)
}

type CommentRepo interface {
	Get(ctx context.Context, id int) (Comment, error)
	Insert(ctx context.Context, c Comment) (Comment, error)
	Delete(ctx context.Context, id int) error

	FindAllByArticleId(ctx context.Context, id int) ([]Comment, error)
}
