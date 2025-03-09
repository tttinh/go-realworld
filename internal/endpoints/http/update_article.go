package httpendpoints

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type updateArticleReq struct {
	Article struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Body        string `json:"body"`
	} `json:"article"`
}

func (h *Handler) updateArticle(c *gin.Context) {
	authorID, _ := h.jwt.GetUserID(c)
	var req updateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		abort(c, err)
		return
	}

	if authorID != a.AuthorID {
		c.AbortWithError(http.StatusForbidden, domain.ErrForbidden)
		return
	}

	err = a.Update(req.Article.Title, req.Article.Description, req.Article.Body)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	a, err = h.articles.Edit(c, a)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	detail, err := h.articles.GetDetail(c, authorID, slug)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res articleRes
	res.fromEntity(detail)
	c.JSON(http.StatusOK, res)
}
