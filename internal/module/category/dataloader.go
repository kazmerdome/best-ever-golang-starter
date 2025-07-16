package category

import (
	"context"

	"github.com/google/uuid"
	"github.com/graph-gophers/dataloader"
	dl "github.com/kazmerdome/best-ever-golang-starter/internal/util/dataloader"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type categoryDataloader struct {
	repository CategoryRepository
	itemLoader *dataloader.Loader
	logger     zerolog.Logger
}

func NewCategoryDataloader(repository CategoryRepository) *categoryDataloader {
	loader := &categoryDataloader{
		repository: repository,
		logger: log.
			With().
			Str("module", "category").
			Str("provider", "dataloader").
			Logger(),
	}
	loader.itemLoader = dataloader.NewBatchedLoader(
		loader.batchItemLoader,
		dataloader.WithCache(&dataloader.NoCache{}),
	)
	return loader
}

func (r *categoryDataloader) ItemLoader(ctx context.Context, id uuid.UUID) (*Category, error) {
	thunk := r.itemLoader.Load(ctx, dl.UuidKey(id))
	result, err := thunk()
	if err != nil {
		r.logger.
			Debug().
			Err(err).
			Str("method", "ItemLoader").
			Str("category", "call itemLoader.Load.thunk").
			Str("id", id.String()).
			Send()
		return nil, err
	}
	org, ok := result.(*Category)
	if !ok {
		return nil, nil
	}
	return org, nil
}

func (r *categoryDataloader) batchItemLoader(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// convert keys to uuids and create placeholders
	uuids := make([]uuid.UUID, len(keys))
	bucket := make(map[uuid.UUID]*dataloader.Result, len(keys))
	for i, key := range keys {
		uid := key.Raw().(uuid.UUID)
		uuids[i] = uid
		bucket[uid] = &dataloader.Result{Data: nil, Error: nil}
	}
	// call repository and add the values to the bucket
	orgs, err := r.repository.GetManyByIds(ctx, uuids)
	if err != nil {
		r.logger.
			Debug().
			Err(err).
			Str("method", "batchItemLoader").
			Str("category", "repository.GetManyByIds call is failed").
			Send()

		return []*dataloader.Result{{Data: nil, Error: err}}
	}
	for _, org := range orgs {
		if org != nil {
			bucket[org.Id] = &dataloader.Result{Data: org, Error: nil}
		}
	}
	// create result array
	results := make([]*dataloader.Result, len(keys))
	for i, key := range keys {
		uid := key.Raw().(uuid.UUID)
		results[i] = bucket[uid]
	}
	return results
}
