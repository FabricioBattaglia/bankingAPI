package api

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	db "github.com/FabricioBattaglia/bankingAPI/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Name   string `json:"name" binding:"required"`
	Cpf    string `json:"cpf" binding:"required"`
	Secret string `json:"secret" binding:"required"`
}

//create account handler
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashsecret, err := bcrypt.GenerateFromPassword([]byte(req.Secret), bcrypt.DefaultCost)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Name:    req.Name,
		Cpf:     req.Cpf,
		Secret:  string(hashsecret),
		Balance: 1000,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

//get account handler
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

//get list account handler

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

//get account handler
func (server *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParam{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
