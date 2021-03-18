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
	// AccountOriginID      int64 `json:"account_origin_id" binding:"required,min=1"`
	AccountDestinationID int64 `json:"account_destination_id" binding:"required,min=1"`
	Amount               int64 `json:"amount" binding:"required,gt=0"`
}

//create transfer handler
func (server *Server) createTransfer(ctx *gin.Context) {
	AccountOriginID, err := ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	fmt.Printf("create_transfer:start\n")
	var req createTransferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if !server.validAccount(ctx, int64(AccountOriginID)) {
		return
	}

	if !server.validAccount(ctx, req.AccountDestinationID) {
		return
	}

	arg := db.TransferTxParams{
		AccountOriginID:      int64(AccountOriginID),
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

//get transfer handler

type listTransferRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=1,max=10"`
}

func (server *Server) listTransfers(ctx *gin.Context) {
	AccountOriginID, err := ValidateToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	var req listTransferRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.ListTransfersParam{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	transfers, err := server.store.ListTransfers(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transfers)
}
