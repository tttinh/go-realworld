package httpendpoints

import (
	"errors"
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
	userID, _ := h.jwt.GetUserID(c)
	var req updateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	slug := c.Param("slug")
	a, err := h.articles.Get(c, slug)
	if err != nil {
		abort(c, err)
		return
	}

	if userID != a.AuthorID {
		c.AbortWithError(http.StatusForbidden, domain.ErrForbidden)
		return
	}

	err = a.Update(req.Article.Title, req.Article.Description, req.Article.Body)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ok := false
	attempts := 3
	for attempts > 0 {
		_, err := h.articles.Edit(c, a)
		if err == nil {
			ok = true
			break
		}

		if errors.Is(err, domain.ErrDuplicateKey) {
			attempts -= 1
			a.NewSlug()
			continue
		} else {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}

	if !ok {
		c.AbortWithError(http.StatusBadRequest, domain.ErrDuplicateKey)
		return
	}

	detail, err := h.articles.GetDetail(c, userID, slug)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res articleRes
	res.fromEntity(detail)
	c.JSON(http.StatusOK, res)
}
