package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHttpHandler() http.Handler {
	router := gin.Default()

	ah := &ArticleHandler{}
	ah.Mount(router)

	return router
}
