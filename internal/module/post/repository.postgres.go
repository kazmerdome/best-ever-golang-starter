package post

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
	postQuerier "github.com/kazmerdome/best-ever-golang-starter/internal/module/post/post-querier"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var ErrPostNotFound = fmt.Errorf("post is not found")

type postgresPostRepository struct {
	querier postQuerier.Querier
	logger  zerolog.Logger
}

func NewPostgresPostRepository(querier postQuerier.Querier) *postgresPostRepository {
	return &postgresPostRepository{
		querier: querier,
		logger: log.
			With().
			Str("module", "post").
			Str("provider", "repository").
			Logger(),
	}
}

func (r *postgresPostRepository) CreateOne(ctx context.Context, data CreateDto) (*Post, error) {
	s := data.Title
	if data.Slug != nil {
		s = *data.Slug
	}
	s = slug.Make(s)
	params := postQuerier.CreateOneParams{
		Title:    data.Title,
		Slug:     s,
		Category: data.Category,
		Status:   string(data.Status),
		Content:  sql.NullString{String: "", Valid: false},
	}
	if data.Content != nil {
		params.Content = sql.NullString{String: *data.Content, Valid: true}
	}
	qm, err := r.querier.CreateOne(ctx, params)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "CreateOne").
			Str("category", "call querier.CreateOne").
			Send()
		return nil, err
	}
	return r.buildModelFromQuerier(qm)
}

func (r *postgresPostRepository) GetOneById(ctx context.Context, id uuid.UUID) (*Post, error) {
	model, err := r.querier.GetOneById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPostNotFound
		}
		r.logger.
			Error().
			Err(err).
			Str("method", "getOneById").
			Str("event", fmt.Sprintf("failed to get post: %v", err)).
			Send()
		return nil, err
	}
	return r.buildModelFromQuerier(model)
}

func (r *postgresPostRepository) GetMany(ctx context.Context, where *WhereDto) ([]Post, error) {
	querierParams := postQuerier.GetManyParams{}
	if where != nil {
		if where.Title != nil {
			if where.Title.Eq != nil {
				querierParams.TitleEq = sql.NullString{Valid: true, String: *where.Title.Eq}
			}
			if where.Title.Regex != nil {
				querierParams.TitleRegex = sql.NullString{Valid: true, String: *where.Title.Regex}
			}
		}
		if where.Slug != nil {
			if where.Slug.Eq != nil {
				querierParams.SlugEq = sql.NullString{Valid: true, String: *where.Slug.Eq}
			}
			if where.Slug.Regex != nil {
				querierParams.SlugRegex = sql.NullString{Valid: true, String: *where.Slug.Regex}
			}
		}
		if where.Category != nil {
			if where.Category.Eq != nil {
				querierParams.CategoryEq = uuid.NullUUID{UUID: *where.Category.Eq, Valid: true}
			}
			if len(where.Category.In) > 0 {
				querierParams.CategoryIn = where.Category.In
			}
		}
		if where.Status != nil {
			status := string(*where.Status)
			querierParams.Status = sql.NullString{String: status, Valid: true}
		}
		orderBy := "created_at"
		sortOrder := "desc"
		if where.Sort != nil {
			orderBy = where.Sort.SortBy
			if where.Sort.SortOrder != nil {
				sortOrder = string(*where.Sort.SortOrder)
			}
		}
		querierParams.SortQuery = sql.NullString{Valid: true, String: fmt.Sprintf("%s__%s", orderBy, sortOrder)}
		if where.Pagination != nil {
			querierParams.Limit = sql.NullInt32{Valid: true, Int32: int32(where.Pagination.Limit)}
			querierParams.Offset = sql.NullInt32{Valid: true, Int32: int32(where.Pagination.Skip)}
		}
	}
	models, err := r.querier.GetMany(ctx, querierParams)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "GetMany").
			Str("category", "call querier.GetMany").
			Send()
		return nil, err
	}
	items := make([]Post, len(models))
	for i, model := range models {
		qi, _ := r.buildModelFromQuerier(model)
		items[i] = *qi
	}
	return items, err
}

func (r *postgresPostRepository) GetManyByIds(ctx context.Context, ids []uuid.UUID) ([]*Post, error) {
	models, err := r.querier.GetManyByIds(ctx, ids)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "GetManyByIds").
			Str("category", "call querier.GetManyByIds").
			Send()
		return nil, err
	}
	items := make([]*Post, len(models))
	for i, model := range models {
		qi, _ := r.buildModelFromQuerier(model)
		items[i] = qi
	}
	return items, err
}

func (r *postgresPostRepository) UpdateOneById(ctx context.Context, id uuid.UUID, data UpdateDto) (*Post, error) {
	params := postQuerier.UpdateOneByIdParams{}
	params.ID = id
	if data.Title != nil {
		params.Title = sql.NullString{Valid: true, String: *data.Title}
	}
	if data.Slug != nil {
		params.Slug = sql.NullString{Valid: true, String: *data.Slug}
	}
	if data.Category != nil {
		params.Category = uuid.NullUUID{Valid: true, UUID: *data.Category}
	}
	if data.Status != nil {
		params.Status = sql.NullString{Valid: true, String: string(*data.Status)}
	}
	if data.Content != nil {
		params.Content = sql.NullString{Valid: true, String: *data.Content}
	}
	qData, err := r.querier.UpdateOneById(ctx, params)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "UpdateOneById").
			Str("category", "call querier.UpdateOneById").
			Send()
		return nil, err
	}
	return r.buildModelFromQuerier(qData)
}

func (r *postgresPostRepository) DeleteOne(ctx context.Context, id uuid.UUID) error {
	return r.querier.DeleteOne(ctx, id)
}

func (r *postgresPostRepository) buildModelFromQuerier(qm postQuerier.Post) (*Post, error) {
	return &Post{
		Id:        qm.ID,
		Title:     qm.Title,
		Slug:      qm.Slug,
		Category:  qm.Category,
		Status:    PostStatus(qm.Status),
		Content:   qm.Content.String,
		CreatedAt: qm.CreatedAt.Time,
		UpdatedAt: qm.UpdatedAt.Time,
	}, nil
}
