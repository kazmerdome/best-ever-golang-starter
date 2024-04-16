package post_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post"
	postQuerier "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post/post-querier"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter"
	"gitlab.com/kazmerdome/best-ever-golang-starter/mocks"
)

type repositoryFixture struct {
	repository post.PostRepository
	mocks      struct {
		postQuerier *mocks.PostQuerier
	}
	data struct {
		ctx      context.Context
		category category.Category

		post      post.Post
		createDto post.CreateDto
		updateDto post.UpdateDto
		whereDto  post.WhereDto

		querierPost                postQuerier.Post
		querierCreateOneParams     postQuerier.CreateOneParams
		querierGetManyParams       postQuerier.GetManyParams
		querierUpdateOneByIdParams postQuerier.UpdateOneByIdParams
	}
}

func newRepositoryFixture(t *testing.T) *repositoryFixture {
	t.Parallel()
	f := &repositoryFixture{}
	f.mocks.postQuerier = mocks.NewPostQuerier(t)
	f.data.ctx = context.Background()
	f.data.category = category.Category{
		Id:        uuid.New(),
		Name:      faker.BeerName(),
		Slug:      slug.Make(f.data.category.Name),
		Status:    category.StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	title := faker.BeerName()
	f.data.post = post.Post{
		Id:        uuid.New(),
		Title:     title,
		Slug:      slug.Make(title),
		Category:  f.data.category.Id,
		Status:    post.StatusActive,
		Content:   faker.Sentence(200),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	f.data.createDto = post.CreateDto{
		Title:    f.data.post.Title,
		Slug:     &f.data.post.Slug,
		Category: f.data.category.Id,
		Status:   f.data.post.Status,
		Content:  &f.data.post.Content,
	}
	f.data.updateDto = post.UpdateDto{
		Title:    &f.data.post.Title,
		Slug:     &f.data.post.Slug,
		Category: &f.data.category.Id,
		Status:   &f.data.post.Status,
		Content:  &f.data.post.Content,
	}
	var SortOrder filter.SortOrder = "desc"
	f.data.whereDto = post.WhereDto{
		Title:      &filter.StringFilter{Eq: &f.data.post.Title, Regex: &f.data.post.Title},
		Slug:       &filter.StringFilter{Eq: &f.data.post.Slug, Regex: &f.data.post.Slug},
		Category:   &filter.UuidFilter{Eq: &f.data.post.Category, In: []uuid.UUID{f.data.post.Category}},
		Status:     &f.data.post.Status,
		Sort:       &filter.SortFilter{SortBy: "createdAt", SortOrder: &SortOrder},
		Pagination: &filter.PaginationFilter{Limit: 1, Skip: 0},
	}

	f.data.querierPost = postQuerier.Post{
		ID:        f.data.post.Id,
		Title:     f.data.post.Title,
		Slug:      f.data.post.Slug,
		Category:  f.data.post.Category,
		Status:    string(f.data.post.Status),
		Content:   sql.NullString{Valid: true, String: f.data.post.Content},
		CreatedAt: sql.NullTime{Valid: true, Time: f.data.post.CreatedAt},
		UpdatedAt: sql.NullTime{Valid: true, Time: f.data.post.UpdatedAt},
	}
	f.data.querierCreateOneParams = postQuerier.CreateOneParams{
		Title:    f.data.createDto.Title,
		Slug:     *f.data.createDto.Slug,
		Category: f.data.createDto.Category,
		Status:   string(f.data.createDto.Status),
		Content:  sql.NullString{Valid: true, String: *f.data.createDto.Content},
	}
	f.data.querierGetManyParams = postQuerier.GetManyParams{
		TitleEq:    sql.NullString{Valid: true, String: f.data.post.Title},
		TitleRegex: sql.NullString{Valid: true, String: f.data.post.Title},
		SlugEq:     sql.NullString{Valid: true, String: f.data.post.Slug},
		SlugRegex:  sql.NullString{Valid: true, String: f.data.post.Slug},
		CategoryEq: uuid.NullUUID{Valid: true, UUID: f.data.post.Category},
		CategoryIn: []uuid.UUID{f.data.post.Category},
		Status:     sql.NullString{String: string(f.data.category.Status), Valid: true},
		SortQuery:  sql.NullString{Valid: true, String: fmt.Sprintf("%s__%s", f.data.whereDto.Sort.SortBy, string(SortOrder))},
		Offset:     sql.NullInt32{Valid: true, Int32: int32(f.data.whereDto.Pagination.Skip)},
		Limit:      sql.NullInt32{Valid: true, Int32: int32(f.data.whereDto.Pagination.Limit)},
	}
	f.data.querierUpdateOneByIdParams = postQuerier.UpdateOneByIdParams{
		Title:    sql.NullString{Valid: true, String: *f.data.updateDto.Title},
		Slug:     sql.NullString{Valid: true, String: *f.data.updateDto.Slug},
		Category: uuid.NullUUID{Valid: true, UUID: *f.data.updateDto.Category},
		Status:   sql.NullString{Valid: true, String: string(*f.data.updateDto.Status)},
		Content:  sql.NullString{Valid: true, String: *f.data.updateDto.Content},
		ID:       f.data.post.Id,
	}
	f.repository = post.NewPostgresPostRepository(f.mocks.postQuerier)
	return f
}

// CreateOne
//

func TestCreateOne_FailsOn_QuerierCreateOne(t *testing.T) {
	f := newRepositoryFixture(t)
	s := "not a valid slug for sure"
	f.data.createDto.Slug = &s
	f.data.querierCreateOneParams.Slug = slug.Make(s)

	f.mocks.postQuerier.EXPECT().CreateOne(f.data.ctx, f.data.querierCreateOneParams).
		Return(f.data.querierPost, fmt.Errorf("QuerierCreateOne error"))

	c, err := f.repository.CreateOne(f.data.ctx, f.data.createDto)
	assert.EqualError(t, err, "QuerierCreateOne error")
	assert.Nil(t, c)
}

func TestCreateOne_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().CreateOne(f.data.ctx, f.data.querierCreateOneParams).
		Return(f.data.querierPost, nil)
	c, err := f.repository.CreateOne(f.data.ctx, f.data.createDto)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.post)
}

