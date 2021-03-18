package api

import (
	"database/sql"
	"fmt"

	//"github.com/FabricioBattaglia/bankingAPI/api"
	//"fmt"
	"net/http"

	db "github.com/FabricioBattaglia/bankingAPI/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createTransferRequest struct {
	AccountOriginID      int64 `json:"account_origin_id" binding:"required,min=1"`
	AccountDestinationID int64 `json:"account_destination_id" binding:"required,min=1"`
	Amount               int64 `json:"amount" binding:"required,gt=0"`
}

//create transfer handler
func (server *Server) createTransfer(ctx *gin.Context) {
	tokenContent := ValidateToken(ctx)

	if claims, ok := tokenContent.Claims.(*MyCustomClaims); ok && tokenContent.Valid {
		fmt.Printf("%v %v", claims.account, claims.StandardClaims.ExpiresAt)
		fmt.Println(claims.account)
	}

	fmt.Println(tokenContent.Claims)

	fmt.Printf("create_transfer:start\n")
	var req createTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, req.AccountOriginID) {
		return
	}

	if !server.validAccount(ctx, req.AccountDestinationID) {
		return
	}

	arg := db.TransferTxParams{
		AccountOriginID:      req.AccountOriginID,
		AccountDestinationID: req.AccountDestinationID,
		Amount:               req.Amount,
	}

	result, err := server.store.TransferTx(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	fmt.Printf("create_transfer:end\n")
	ctx.JSON(http.StatusOK, result)
}

func (server *Server) validAccount(ctx *gin.Context, accountID int64) bool {
	account, err := server.store.GetAccount(ctx, accountID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return false
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return false
	}
	if account.ID != accountID {
		return false
	}
	return true
}
