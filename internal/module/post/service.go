package post

import (
	"context"

	"github.com/google/uuid"
)

type postService struct {
	repository PostRepository
	dataloader PostDataloader
}

func NewPostService(repository PostRepository, dataloader PostDataloader) *postService {
	return &postService{
		repository: repository,
		dataloader: dataloader,
	}
}

// Crud
//

func (s *postService) CreatePost(ctx context.Context, data CreateDto) (*Post, error) {
	return s.repository.CreateOne(ctx, data)
}
func (s *postService) GetPost(ctx context.Context, id uuid.UUID) (*Post, error) {
	return s.repository.GetOneById(ctx, id)
}
func (s *postService) ListPosts(ctx context.Context, where *WhereDto) ([]Post, error) {
	return s.repository.GetMany(ctx, where)
}
func (s *postService) UpdatePost(ctx context.Context, id uuid.UUID, data UpdateDto) (*Post, error) {
	return s.repository.UpdateOneById(ctx, id, data)
}
func (s *postService) DeletePost(ctx context.Context, id uuid.UUID) error {
	return s.repository.DeleteOne(ctx, id)
}

// Dataloader
//

func (r *postService) LoadPost(ctx context.Context, id uuid.UUID) (*Post, error) {
	return r.dataloader.ItemLoader(ctx, id)
}
