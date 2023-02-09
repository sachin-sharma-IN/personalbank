package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/sachin-sharma-IN/personalbank/db/sqlc"
)

// Creating this struct to store new createA/C req. We'll get below params from
//
//		body of HTTP req.
//	 gin uses go validator package internally. binding:required makes field mandatory in req.
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

// Decclare createA/c func with server pointer receiver and input is ctx of gin.Context type. Why?
// look at router.post sign. in api/server.go. HandlerFunc is decalred with context input.
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, account)
}
