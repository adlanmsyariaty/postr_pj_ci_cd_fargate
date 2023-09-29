package api

import (
	"github.com/adlanmsyariaty/postr/config"
	db "github.com/adlanmsyariaty/postr/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config config.Config
	store  *db.Store
	router *gin.Engine
}

func NewServer(config config.Config, store *db.Store) (*Server, error) {
	server := &Server{
		config: config,
		store:  store,
	}

	// Initialize router
	router := gin.Default()

	// Initialize routes
	router.POST("/posts", server.createPost)

	server.router = router
	return server, nil
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(errorMessage string) gin.H {
	return gin.H{
		"success": false,
		"message": errorMessage,
	}
}

func successResponse[R any](data R) gin.H {
	return gin.H{
		"success": true,
		"data":    data,
	}
}
