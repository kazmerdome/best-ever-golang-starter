package category

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/kazmerdome/best-ever-golang-starter/internal/util/filter"
)

// Interface
//

type CategoryService interface {
	// Crud
	CreateCategory(ctx context.Context, data CreateDto) (*Category, error)
	GetCategory(ctx context.Context, id uuid.UUID) (*Category, error)
	ListCategories(ctx context.Context, where *WhereDto) ([]Category, error)
	UpdateCategory(ctx context.Context, id uuid.UUID, data UpdateDto) (*Category, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) error
	// Dataloader
	LoadCategory(ctx context.Context, id uuid.UUID) (*Category, error)
}

type CategoryRepository interface {
	CreateOne(ctx context.Context, data CreateDto) (*Category, error)
	GetOneById(ctx context.Context, id uuid.UUID) (*Category, error)
	GetMany(ctx context.Context, where *WhereDto) ([]Category, error)
	GetManyByIds(ctx context.Context, ids []uuid.UUID) ([]*Category, error)
	UpdateOneById(ctx context.Context, id uuid.UUID, data UpdateDto) (*Category, error)
	DeleteOne(ctx context.Context, id uuid.UUID) error
}

type CategoryDataloader interface {
	ItemLoader(ctx context.Context, id uuid.UUID) (*Category, error)
}

// Entity & Enum
//

type Category struct {
	Id        uuid.UUID      `json:"id" bson:"_id"`
	Name      string         `json:"name" bson:"name"`
	Slug      string         `json:"slug" bson:"slug"`
	Status    CategoryStatus `json:"status" bson:"status"`
	CreatedAt time.Time      `json:"createdAt" bson:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt" bson:"updatedAt"`
}

type CategoryStatus string

const (
	StatusActive   CategoryStatus = "ACTIVE"
	StatusPending  CategoryStatus = "PENDING"
	StatusArchived CategoryStatus = "ARCHIVED"
)

// Dto
//

type CreateDto struct {
	Name   string         `json:"name" bson:"name"`
	Slug   *string        `json:"slug" bson:"slug"`
	Status CategoryStatus `json:"status" bson:"status"`
}

type WhereDto struct {
	Name       *filter.StringFilter     `json:"name" bson:"name"`
	Slug       *filter.StringFilter     `json:"slug" bson:"slug"`
	Status     *CategoryStatus          `json:"status" bson:"status"`
	Sort       *filter.SortFilter       `json:"sort" bson:"sort"`
	Pagination *filter.PaginationFilter `json:"pagination" bson:"pagination"`
}

type UpdateDto struct {
	Name   *string         `json:"name" bson:"name"`
	Slug   *string         `json:"slug" bson:"slug"`
	Status *CategoryStatus `json:"status" bson:"status,omitempty"`
}
