package postgres

import (
	"context"

	"github.com/tinhtt/go-realworld/internal/domain"
)

func (r *Articles) GetAllArticles(
	ctx context.Context,
	viewerID int,
	offset int,
	limit int,
) ([]domain.ArticleDetail, error) {
	return nil, nil
}

func (r *Articles) GetAllArticlesByAuthor(
	ctx context.Context,
	viewerID int,
	offset int,
	limit int,
	author string,
) ([]domain.ArticleDetail, error) {
	return nil, nil

}

func (r *Articles) GetAllArticlesByFavorited(
	ctx context.Context,
	viewerID int,
	offset int,
	limit int,
	favoritedUser string,
) ([]domain.ArticleDetail, error) {
	return nil, nil

}

func (r *Articles) GetAllArticlesByTag(
	ctx context.Context,
	viewerID int,
	offset int,
	limit int,
	tag string,
) ([]domain.ArticleDetail, error) {
	return nil, nil
}
