package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type createArticleReq struct {
	Article struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Body        string   `json:"body"`
		Tags        []string `json:"tagList"`
	} `json:"article"`
}

func (h *Handler) createArticle(c *gin.Context) {
	authorID, _ := h.jwt.GetUserID(c)
	var req createArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	a := domain.NewArticle(
		authorID,
		req.Article.Title,
		req.Article.Description,
		req.Article.Body,
		req.Article.Tags,
	)

	a, err := h.articles.Add(c, a)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	detail, err := h.articles.GetDetail(c, authorID, a.Slug)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res articleRes
	res.fromEntity(detail)
	c.JSON(http.StatusOK, res)
}
