package category_test

import (
	"context"
	"testing"
	"time"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"github.com/kazmerdome/best-ever-golang-starter/internal/util/filter"
	"github.com/kazmerdome/best-ever-golang-starter/mocks"
	"github.com/stretchr/testify/assert"
)

type serviceFixture struct {
	service category.CategoryService
	mocks   struct {
		repository *mocks.CategoryRepository
		dataloader *mocks.CategoryDataloader
	}
	data struct {
		ctx       context.Context
		category  category.Category
		createDto category.CreateDto
		updateDto category.UpdateDto
		whereDto  category.WhereDto
	}
}

func newServiceFixture(t *testing.T) *serviceFixture {
	f := &serviceFixture{}
	f.mocks.repository = mocks.NewCategoryRepository(t)
	f.mocks.dataloader = mocks.NewCategoryDataloader(t)
	f.data.category = category.Category{
		Id:        uuid.New(),
		Name:      faker.BeerName(),
		Slug:      slug.Make(f.data.createDto.Name),
		Status:    category.StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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
	f.data.whereDto = category.WhereDto{
		Pagination: &filter.PaginationFilter{Limit: 1, Skip: 0},
	}
	f.service = category.NewCategoryService(f.mocks.repository, f.mocks.dataloader)
	return f
}

// Crud
//

func TestCreateCategory(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.repository.EXPECT().CreateOne(f.data.ctx, f.data.createDto).Return(&f.data.category, nil)
	c, err := f.service.CreateCategory(f.data.ctx, f.data.createDto)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.category)
}

func TestGetCategory(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.repository.EXPECT().GetOneById(f.data.ctx, f.data.category.Id).Return(&f.data.category, nil)
	c, err := f.service.GetCategory(f.data.ctx, f.data.category.Id)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.category)
}

func TestListCategories(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.repository.EXPECT().GetMany(f.data.ctx, &f.data.whereDto).Return([]category.Category{f.data.category}, nil)
	cs, err := f.service.ListCategories(f.data.ctx, &f.data.whereDto)
	assert.NoError(t, err)
	for _, c := range cs {
		assert.Equal(t, c, f.data.category)
	}
}

func TestUpdateCategory(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.repository.EXPECT().UpdateOneById(f.data.ctx, f.data.category.Id, f.data.updateDto).Return(&f.data.category, nil)
	c, err := f.service.UpdateCategory(f.data.ctx, f.data.category.Id, f.data.updateDto)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.category)
}

func TestDeleteCategory(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.repository.EXPECT().DeleteOne(f.data.ctx, f.data.category.Id).Return(nil)
	err := f.service.DeleteCategory(f.data.ctx, f.data.category.Id)
	assert.NoError(t, err)
}

// Dataloader
func TestLoadCategory(t *testing.T) {
	f := newServiceFixture(t)
	f.mocks.dataloader.EXPECT().ItemLoader(f.data.ctx, f.data.category.Id).Return(&f.data.category, nil)
	c, err := f.service.LoadCategory(f.data.ctx, f.data.category.Id)
	assert.NoError(t, err)
	assert.Equal(t, c, &f.data.category)
}
