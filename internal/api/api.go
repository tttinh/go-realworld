package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/repo"
)

func NewHttpHandler(
	articles repo.Articles,
	comments repo.Comments,
) http.Handler {
	router := gin.Default()

	ah := ArticleHandler{articles: articles}
	ah.Mount(router)

	ch := CommentHandler{articles: articles, comments: comments}
	ch.Mount(router)

	uh := UserHandler{}
	uh.Mount(router)

	return router
}
