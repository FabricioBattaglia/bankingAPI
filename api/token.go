package api

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type MyCustomClaims struct {
	account string `json:"account"`
	jwt.StandardClaims
}

func ValidateToken(ctx *gin.Context) *jwt.Token {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "No Authorization header provided"})
		return nil
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "No Authorization header provided"})
		return nil
	}

	authData, err := jwt.ParseWithClaims(token, &MyCustomClaims{}, func(jsonwebtoken *jwt.Token) (interface{}, error) { return []byte("CHAVE_SECRETA"), nil })
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return nil
	}

	return authData
}
