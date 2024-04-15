package category_test

import (
	"context"
	"testing"
	"time"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/actor/db/mongodb"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type repositoryFixture struct {
	repository category.CategoryRepository
	mocks      struct {
		mt *mtest.T
	}
	data struct {
		ctx           context.Context
		category      category.Category
		categoryBsonD bson.D
		createDto     category.CreateDto
		updateDto     category.UpdateDto
		whereDto      category.WhereDto
	}
}

func newRepositoryFixture(t *testing.T) *repositoryFixture {
	t.Parallel()
	f := &repositoryFixture{}
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	f.data.ctx = context.Background()
	f.data.category = category.Category{
		Id:        uuid.New(),
		Name:      faker.BeerName(),
		Slug:      slug.Make(f.data.createDto.Name),
		Status:    category.StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	f.data.categoryBsonD = bson.D{
		{Key: "_id", Value: f.data.category.Id},
		{Key: "name", Value: f.data.category.Name},
		{Key: "slug", Value: f.data.category.Slug},
		{Key: "status", Value: f.data.category.Status},
		{Key: "createdAt", Value: f.data.category.CreatedAt.UTC()},
		{Key: "updatedAt", Value: f.data.category.UpdatedAt.UTC()},
	}

	f.data.createDto = category.CreateDto{
		Name:   f.data.category.Name,
		Slug:   &f.data.category.Slug,
		Status: f.data.category.Status,
	}
	f.data.updateDto = category.UpdateDto{
		Name:   &f.data.category.Name,
		Slug:   &f.data.category.Slug,
		Status: &f.data.category.Status,
	}
	var SortOrder filter.SortOrder = "desc"
	f.data.whereDto = category.WhereDto{
		Name:       &filter.StringFilter{Eq: &f.data.category.Name},
		Slug:       &filter.StringFilter{Eq: &f.data.category.Slug},
		Status:     &f.data.category.Status,
		Pagination: &filter.PaginationFilter{Limit: 1, Skip: 0},
		Sort:       &filter.SortFilter{SortBy: "createdAt", SortOrder: &SortOrder},
	}
	mt.Run("test", func(mtcb *mtest.T) {
		f.mocks.mt = mtcb
		f.repository = category.NewMongodbCategoryRepository(mongodb.NewDatabase(mtcb.DB))
	})
	return f
}

// CreateOne
//

func TestCreateOne_FailsOn_CollectionInsertOne(t *testing.T) {
	f := newRepositoryFixture(t)
	s := "not a valid slug for sure"
	f.data.createDto.Slug = &s
	f.mocks.mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{Message: "an error"}))
	c, err := f.repository.CreateOne(f.data.ctx, f.data.createDto)
	assert.EqualError(t, err, "write exception: write errors: [an error]")
	assert.Nil(t, c)
}

func TestCreateOne_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	s := "not a valid slug for sure"
	f.data.createDto.Slug = &s
	f.mocks.mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, f.data.categoryBsonD))
	c, err := f.repository.CreateOne(f.data.ctx, f.data.createDto)
	assert.NoError(t, err)
	assert.NotNil(t, c.Id)
	assert.Equal(t, c.Name, f.data.createDto.Name)
	assert.Equal(t, c.Slug, slug.Make(s))
	assert.Equal(t, c.Status, f.data.createDto.Status)
	assert.WithinDuration(t, time.Now(), c.CreatedAt, 10*time.Second)
	assert.WithinDuration(t, time.Now(), c.UpdatedAt, 10*time.Second)
}

// GetOneById
//

func TestGetOneById_FailsOn_CollectionFindOne(t *testing.T) {
	f := newRepositoryFixture(t)
	c, err := f.repository.GetOneById(f.data.ctx, f.data.category.Id)
	assert.EqualError(t, err, "failed to get category: no responses remaining")
	assert.Nil(t, c)
}

func TestGetOneById_FailsOn_NoDocuments(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))
	c, err := f.repository.GetOneById(f.data.ctx, f.data.category.Id)
	assert.EqualError(t, err, "category is not found")
	assert.Nil(t, c)
}

func TestGetOneById_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, f.data.categoryBsonD))
	c, err := f.repository.GetOneById(f.data.ctx, f.data.category.Id)
	assert.NoError(t, err)
	assert.Equal(t, c.Id, f.data.category.Id)
	assert.Equal(t, c.Name, f.data.category.Name)
	assert.Equal(t, c.Slug, f.data.category.Slug)
	assert.Equal(t, c.Status, f.data.category.Status)
	assert.WithinDuration(t, c.CreatedAt, f.data.category.CreatedAt, 1*time.Millisecond)
	assert.WithinDuration(t, c.UpdatedAt, f.data.category.UpdatedAt, 1*time.Millisecond)
}

// GetMany
//

