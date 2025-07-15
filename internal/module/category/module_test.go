package category_test

import (
	"testing"

	"github.com/kazmerdome/best-ever-golang-starter/internal/module/category"
	"github.com/kazmerdome/best-ever-golang-starter/mocks"
	"github.com/stretchr/testify/assert"
)

type moduleFixture struct {
	mocks struct {
		database   *mocks.Database
		collection *mocks.Collection
	}
}

func newModuleFixture(t *testing.T) *moduleFixture {
	f := &moduleFixture{}
	f.mocks.database = mocks.NewDatabase(t)
	f.mocks.collection = mocks.NewCollection(t)
	return f
}

func TestModule(t *testing.T) {
	f := newModuleFixture(t)
	f.mocks.database.EXPECT().Collection("category").Return(f.mocks.collection)
	categoryModule := category.NewCategoryModule(f.mocks.database)
	categoryService := categoryModule.GetService()
	assert.Implements(t, (*category.CategoryService)(nil), categoryService)
}
