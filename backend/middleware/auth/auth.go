package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	UserID string
	Email  string
	jwt.StandardClaims
}

var jwtKey = []byte("my_secret_key")

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA)+1:]

		// Initialize a new instance of `Claims`
		claims := &Claims{}

		tkn, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				c.JSON(http.StatusInternalServerError, nil)
				return
			}
			c.JSON(http.StatusBadRequest, nil)
			return
		}
		if !tkn.Valid {
			c.JSON(http.StatusUnauthorized, nil)
			return
		}

		c.Set("email", claims.Email)
		c.Set("user_id", claims.UserID)

		c.Next()
	}
}
