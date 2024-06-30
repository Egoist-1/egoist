package cache

import (
	"7day/webook/internal/domain"
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"time"
)

type ArticleCache interface {
	GetFirstPage(ctx context.Context, key string) ([]domain.Article, error)
	SetFirstPage(ctx context.Context, key string, val []domain.Article) error
	SetArtDetail(ctx context.Context, key string, val domain.Article) error
	GetDetail(ctx context.Context, Key string)(domain.Article,error)
}

func NewArticleRedis(redis redis.Cmdable) ArticleCache {
	return &ArticleRedis{redis: redis}
}

type ArticleRedis struct {
	redis redis.Cmdable
}

func (a *ArticleRedis) GetDetail(ctx context.Context, Key string) (domain.Article, error) {
	bytes, err := a.redis.Get(ctx, Key).Bytes()
	if err != nil {
		return domain.Article{}, err
	}
	var article domain.Article
	err = json.Unmarshal(bytes, &article)
	return article,err
}

func (a *ArticleRedis) SetArtDetail(ctx context.Context, key string, val domain.Article) error {
	bytes, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return a.redis.Set(ctx, key, bytes, time.Minute*3).Err()
}

func (a *ArticleRedis) GetFirstPage(ctx context.Context, key string) ([]domain.Article, error) {
	var articles []domain.Article
	bytes, err := a.redis.Get(ctx, key).Bytes()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, &articles)
	return articles, err
}

func (a *ArticleRedis) SetFirstPage(ctx context.Context, key string, val []domain.Article) error {
	for _, art := range val {
		art.Content = art.Abstract()
	}
	marshal, err := json.Marshal(val)
	if err != nil {
		return err
	}
	return a.redis.Set(ctx, key, marshal, time.Minute*10).Err()

}
