package post_test

import (
	"context"
	"testing"
	"time"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/util/filter"
	"gitlab.com/kazmerdome/best-ever-golang-starter/mocks"
)

type serviceFixture struct {
	service post.PostService
	mocks   struct {
		repository *mocks.PostRepository
		dataloader *mocks.PostDataloader
	}
	data struct {
		ctx       context.Context
		category  category.Category
		post      post.Post
		createDto post.CreateDto
		updateDto post.UpdateDto
		whereDto  post.WhereDto
	}
}

func newServiceFixture(t *testing.T) *serviceFixture {
	f := &serviceFixture{}
	f.mocks.repository = mocks.NewPostRepository(t)
	f.mocks.dataloader = mocks.NewPostDataloader(t)
	f.data.category = category.Category{
		Id:        uuid.New(),
		Name:      faker.BeerName(),
		Slug:      slug.Make(f.data.category.Name),
		Status:    category.StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	f.data.post = post.Post{
		Id:        uuid.New(),
		Title:     faker.BeerName(),
		Slug:      slug.Make(f.data.post.Title),
		Category:  f.data.category.Id,
		Status:    post.StatusActive,
		Content:   faker.Sentence(300),
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
	f.data.whereDto = post.WhereDto{
		Pagination: &filter.PaginationFilter{Limit: 1, Skip: 0},
	}
	f.service = post.NewPostService(f.mocks.repository, f.mocks.dataloader)
	return f
}

// Crud
//

func TestCreatePost(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.repository.EXPECT().CreateOne(f.data.ctx, f.data.createDto).Return(&f.data.post, nil)
	c, err := f.service.CreatePost(f.data.ctx, f.data.createDto)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.post)
}

func TestGetPost(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.repository.EXPECT().GetOneById(f.data.ctx, f.data.post.Id).Return(&f.data.post, nil)
	c, err := f.service.GetPost(f.data.ctx, f.data.post.Id)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.post)
}

func TestListCategories(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.repository.EXPECT().GetMany(f.data.ctx, &f.data.whereDto).Return([]post.Post{f.data.post}, nil)
	cs, err := f.service.ListPosts(f.data.ctx, &f.data.whereDto)
	assert.NoError(t, err)
	for _, c := range cs {
		assert.Equal(t, c, f.data.post)
	}
}

func TestUpdatePost(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.repository.EXPECT().UpdateOneById(f.data.ctx, f.data.post.Id, f.data.updateDto).Return(&f.data.post, nil)
	c, err := f.service.UpdatePost(f.data.ctx, f.data.post.Id, f.data.updateDto)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.post)
}

func TestDeletePost(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.repository.EXPECT().DeleteOne(f.data.ctx, f.data.post.Id).Return(nil)
	err := f.service.DeletePost(f.data.ctx, f.data.post.Id)
	assert.NoError(t, err)
}

// Dataloader
func TestLoadPost(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.dataloader.EXPECT().ItemLoader(f.data.ctx, f.data.post.Id).Return(&f.data.post, nil)
	c, err := f.service.LoadPost(f.data.ctx, f.data.post.Id)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.post)
}
