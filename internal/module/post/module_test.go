package post_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post"
	"gitlab.com/kazmerdome/best-ever-golang-starter/mocks"
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