func TestGetMany_FailsOn_CollectionFind(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{Message: "find error"}))
	cs, err := f.repository.GetMany(f.data.ctx, &f.data.whereDto)
	assert.EqualError(t, err, "write command error: [{write errors: [{find error}]}, {<nil>}]")
	assert.Empty(t, cs)
}

func TestGetMany_FailsOn_Cursor(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(
		// Indicate a Cursor Error
		mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{{Key: "ló", Value: "fasz"}}),
	)
	cs, err := f.repository.GetMany(f.data.ctx, &f.data.whereDto)
	assert.EqualError(t, err, "no responses remaining")
	assert.Empty(t, cs)
}

func TestGetMany_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(
		mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, f.data.categoryBsonD),
		mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch),
	)
	cs, err := f.repository.GetMany(f.data.ctx, &f.data.whereDto)
	assert.NoError(t, err)
	for _, c := range cs {
		assert.Equal(t, c.Id, f.data.category.Id)
		assert.Equal(t, c.Name, f.data.category.Name)
		assert.Equal(t, c.Slug, f.data.category.Slug)
		assert.Equal(t, c.Status, f.data.category.Status)
		assert.WithinDuration(t, c.CreatedAt, f.data.category.CreatedAt, 1*time.Millisecond)
		assert.WithinDuration(t, c.UpdatedAt, f.data.category.UpdatedAt, 1*time.Millisecond)
	}
}

// GetManyByIds
//

func TestGetManyByIds_FailsOn_CollectionFind(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{Message: "find error"}))
	cs, err := f.repository.GetManyByIds(f.data.ctx, []uuid.UUID{f.data.category.Id})
	assert.EqualError(t, err, "write command error: [{write errors: [{find error}]}, {<nil>}]")
	assert.Empty(t, cs)
}

func TestGetManyByIds_FailsOn_Cursor(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(
		// Indicate a Cursor Error
		mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{{Key: "ló", Value: "fasz"}}),
	)
	cs, err := f.repository.GetManyByIds(f.data.ctx, []uuid.UUID{f.data.category.Id})
	assert.EqualError(t, err, "no responses remaining")
	assert.Empty(t, cs)
}

func TestGetManyByIds_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(
		mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, f.data.categoryBsonD),
		mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch),
	)
	cs, err := f.repository.GetManyByIds(f.data.ctx, []uuid.UUID{f.data.category.Id})
	assert.NoError(t, err)
	for _, c := range cs {
		assert.Equal(t, c.Id, f.data.category.Id)
		assert.Equal(t, c.Name, f.data.category.Name)
		assert.Equal(t, c.Slug, f.data.category.Slug)
		assert.Equal(t, c.Status, f.data.category.Status)
		assert.WithinDuration(t, c.CreatedAt, f.data.category.CreatedAt, 1*time.Millisecond)
		assert.WithinDuration(t, c.UpdatedAt, f.data.category.UpdatedAt, 1*time.Millisecond)
	}
}

// UpdateOneById
//

func TestUpdateOneById_FailsOn_UpdateOne(t *testing.T) {
	f := newRepositoryFixture(t)
	c, err := f.repository.UpdateOneById(f.data.ctx, f.data.category.Id, f.data.updateDto)
	assert.EqualError(t, err, "no responses remaining")
	assert.Nil(t, c)
}

func TestUpdateOneById_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(
		mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, f.data.categoryBsonD),
		mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch, f.data.categoryBsonD),
	)
	c, err := f.repository.UpdateOneById(f.data.ctx, f.data.category.Id, f.data.updateDto)
	assert.NoError(t, err)
	assert.NotNil(t, c.Id)
	assert.Equal(t, c.Name, f.data.category.Name)
	assert.Equal(t, c.Slug, f.data.category.Slug)
	assert.Equal(t, c.Status, f.data.category.Status)
	assert.WithinDuration(t, c.CreatedAt, f.data.category.CreatedAt, 1*time.Millisecond)
	assert.WithinDuration(t, c.UpdatedAt, f.data.category.UpdatedAt, 1*time.Millisecond)
}

// DeleteOne
//

func TestDeleteOne_FailsOn_CollectionDeleteOne(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(mtest.CreateWriteErrorsResponse(mtest.WriteError{Message: "delete error"}))
	err := f.repository.DeleteOne(f.data.ctx, f.data.category.Id)
	assert.EqualError(t, err, "write exception: write errors: [delete error]")
}
func TestDeleteOne_FailsOn_DeleteCount(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(mtest.CreateCursorResponse(0, "foo.bar", mtest.FirstBatch))
	err := f.repository.DeleteOne(f.data.ctx, f.data.category.Id)
	assert.EqualError(t, err, "category is not found")
}
func TestDeleteOne_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.mt.AddMockResponses(
		bson.D{{Key: "ok", Value: 1}, {Key: "acknowledged", Value: true}, {Key: "n", Value: 1}},
	)
	err := f.repository.DeleteOne(f.data.ctx, f.data.category.Id)
	assert.NoError(t, err)
}
