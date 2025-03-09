package httpendpoints

import (
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	secret   string
	duration time.Duration
}

func NewJWT(secret string, duration time.Duration) JWT {
	return JWT{
		secret:   secret,
		duration: duration,
	}
}

func (t JWT) Get(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func (t JWT) Validate(c *gin.Context) error {
	jwtToken, err := jwt.ParseWithClaims(
		t.Get(c),
		&jwt.RegisteredClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(t.secret), nil
		},
	)
	if err != nil {
		return err
	}
	if _, ok := jwtToken.Claims.(*jwt.RegisteredClaims); ok && jwtToken.Valid {
		return nil
	}
	return err
}

func (t JWT) GetUserID(c *gin.Context) (int, bool) {
	tokenString := t.Get(c)
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(t.secret), nil
	})
	if err != nil {
		return 0, false
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return 0, false
	}

	if !token.Valid {
		return 0, false
	}

	id, _ := strconv.Atoi(claims.Subject)
	return id, true
}

func (t JWT) Generate(id string) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   id,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(t.duration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(t.secret))
}
