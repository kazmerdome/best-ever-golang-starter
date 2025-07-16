package post_test

import (
	"testing"

	"github.com/kazmerdome/best-ever-golang-starter/internal/module/post"
	"github.com/kazmerdome/best-ever-golang-starter/mocks"
	"github.com/stretchr/testify/assert"
)

type moduleFixture struct {
	mocks struct {
		db    *mocks.DB
		sqlDb *mocks.SqlDB
	}
}

func newModuleFixture(t *testing.T) *moduleFixture {
	f := &moduleFixture{}
	f.mocks.db = mocks.NewDB(t)
	f.mocks.sqlDb = mocks.NewSqlDB(t)
	return f
}

func TestModule(t *testing.T) {
	f := newModuleFixture(t)
	f.mocks.db.EXPECT().GetDB().Return(f.mocks.sqlDb)
	postModule := post.NewPostModule(f.mocks.db)
	postService := postModule.GetService()
	assert.Implements(t, (*post.PostService)(nil), postService)
}
