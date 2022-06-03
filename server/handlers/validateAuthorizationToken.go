package handlers

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"jwt-filter/server/conf"
	"net/http"
	"strings"
)

func ValidateAuthorizationToken(c *gin.Context) {
	var statusCode = http.StatusUnauthorized
	var msg = ""

	AuthorizationHeader := c.GetHeader("Authorization")
	if !strings.HasPrefix(AuthorizationHeader, "Bearer ") {
		c.JSON(statusCode, gin.H{"error": msg})
		return
	}

	tokenString := strings.Split(AuthorizationHeader, " ")[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.JwtSecretKey), nil
	})

	if token.Valid {
		statusCode = http.StatusOK
		msg = "Good token"
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			msg = "That's not even a token"
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			msg = "Token expired2"
		} else {
			msg = "Couldn't handle this token:"
		}
	} else {
		msg = "Couldn't handle this token:"
	}

	c.JSON(statusCode, gin.H{"error": msg})
}
