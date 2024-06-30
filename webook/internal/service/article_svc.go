package service

import (
	"7day/webook/internal/domain"
	"7day/webook/internal/repository"
	"context"
)

type ArticleSVC interface {
	Save(ctx context.Context, article domain.Article) (id int64, err error)
	Publish(ctx context.Context, art domain.Article) (int64, error)
	Withdraw(ctx context.Context, article domain.Article) error
	List(ctx context.Context, id int64, offset, limit int) ([]domain.Article, error)
	GetDetail(ctx context.Context, id int64)(domain.Article,error)
}

func NewArticleSVCImpl(rpo repository.ArticleRepo) ArticleSVC {
	return &ArticleSVCImpl{
		repo: rpo,
	}
}

type ArticleSVCImpl struct {
	repo repository.ArticleRepo
}

func (svc *ArticleSVCImpl) GetDetail(ctx context.Context, id int64) (domain.Article, error) {
	return svc.repo.GetDetail(ctx,id)
}

func (svc *ArticleSVCImpl) List(ctx context.Context, id int64, offset, limit int) ([]domain.Article, error) {
	return svc.repo.List(ctx, id, offset, limit)
}

func (svc *ArticleSVCImpl) Withdraw(ctx context.Context, article domain.Article) error {
	return svc.repo.SyncStatus(ctx, article)
}

func (svc *ArticleSVCImpl) Publish(ctx context.Context, art domain.Article) (int64, error) {
	art.Status = 1
	return svc.repo.Sync(ctx, art)

}

func (svc *ArticleSVCImpl) Save(ctx context.Context, article domain.Article) (id int64, err error) {
	return svc.repo.Upset(ctx, article)
}
