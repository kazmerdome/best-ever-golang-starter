package post_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	faker "github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"github.com/kazmerdome/best-ever-golang-starter/internal/module/post"
	"github.com/kazmerdome/best-ever-golang-starter/mocks"
	"github.com/stretchr/testify/assert"
)

type dataloaderFixture struct {
	dataloader post.PostDataloader
	mocks      struct {
		repository *mocks.PostRepository
	}
	data struct {
		ctx      context.Context
		post     post.Post
		category category.Category
	}
}

func newDataloaderFixture(t *testing.T) *dataloaderFixture {
	t.Parallel()
	f := &dataloaderFixture{}
	f.mocks.repository = mocks.NewPostRepository(t)
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
		Category:  f.data.category.Id,
		Slug:      slug.Make(f.data.post.Title),
		Status:    post.StatusActive,
		Content:   faker.Sentence(300),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	f.dataloader = post.NewPostDataloader(f.mocks.repository)
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
	f.mocks.repository.EXPECT().GetManyByIds(f.data.ctx, []uuid.UUID{uuid1}).Return([]*post.Post{&f.data.post}, nil)
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
	f.mocks.repository.EXPECT().GetManyByIds(f.data.ctx, []uuid.UUID{f.data.post.Id}).Return([]*post.Post{&f.data.post}, nil)
	go func() {
		defer wg.Done()
		result, err := f.dataloader.ItemLoader(f.data.ctx, f.data.post.Id)
		assert.NoError(t, err)
		assert.Equal(t, f.data.post, *result)
	}()
	wg.Wait()
}
