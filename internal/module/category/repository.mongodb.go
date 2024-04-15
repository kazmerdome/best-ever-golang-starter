package category

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/actor/db/mongodb"

	"github.com/gosimple/slug"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ErrCategoryNotFound = errors.New("category is not found")

type mongodbCategoryRepository struct {
	collection mongodb.Collection
	logger     zerolog.Logger
}

func NewMongodbCategoryRepository(db mongodb.Database) *mongodbCategoryRepository {
	return &mongodbCategoryRepository{
		collection: db.Collection("category"),
		logger: log.
			With().
			Str("module", "category").
			Str("provider", "repository").
			Logger(),
	}
}

func (r *mongodbCategoryRepository) CreateOne(ctx context.Context, data CreateDto) (*Category, error) {
	category := &Category{
		Id:        uuid.New(),
		Name:      data.Name,
		Slug:      data.Name,
		Status:    data.Status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if data.Slug != nil {
		category.Slug = *data.Slug
	}
	if !slug.IsSlug(category.Slug) {
		s := slug.Make(category.Slug)
		category.Slug = s
	}
	_, err := r.collection.InsertOne(ctx, category)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "CreateOne").
			Str("category", "call collection.InsertOne").
			Send()
		return nil, err
	}
	return category, nil
}

func (r *mongodbCategoryRepository) GetOneById(ctx context.Context, id uuid.UUID) (*Category, error) {
	filter := bson.M{"_id": id}
	var category Category
	err := r.collection.FindOne(ctx, filter).Decode(&category)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, ErrCategoryNotFound
		}
		r.logger.
			Error().
			Err(err).
			Str("method", "GetOneById").
			Str("category", "call collection.FindOne").
			Send()
		return nil, fmt.Errorf("failed to get category: %v", err)
	}
	return &category, nil
}

func (r *mongodbCategoryRepository) GetMany(ctx context.Context, where *WhereDto) ([]Category, error) {
	// Filter
	filter := bson.M{}
	options := options.Find()
	if where != nil {
		if where.Name != nil {
			filter["name"] = mongodb.GetStringFilter(where.Name)
		}
		if where.Slug != nil {
			filter["slug"] = mongodb.GetStringFilter(where.Slug)
		}
		if where.Status != nil {
			filter["status"] = mongodb.GetRegexFilter(string(*where.Status))
		}
		// Sort
		sort := bson.D{}
		if where.Sort != nil {
			sort = mongodb.GetSortFilter(where.Sort)
		}
		options.SetSort(sort)
		// Pagination
		if where.Pagination != nil {
			options.SetLimit(int64(where.Pagination.Limit))
			options.SetSkip(int64(where.Pagination.Skip))
		}
	}
	// Cursor
	cursor, err := r.collection.Find(ctx, filter, options)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "GetMany").
			Str("category", "call collection.Find").
			Send()
		return nil, err
	}
	defer cursor.Close(ctx)
	var categories []Category
	if err = cursor.All(ctx, &categories); err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "GetMany").
			Str("category", "call cursor.All").
			Send()
		return nil, err
	}
	return categories, nil
}

func (r *mongodbCategoryRepository) GetManyByIds(ctx context.Context, ids []uuid.UUID) ([]*Category, error) {
	filter := bson.M{"_id": bson.M{"$in": ids}}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "GetManyByIds").
			Str("category", "call collection.Find").
			Send()
		return nil, err
	}
	defer cursor.Close(ctx)
	var categories []*Category
	if err = cursor.All(ctx, &categories); err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "GetManyByIds").
			Str("category", "call cursor.All").
			Send()
		return nil, err
	}
	return categories, nil
}

func (r *mongodbCategoryRepository) UpdateOneById(ctx context.Context, id uuid.UUID, data UpdateDto) (*Category, error) {
	filter := bson.M{"_id": id}
	update := bson.M{"$set": data}

	opts := options.Update().SetUpsert(false)
	_, err := r.collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "UpdateOneById").
			Str("category", "call collection.UpdateOne").
			Send()
		return nil, err
	}
	return r.GetOneById(ctx, id)
}

func (r *mongodbCategoryRepository) DeleteOne(ctx context.Context, id uuid.UUID) error {
	filter := bson.M{"_id": id}
	deleteResult, err := r.collection.DeleteOne(ctx, filter)
	if err != nil {
		r.logger.
			Error().
			Err(err).
			Str("method", "DeleteOne").
			Str("category", "call collection.DeleteOne").
			Send()
		return err
	}
	if deleteResult.DeletedCount == 0 {
		r.logger.
			Debug().
			Err(err).
			Str("method", "DeleteOne").
			Str("category", "Item is not found in db").
			Send()
		return ErrCategoryNotFound
	}
	return nil
}
