package httpendpoints

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func authMiddleware(tokenSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		err := validateToken(getToken(c), tokenSecret)
		if err != nil {
			c.JSON(http.StatusUnauthorized, newErrorRes(err))
			c.Abort()
			return
		}
		c.Next()
	}
}

func getToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""

}

func validateToken(token string, tokenSecret string) error {
	jwtToken, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return err
	}
	if _, ok := jwtToken.Claims.(*jwt.RegisteredClaims); ok && jwtToken.Valid {
		return nil
	}
	return err
}

// func getIDFromJWT(tokenString string) (string, error) {
// 	secret := viper.GetString("API_SECRET")
// 	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		return []byte(secret), nil
// 	})
// 	if err != nil {
// 		return "", err
// 	}
// 	claims, ok := token.Claims.(*jwt.RegisteredClaims)
// 	if !ok {
// 		return "", err
// 	}
// 	// check if token is not expired
// 	if !token.Valid {
// 		return "", err
// 	}
// 	return claims.Subject, nil
// }

// func getIDFromToken(token string) string {
// 	id, _ := getIDFromJWT(token)
// 	return id
// }

// func getIDFromHeaderr(c *gin.Context) string {
// 	tokenString := getToken(c)
// 	id, _ := getIDFromJWT(tokenString)
// 	return id
// }

func generateToken(id string, tokenSecret string, tokenDuration time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   id,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenDuration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(tokenSecret))
}
