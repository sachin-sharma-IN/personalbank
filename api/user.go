package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	db "github.com/sachin-sharma-IN/personalbank/db/sqlc"
	"github.com/sachin-sharma-IN/personalbank/util"
)

/*
	  Creating this struct to store new CreateUser req. We'll get below params from
			body of HTTP req.
			gin uses go validator package internally. binding:required makes field mandatory in req.
	   and alphanum prevents it from containing any special character.
*/
type createUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
	FullName string `json:"full_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

// Declare createA/c func with server pointer receiver and input is ctx of gin.Context type. Why?
// look at router.post sign. in api/server.go. HandlerFunc is decalred with context input.
func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
	}
	arg := db.CreateUserParams{
		Username:       req.Username,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
		Email:          req.Email,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		// Try returning err of *pq.error type
		if pqErr, ok := err.(*pq.Error); ok {
			// Print log to see this error's codename
			// log.Println(pqErr.Code.Name())
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
