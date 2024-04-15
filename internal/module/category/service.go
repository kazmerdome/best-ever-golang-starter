package category

import (
	"context"

	"github.com/google/uuid"
)

type categoryService struct {
	repository CategoryRepository
	dataloader CategoryDataloader
}

func NewCategoryService(repository CategoryRepository, dataloader CategoryDataloader) *categoryService {
	return &categoryService{
		repository: repository,
		dataloader: dataloader,
	}
}

// Crud
//

func (s *categoryService) CreateCategory(ctx context.Context, data CreateDto) (*Category, error) {
	return s.repository.CreateOne(ctx, data)
}
func (s *categoryService) GetCategory(ctx context.Context, id uuid.UUID) (*Category, error) {
	return s.repository.GetOneById(ctx, id)
}
func (s *categoryService) ListCategories(ctx context.Context, where *WhereDto) ([]Category, error) {
	return s.repository.GetMany(ctx, where)
}
func (s *categoryService) UpdateCategory(ctx context.Context, id uuid.UUID, data UpdateDto) (*Category, error) {
	return s.repository.UpdateOneById(ctx, id, data)
}
func (s *categoryService) DeleteCategory(ctx context.Context, id uuid.UUID) error {
	return s.repository.DeleteOne(ctx, id)
}

// Dataloader
//

func (r *categoryService) LoadCategory(ctx context.Context, id uuid.UUID) (*Category, error) {
	return r.dataloader.ItemLoader(ctx, id)
}
