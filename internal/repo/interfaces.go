package repo

import (
	"github.com/tinhtt/go-realworld/internal/entity"
)

type Article interface {
	FindBySlug(slug string) (entity.Article, error)
	Insert(a entity.Article) (entity.Article, error)
	Update(a entity.Article) (entity.Article, error)
	Delete(slug string) error
}

type Comment interface {
	FindByArticleId(id int) ([]entity.Comment, error)
	Insert(slug string, c entity.Comment) (entity.Comment, error)
	Delete(id int) error
}
