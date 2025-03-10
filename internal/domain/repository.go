package domain

import "context"

type UserRepo interface {
	GetByID(ctx context.Context, id int) (User, error)
	GetByEmail(ctx context.Context, email string) (User, error)
	Add(ctx context.Context, u User) (User, error)
	Edit(ctx context.Context, u User) (User, error)

	GetProfile(ctx context.Context, followerID int, followingUsername string) (Profile, error)
	Follow(ctx context.Context, followerID int, followingID int) error
	Unfollow(ctx context.Context, followerID int, followingID int) error
}

type ArticleRepo interface {
	GetDetail(ctx context.Context, viewerID int, slug string) (ArticleDetail, error)

	Get(ctx context.Context, slug string) (Article, error)
	Add(ctx context.Context, a Article) (Article, error)
	Edit(ctx context.Context, a Article) (Article, error)
	Remove(ctx context.Context, id int) error

	AddFavorite(ctx context.Context, userID int, articleID int) error
	RemoveFavorite(ctx context.Context, userID int, articleID int) error

	AddComment(ctx context.Context, c Comment) (Comment, error)
	RemoveComment(ctx context.Context, id int) error
	GetAllComments(ctx context.Context, id int) ([]Comment, error)
}
