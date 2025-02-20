package utils

import (
	"backend/env"
	"backend/models/dto"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func SetAuthToken(c *gin.Context, username string) string {

	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username required"})
		return ""
	}

	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &dto.AuthClaims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtkey := env.Get("JWTKEY")
	tokenString, err := token.SignedString([]byte(jwtkey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return ""
	}

	return tokenString

}
