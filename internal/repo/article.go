package repo

import (
	"github.com/tinhtt/go-realworld/internal/entity"
)

type Article interface {
	InsertArticle(userId int, a entity.Article) (entity.Article, error)
}
