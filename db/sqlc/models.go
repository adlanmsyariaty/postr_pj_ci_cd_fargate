// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID      `json:"id"`
	Post      string         `json:"post"`
	CreatedAt time.Time      `json:"created_at"`
	CreatedBy string         `json:"created_by"`
	UpdatedAt sql.NullTime   `json:"updated_at"`
	UpdatedBy sql.NullString `json:"updated_by"`
}
