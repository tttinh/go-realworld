package httpport

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func NewHandler(
	users domain.UserRepo,
	articles domain.ArticleRepo,
) http.Handler {
	h := gin.Default()
	r := h.Group("/api")

	u := userHandler{users: users}
	r.POST("/users/login", u.login)
	r.POST("/users", u.register)
	r.GET("/user", u.read)
	r.PUT("/user", u.edit)

	a := articleHandler{articles: articles}
	r.GET("/articles/feed", a.browseFeed)
	r.GET("/articles", a.browse)
	r.POST("/articles", a.add)
	r.GET("/articles/:slug", a.read)
	r.PUT("/articles/:slug", a.edit)
	r.DELETE("/articles/:slug", a.delete)
	r.POST("/articles/:slug/favorite", a.favorite)
	r.DELETE("/articles/:slug/favorite", a.unfavorite)

	r.GET("/articles/:slug/comments", a.browseComments)
	r.POST("/articles/:slug/comments", a.addComment)
	r.DELETE("/articles/:slug/comments/:id", a.deleteComment)

	return h
}
