package httpendpoints

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

func NewHandler(
	users domain.UserRepo,
	articles domain.ArticleRepo,
) http.Handler {
	const tokenSecret = "ABC"
	const tokenDuration = 1 * time.Minute
	h := gin.Default()
	r := h.Group("/api")

	a := articleHandler{articles: articles}
	u := userHandler{
		tokenSecret:   tokenSecret,
		tokenDuration: tokenDuration,
		users:         users,
	}

	// public APIs
	r.POST("/users/login", u.loginUser)
	r.POST("/users", u.registerUser)
	r.GET("/articles", a.browseArticles)
	r.GET("/articles/:slug", a.getArticle)
	r.GET("/articles/:slug/comments", a.browseComments)
	r.GET("/profiles/:username", u.getProfile)

	// private APIs
	r.Use(authMiddleware(tokenSecret))

	// user
	r.GET("/user", u.getCurrentUser)
	r.PUT("/user", u.updateCurrentUser)

	// profile
	r.POST("/profiles/:username/follow", u.followUser)
	r.DELETE("/profiles/:username/follow", u.unfollowUser)

	// article
	r.GET("/articles/feed", a.getFeed)
	r.POST("/articles", a.createArticle)
	r.PUT("/articles/:slug", a.updateArticle)
	r.DELETE("/articles/:slug", a.deleteArticle)
	r.POST("/articles/:slug/favorite", a.favoriteArticle)
	r.DELETE("/articles/:slug/favorite", a.unfavoriteArticle)

	// comment
	r.POST("/articles/:slug/comments", a.createComment)
	r.DELETE("/articles/:slug/comments/:id", a.deleteComment)

	return h
}
