package httpport

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type articleHandler struct {
	articles domain.ArticleRepo
}

func (h *articleHandler) browseFeed(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *articleHandler) browse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "OK"})
}

func (h *articleHandler) read(c *gin.Context) {
	viewerID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetDetail(c, viewerID, slug)
	if err != nil {
		error400(c, err)
		return
	}

	var res articleRes
	res.fromEntity(a)
	ok(c, res)
}

func (h *articleHandler) edit(c *gin.Context) {
	authorID := 1
	var req updateArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		error400(c, err)
		return
	}

	if authorID != a.AuthorID {
		error403(c)
		return
	}

	err = a.Update(req.Article.Title, req.Article.Description, req.Article.Body)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	a, err = h.articles.Edit(c, a)
	if err != nil {
		error400(c, err)
		return
	}

	detail, err := h.articles.GetDetail(c, authorID, slug)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res articleRes
	res.fromEntity(detail)
	ok(c, res)
}

func (h *articleHandler) add(c *gin.Context) {
	authorID := 1
	var req createArticleReq
	if err := c.ShouldBindJSON(&req); err != nil {
		error400(c, err)
		return
	}

	a := domain.NewArticle(
		authorID,
		req.Article.Title,
		req.Article.Description,
		req.Article.Body,
	)
	a, err := h.articles.Add(c, a)
	if err != nil {
		error400(c, err)
		return
	}

	detail, err := h.articles.GetDetail(c, authorID, a.Slug)
	if err != nil {
		log.Println(err)
		error500(c)
		return
	}

	var res articleRes
	res.fromEntity(detail)
	ok(c, res)
}

func (h *articleHandler) delete(c *gin.Context) {
	authorID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		error400(c, err)
		return
	}

	if authorID != a.AuthorID {
		error403(c)
		return
	}

	if err := h.articles.Remove(c, a.ID); err != nil {
		log.Println(err)
		error500(c)
		return
	}
}

func (h *articleHandler) favorite(c *gin.Context) {
	authorID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		error400(c, err)
		return
	}

	if err := h.articles.AddFavorite(c, authorID, a.ID); err != nil {
		log.Println(err)
		error500(c)
		return
	}
}

func (h *articleHandler) unfavorite(c *gin.Context) {
	authorID := 1
	slug := c.Param("slug")
	a, err := h.articles.GetBySlug(c, slug)
	if err != nil {
		error400(c, err)
		return
	}

	if err := h.articles.RemoveFavorite(c, authorID, a.ID); err != nil {
		log.Println(err)
		error500(c)
		return
	}
}
