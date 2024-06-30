package repository

import (
	"7day/webook/internal/repository/cache"
	"7day/webook/pkg/logger"
	"context"
)

type CodeRepo interface {
	Store(ctx context.Context, key, val string) error
	Verify(ctx context.Context, key string, val string) error
}

func NewCodeRepoImpl(l logger.Logger, cache cache.CodeCache) CodeRepo {
	return &CodeRepoImpl{l: l, cache: cache}
}

type CodeRepoImpl struct {
	l     logger.Logger
	cache cache.CodeCache
}

func (repo CodeRepoImpl) Verify(ctx context.Context, key string, val string) error {
	return repo.cache.Verify(ctx, key, val)
}

func (repo CodeRepoImpl) Store(ctx context.Context, key, val string) error {
	return repo.cache.Set(ctx, key, val)
}
