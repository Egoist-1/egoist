package article

import (
	"context"
)

type ArticleDao interface {
	Upset(ctx context.Context, article Article) (Article, error)
	Sync(ctx context.Context, art Article) (int64, error)
	FindById(ctx context.Context, id int64) (Article, error)
	SyncStatus(ctx context.Context, art Article) error
	OnePage(ctx context.Context, uid int64, offset int, limit int) ([]Article, error)
	GetById(ctx context.Context, id int64)(Article,error)
}
