// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	GetPostById(ctx context.Context, id uuid.UUID) (Post, error)
}

var _ Querier = (*Queries)(nil)
