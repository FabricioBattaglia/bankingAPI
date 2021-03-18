package api

import (
	"errors"
	"strings"

	jwtToken "github.com/FabricioBattaglia/bankingAPI/token"

	"github.com/gin-gonic/gin"
)

func ValidateToken(ctx *gin.Context) (int64, error) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		return 0, errors.New("no Authorization header provided")
	}

	token := strings.TrimPrefix(auth, "Bearer ")
	if token == auth {
		return 0, errors.New("no Authorization header provided")
	}

	claims, err := jwtToken.ParseJWT(token)
	if err != nil {
		return 0, err
	}

	return claims.Account, nil
}
