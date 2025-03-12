package httpendpoints

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tinhtt/go-realworld/internal/domain"
)

type commentData struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Author    struct {
		Name      string `json:"username"`
		Bio       string `json:"bio"`
		Image     string `json:"image"`
		Following bool   `json:"following"`
	} `json:"author"`
}

func (res *commentData) fromEntity(c domain.Comment) {
	res.ID = c.ID
	res.Body = c.Body
	res.CreatedAt = c.CreatedAt
	res.UpdatedAt = c.UpdatedAt
	res.Author.Name = c.Author.Name
	res.Author.Bio = c.Author.Bio
	res.Author.Image = c.Author.Image
	res.Author.Following = c.Author.Following
}

type commentRes struct {
	Comment commentData `json:"comment"`
}

func (res *commentRes) fromEntity(c domain.Comment) {
	res.Comment.fromEntity(c)
}

type createCommentReq struct {
	Comment struct {
		Body string `json:"body"`
	} `json:"comment"`
}

func (h *Handler) createComment(c *gin.Context) {
	userID, _ := h.jwt.GetUserID(c)
	var req createCommentReq
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

	comment := domain.Comment{
		Author:    domain.Author{ID: userID},
		ArticleID: a.ID,
		Body:      req.Comment.Body,
	}
	comment, err = h.articles.AddComment(c, comment)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	var res commentRes
	res.fromEntity(comment)
	c.JSON(http.StatusOK, res)
}
