package httpport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func NewHandler(
	users domain.UserRepo,
	articles domain.ArticleRepo,
	comments domain.CommentRepo,
) http.Handler {
	h := gin.Default()
	r := h.Group("/api")

	uh := userHandler{users: users}
	ah := articleHandler{articles: articles}
	ch := commentHandler{articles: articles, comments: comments}

	r.POST("/users/login", uh.login)
	r.POST("/users", uh.register)
	r.GET("/user", uh.read)
	r.PUT("/user", uh.edit)

	r.GET("/articles/feed", ah.browseFeed)
	r.GET("/articles", ah.browse)
	r.POST("/articles", ah.add)
	r.GET("/articles/:slug", ah.read)
	r.PUT("/articles/:slug", ah.edit)
	r.DELETE("/articles/:slug", ah.delete)

	r.GET("/articles/:slug/comments", ch.browse)
	r.POST("/articles/:slug/comments", ch.add)
	r.DELETE("/articles/:slug/comments/:id", ch.delete)

	return h
}
