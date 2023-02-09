// We'll implement http server here

package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/sachin-sharma-IN/personalbank/db/sqlc"
)

// Server serves HTTP requests for our banking service.
// store of type db.store will allow us to interact with db.
// router will help us send each API req to correct handler.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer creats a new HTTP server and setup routing.
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// handler functions.
	// we can pass one or multiple handler func. if we pass multiple funcs,
	// last one should be real handler and all other funcs should be middleware.

	// For now, we don't have middleware so we'll just pass createA/c. This is method of server struct which needs to be implemented.
	// Reason it is of server struct is bcz it needs access to store obj to create a/c in db.
	router.POST("/accounts", server.createAccount)

	server.router = router
	return server
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
