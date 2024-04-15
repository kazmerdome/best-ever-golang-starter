package post

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter"
)

// Interface
//

type PostService interface {
	// Crud
	CreatePost(ctx context.Context, data CreateDto) (*Post, error)
	GetPost(ctx context.Context, id uuid.UUID) (*Post, error)
	ListPosts(ctx context.Context, where *WhereDto) ([]Post, error)
	UpdatePost(ctx context.Context, id uuid.UUID, data UpdateDto) (*Post, error)
	DeletePost(ctx context.Context, id uuid.UUID) error
	// Dataloader
	LoadPost(ctx context.Context, id uuid.UUID) (*Post, error)
}

type PostRepository interface {
	CreateOne(ctx context.Context, data CreateDto) (*Post, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*Post, error)
	GetMany(ctx context.Context, where *WhereDto) ([]Post, error)
	GetManyByIds(ctx context.Context, ids []uuid.UUID) ([]*Post, error)
	UpdateOneById(ctx context.Context, id uuid.UUID, data UpdateDto) (*Post, error)
	DeleteOne(ctx context.Context, id uuid.UUID) error
}

type PostDataloader interface {
	ItemLoader(ctx context.Context, id uuid.UUID) (*Post, error)
}

// Entity & Enum
//

type Post struct {
	Id        uuid.UUID  `json:"id" bson:"_id"`
	Title     string     `json:"title" bson:"title"`
	Slug      string     `json:"slug" bson:"slug"`
	Category  uuid.UUID  `json:"category" bson:"category"`
	Status    PostStatus `json:"status" bson:"status"`
	Content   string     `json:"content" bson:"content"`
	CreatedAt time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt" bson:"updatedAt"`
}

type PostStatus string

const (
	StatusActive   PostStatus = "ACTIVE"
	StatusPending  PostStatus = "PENDING"
	StatusArchived PostStatus = "ARCHIVED"
)

// Dto
//

type CreateDto struct {
	Title    string     `json:"title" bson:"title"`
	Slug     *string    `json:"slug" bson:"slug"`
	Category uuid.UUID  `json:"category" bson:"category"`
	Status   PostStatus `json:"status" bson:"status"`
	Content  *string    `json:"content" bson:"content"`
}

type WhereDto struct {
	Title      *filter.StringFilter     `json:"title" bson:"title"`
	Slug       *filter.StringFilter     `json:"slug" bson:"slug"`
	Category   *filter.UuidFilter       `json:"category" bson:"category"`
	Status     *PostStatus              `json:"status" bson:"status"`
	Sort       *filter.SortFilter       `json:"sort" bson:"sort"`
	Pagination *filter.PaginationFilter `json:"pagination" bson:"pagination"`
}

type UpdateDto struct {
	Title    *string     `json:"title" bson:"title"`
	Slug     *string     `json:"slug" bson:"slug"`
	Category *uuid.UUID  `json:"category" bson:"category"`
	Status   *PostStatus `json:"status" bson:"status,omitempty"`
	Content  *string     `json:"content" bson:"content"`
}
