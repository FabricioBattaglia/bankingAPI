package api

import (
	db "github.com/FabricioBattaglia/bankingAPI/db/sqlc"
	"github.com/gin-gonic/gin"
)

//serves HTTP requests
type Server struct {
	store  *db.Store
	router *gin.Engine
}

//create new HTTP server
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.GET("/transfers", server.listTransfers)
	router.POST("/transfers", server.createTransfer)
	router.POST("/login", server.login)

	server.router = router
	return server
}

//run server HTTP on address provided
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
