package api

import (
	"net/http"
	"time"

	db "github.com/adlanmsyariaty/postr/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreatePostRequest struct {
	Post string `json:"post" binding:"required"`
}

type CreatePostResponse struct {
	PostID uuid.UUID `json:"postId"`
}

func (server *Server) createPost(ctx *gin.Context) {
	var req CreatePostRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	post, err := server.store.CreatePost(ctx, db.CreatePostParams{
		ID:   uuid.New(),
		Post: req.Post,
		CreatedAt: time.Now(),
		CreatedBy: "system",
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err.Error()))
		return
	}

	res := CreatePostResponse{
		PostID: post.ID,
	}
	ctx.JSON(http.StatusCreated, successResponse(res))
}
