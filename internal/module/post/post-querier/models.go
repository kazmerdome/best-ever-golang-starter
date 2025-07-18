// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package postQuerier

import (
	"database/sql"

	"github.com/google/uuid"
)

type Post struct {
	ID        uuid.UUID
	Title     string
	Slug      string
	Category  uuid.UUID
	Status    string
	Content   sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
