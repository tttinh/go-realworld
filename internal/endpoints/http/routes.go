package httpendpoints

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/config"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type Handler struct {
	jwt      JWT
	articles domain.ArticleRepo
	users    domain.UserRepo
}

func NewHandler(
	log *slog.Logger,
	cfg config.Config,
	users domain.UserRepo,
	articles domain.ArticleRepo,
) http.Handler {
	jwt := NewJWT(cfg.HTTPServer.JWTSecret, cfg.HTTPServer.JWTDuration)
	h := Handler{
		jwt:      jwt,
		articles: articles,
		users:    users,
	}

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	api := router.Group("/api").Use(LogMiddleware(log), ErrorMiddleware())

	// public APIs
	api.POST("/users/login", h.loginUser)
	api.POST("/users", h.registerUser)
	api.GET("/articles", h.browseArticles)
	api.GET("/articles/:slug/comments", h.browseComments)
	api.GET("/articles/:slug", h.getArticle)
	api.GET("/profiles/:username", h.getProfile)
	api.GET("/tags", h.browseTags)

	// private APIs
	api.Use(AuthMiddleware(jwt))

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
