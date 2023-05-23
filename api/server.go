// We'll implement http server here

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/sachin-sharma-IN/personalbank/db/sqlc"
)

// Server serves HTTP requests for our banking service.
// store of type db.store will allow us to interact with db.
// router will help us send each API req to correct handler.
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(store db.Store) (*Server, error) {
	server := &Server{store: store}
	router := gin.Default()

	// binding.Validator.Engine() will return the current validator which gin is using.
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		err := v.RegisterValidation("currency", validCurrency)
		if err != nil {
			return nil, fmt.Errorf("RegisterValidation failed with %v", err)
		}
	}

	// handler functions.
	// we can pass one or multiple handler func. if we pass multiple funcs,
	// last one should be real handler and all other funcs should be middleware.

	// For now, we don't have middleware, so we'll just pass createA/c. This is method of server struct which needs to be implemented.
	// Reason it is of server struct is bcz it needs access to store obj to create a/c in db.
	router.POST("/accounts", server.createAccount)
	// Api to get a specific account by ID.
	// colon aka : before id is the way to tell gin that this is URI param.
	router.GET("/accounts/:id", server.getAccount)
	// endpoint to listAccounts. We'll use pagination and get input params
	// from query params in request body. Since input will come from query params, we'll use
	// /accounts endpoint
	router.GET("/accounts", server.listAccount)

	router.POST("/transfers", server.createTransfer)
	router.POST("/users", server.createUser)

	server.router = router
	return server, nil
}

// Start runs the HTTP server on a specific address.
// Notice router field is private in server i.e. it can't be accessed outside api package.
// That's why we've public Start().
// TO-DO: we'll add more logic to gracefully shutdown server.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
