package category

import "github.com/kazmerdome/best-ever-golang-starter/internal/actor/db/mongodb"

type categoryModule struct {
	service CategoryService
}

func NewCategoryModule(db mongodb.Database) *categoryModule {
	repository := NewMongodbCategoryRepository(db)
	loader := NewCategoryDataloader(repository)
	service := NewCategoryService(repository, loader)
	return &categoryModule{
		service: service,
	}
}

func (m *categoryModule) GetService() CategoryService {
	return m.service
}
