package middlewares

import (
	"backend/env"
	"backend/models/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthenticateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Request.Cookie("auth_token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}

		tokenString := cookie.Value
		AuthClaims := &dto.AuthClaims{}

		token, err := jwt.ParseWithClaims(tokenString, AuthClaims, func(token *jwt.Token) (interface{}, error) {
			return []byte(env.Get("JWTKEY")), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized - invalid token"})
			c.Abort()
			return
		}

		c.Set("username", AuthClaims.Username)
		c.Next()
	}
}
