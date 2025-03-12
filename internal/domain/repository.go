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
	GetAllArticles(
		ctx context.Context,
		viewerID int,
		offset int,
		limit int,
	) ([]ArticleDetail, error)

	GetAllArticlesByAuthor(
		ctx context.Context,
		viewerID int,
		offset int,
		limit int,
		author string,
	) ([]ArticleDetail, error)

	GetAllArticlesByFavorited(
		ctx context.Context,
		viewerID int,
		offset int,
		limit int,
		favoritedUser string,
	) ([]ArticleDetail, error)

	GetAllArticlesByTag(
		ctx context.Context,
		viewerID int,
		offset int,
		limit int,
		tag string,
	) ([]ArticleDetail, error)

	GetDetail(ctx context.Context, viewerID int, slug string) (ArticleDetail, error)

	Get(ctx context.Context, slug string) (Article, error)
	Add(ctx context.Context, a Article) (Article, error)
	Edit(ctx context.Context, a Article) (Article, error)
	Remove(ctx context.Context, id int) error

	AddFavorite(ctx context.Context, userID int, articleID int) error
	RemoveFavorite(ctx context.Context, userID int, articleID int) error

	GetAllComments(ctx context.Context, viewerID int, articleID int) ([]Comment, error)
	GetComment(ctx context.Context, viewerID int, commentID int) (Comment, error)
	AddComment(ctx context.Context, c Comment) (Comment, error)
	RemoveComment(ctx context.Context, id int) error

	GetAllTags(ctx context.Context) ([]string, error)
}
