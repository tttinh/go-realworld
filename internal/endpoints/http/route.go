package httpendpoints

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type Handler struct {
	token    Token
	articles domain.ArticleRepo
	users    domain.UserRepo
}

func NewHandler(
	users domain.UserRepo,
	articles domain.ArticleRepo,
) http.Handler {
	t := Token{
		secret:   "ABC",
		duration: 5 * time.Minute,
	}

	h := Handler{
		token:    t,
		articles: articles,
		users:    users,
	}

	router := gin.Default()
	api := router.Group("/api")

	// public APIs
	api.POST("/users/login", h.loginUser)
	api.POST("/users", h.registerUser)
	api.GET("/articles", h.browseArticles)
	api.GET("/articles/:slug", h.getArticle)
	api.GET("/articles/:slug/comments", h.browseComments)
	api.GET("/profiles/:username", h.getProfile)

	// private APIs
	api.Use(authMiddleware(t))

	// user
	api.GET("/user", h.getCurrentUser)
	api.PUT("/user", h.updateCurrentUser)

	// profile
	api.POST("/profiles/:username/follow", h.followUser)
	api.DELETE("/profiles/:username/follow", h.unfollowUser)

	// article
	api.GET("/articles/feed", h.browseFeed)
	api.POST("/articles", h.createArticle)
	api.PUT("/articles/:slug", h.updateArticle)
	api.DELETE("/articles/:slug", h.deleteArticle)
	api.POST("/articles/:slug/favorite", h.favoriteArticle)
	api.DELETE("/articles/:slug/favorite", h.unfavoriteArticle)

	// comment
	api.POST("/articles/:slug/comments", h.createComment)
	api.DELETE("/articles/:slug/comments/:id", h.deleteComment)

	return router
}
