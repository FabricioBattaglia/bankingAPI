package api

import (
	"database/sql"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginRequest struct {
	Cpf    string `json:"cpf" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}

//create login handler
func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccountByCpf(ctx, req.Cpf)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Secret), []byte(req.Secret))
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.MapClaims{"account": account.ID})
	jsonwebtoken, err := token.SignedString([]byte("CHAVE_SECRETA"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": jsonwebtoken})
}
