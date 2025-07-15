package post

import (
	db "github.com/kazmerdome/best-ever-golang-starter/internal/actor/db/sql"
	postQuerier "github.com/kazmerdome/best-ever-golang-starter/internal/module/post/post-querier"
)

//go:generate make name=PostService mock
//go:generate make name=PostRepository mock
//go:generate make name=PostDataloader mock
//go:generate make name=Querier structname=PostQuerier filename=PostQuerier.go srcpkg=github.com/kazmerdome/best-ever-golang-starter/internal/module/post/post-querier mock

type postModule struct {
	service PostService
}

func NewPostModule(db db.DB) *postModule {
	querier := postQuerier.New(db.GetDB())
	repository := NewPostgresPostRepository(querier)
	loader := NewPostDataloader(repository)
	service := NewPostService(repository, loader)
	return &postModule{
		service: service,
	}
}

func (m *postModule) GetService() PostService {
	return m.service
}
