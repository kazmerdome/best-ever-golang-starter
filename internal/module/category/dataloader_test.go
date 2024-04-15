package category_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"gitlab.com/kazmerdome/best-ever-golang-starter/mocks"
)

type dataloaderFixture struct {
	dataloader category.CategoryDataloader
	mocks      struct {
		repository *mocks.CategoryRepository
	}
	data struct {
		ctx       context.Context
		category  category.Category
		createDto category.CreateDto
		updateDto category.UpdateDto
		whereDto  category.WhereDto
	}
}

func newDataloaderFixture(t *testing.T) *dataloaderFixture {
	t.Parallel()
	f := &dataloaderFixture{}
	f.mocks.repository = mocks.NewCategoryRepository(t)
	f.data.category = category.Category{
		Id:        uuid.New(),
		Name:      faker.BeerName(),
		Slug:      slug.Make(f.data.createDto.Name),
		Status:    category.StatusActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	f.dataloader = category.NewCategoryDataloader(f.mocks.repository)
	return f
}

// ItemLoader
//

func TestItemLoader_FailsOn_Thunk(t *testing.T) {
	f := newDataloaderFixture(t)
	var wg sync.WaitGroup
	wg.Add(1)
	f.mocks.repository.EXPECT().GetManyByIds(f.data.ctx, []uuid.UUID{f.data.category.Id}).Return(nil, fmt.Errorf("error fetching category"))
	go func() {
		defer wg.Done()
		result, err := f.dataloader.ItemLoader(f.data.ctx, f.data.category.Id)
		assert.EqualError(t, err, "error fetching category")
		assert.Nil(t, result)
	}()
	wg.Wait()
}

func TestItemLoader_SilentFailsOn_CategoryCasting(t *testing.T) {
	f := newDataloaderFixture(t)
	var wg sync.WaitGroup
	wg.Add(1)
	uuid1 := uuid.New()
	f.mocks.repository.EXPECT().GetManyByIds(f.data.ctx, []uuid.UUID{uuid1}).Return([]*category.Category{&f.data.category}, nil)
	go func() {
		defer wg.Done()
		result, err := f.dataloader.ItemLoader(f.data.ctx, uuid1)
		assert.Nil(t, err)
		assert.Nil(t, result)
	}()
	wg.Wait()
}

func TestItemLoader_Success(t *testing.T) {
	f := newDataloaderFixture(t)
	var wg sync.WaitGroup
	wg.Add(1)
	f.mocks.repository.EXPECT().GetManyByIds(f.data.ctx, []uuid.UUID{f.data.category.Id}).Return([]*category.Category{&f.data.category}, nil)
	go func() {
		defer wg.Done()
		result, err := f.dataloader.ItemLoader(f.data.ctx, f.data.category.Id)
		assert.NoError(t, err)
		assert.Equal(t, f.data.category, *result)
	}()
	wg.Wait()
}
