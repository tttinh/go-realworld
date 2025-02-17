package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/repo"
)

func NewHttpHandler(articleRepo repo.Articles) http.Handler {
	router := gin.Default()

	ah := &ArticleHandler{articles: articleRepo}
	ah.Mount(router)

	return router
}
