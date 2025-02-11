package mem

import (
	"errors"
	"time"

	"github.com/tinhtt/go-realworld/internal/entity"
)

type ArticleRepo struct {
	count    int
	articles []*entity.Article
	slugs    map[string]*entity.Article
}

func NewArticleRepo() *ArticleRepo {
	return &ArticleRepo{
		count:    0,
		articles: []*entity.Article{},
		slugs:    map[string]*entity.Article{},
	}
}

func (repo *ArticleRepo) FindBySlug(slug string) (entity.Article, error) {
	a := repo.slugs[slug]
	if a == nil {
		return entity.Article{}, errors.New("not found")
	}

	return *a, nil
}

func (repo *ArticleRepo) Insert(a entity.Article) (entity.Article, error) {
	if _, ok := repo.slugs[a.Slug]; ok {
		return entity.Article{}, errors.New("already existed")
	}

	repo.count += 1
	a.Id = repo.count
	a.CreatedAt = time.Now()
	a.UpdatedAt = time.Now()
	repo.articles = append(repo.articles, &a)
	repo.slugs[a.Slug] = &a

	return a, nil
}

func (repo *ArticleRepo) Update(a entity.Article) (entity.Article, error) {
	dbArticle, ok := repo.slugs[a.Slug]
	if !ok {
		return entity.Article{}, errors.New("not found")
	}

	dbArticle.Title = a.Title
	dbArticle.Description = a.Description
	dbArticle.Body = a.Body
	dbArticle.UpdatedAt = time.Now()

	return *dbArticle, nil
}

func (repo *ArticleRepo) Delete(slug string) error {
	for i, v := range repo.articles {
		if v.Slug == slug {
			repo.articles = append(repo.articles[:i], repo.articles[i+1:]...)
		}
	}
	delete(repo.slugs, slug)
	return nil
}
