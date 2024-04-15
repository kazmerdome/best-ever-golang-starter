package post

import (
	db "gitlab.com/kazmerdome/best-ever-golang-starter/internal/actor/db/sql"
	postQuerier "gitlab.com/kazmerdome/best-ever-golang-starter/internal/module/post/post-querier"
)

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