// GetOneById
//

func TestGetOneById_FailsOn_QuerierGetOneById(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().GetOneById(f.data.ctx, f.data.post.Id).
		Return(f.data.querierPost, fmt.Errorf("QuerierGetOneById error"))
	c, err := f.repository.GetOneById(f.data.ctx, f.data.post.Id)
	assert.EqualError(t, err, "QuerierGetOneById error")
	assert.Nil(t, c)
}

func TestGetOneById_FailsOn_NotFound(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().GetOneById(f.data.ctx, f.data.post.Id).
		Return(f.data.querierPost, sql.ErrNoRows)
	c, err := f.repository.GetOneById(f.data.ctx, f.data.post.Id)
	assert.EqualError(t, err, post.ErrPostNotFound.Error())
	assert.Nil(t, c)
}

func TestGetOneById_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().GetOneById(f.data.ctx, f.data.post.Id).
		Return(f.data.querierPost, nil)
	c, err := f.repository.GetOneById(f.data.ctx, f.data.post.Id)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.post)
}

// GetMany
//

func TestGetMany_FailsOn_QuerierGetMany(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().GetMany(f.data.ctx, f.data.querierGetManyParams).
		Return([]postQuerier.Post{f.data.querierPost}, fmt.Errorf("QuerierGetMany error"))
	c, err := f.repository.GetMany(f.data.ctx, &f.data.whereDto)
	assert.EqualError(t, err, "QuerierGetMany error")
	assert.Nil(t, c)
}

func TestGetMany_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().GetMany(f.data.ctx, f.data.querierGetManyParams).
		Return([]postQuerier.Post{f.data.querierPost}, nil)
	cs, err := f.repository.GetMany(f.data.ctx, &f.data.whereDto)
	assert.NoError(t, err)
	for _, c := range cs {
		assert.Equal(t, c, f.data.post)
	}
}

// GetManyByIds
//

func TestGetManyByIds_FailsOn_QuerierGetManyByIds(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().GetManyByIds(f.data.ctx, []uuid.UUID{f.data.post.Id}).
		Return([]postQuerier.Post{f.data.querierPost}, fmt.Errorf("QuerierGetManyByIds error"))
	c, err := f.repository.GetManyByIds(f.data.ctx, []uuid.UUID{f.data.post.Id})
	assert.EqualError(t, err, "QuerierGetManyByIds error")
	assert.Nil(t, c)
}

func TestGetManyByIds_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().GetManyByIds(f.data.ctx, []uuid.UUID{f.data.post.Id}).
		Return([]postQuerier.Post{f.data.querierPost}, nil)
	cs, err := f.repository.GetManyByIds(f.data.ctx, []uuid.UUID{f.data.post.Id})
	assert.NoError(t, err)
	for _, c := range cs {
		assert.Equal(t, c, &f.data.post)
	}
}

// UpdateOneById
//

func TestUpdateOneById_FailsOn_QuerierUpdateOneById(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().UpdateOneById(f.data.ctx, f.data.querierUpdateOneByIdParams).
		Return(f.data.querierPost, fmt.Errorf("QuerierUpdateOneById error"))
	c, err := f.repository.UpdateOneById(f.data.ctx, f.data.post.Id, f.data.updateDto)
	assert.EqualError(t, err, "QuerierUpdateOneById error")
	assert.Nil(t, c)
}

func TestUpdateOneById_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().UpdateOneById(f.data.ctx, f.data.querierUpdateOneByIdParams).
		Return(f.data.querierPost, nil)
	c, err := f.repository.UpdateOneById(f.data.ctx, f.data.post.Id, f.data.updateDto)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.post)
}

// DeleteOne
//

func TestDeleteOne_Success(t *testing.T) {
	f := newRepositoryFixture(t)
	f.mocks.postQuerier.EXPECT().DeleteOne(f.data.ctx, f.data.post.Id).
		Return(nil)
	err := f.repository.DeleteOne(f.data.ctx, f.data.post.Id)
	assert.NoError(t, err)
}
