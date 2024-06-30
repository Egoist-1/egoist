package repository

import (
	"7day/webook/internal/domain"
	"7day/webook/internal/repository/cache"
	dao "7day/webook/internal/repository/dao/article"
	"7day/webook/pkg/logger"
	"7day/webook/pkg/tools"
	"context"
	"fmt"

	"time"
)

type ArticleRepo interface {
	Upset(ctx context.Context, article domain.Article) (int64, error)
	Sync(ctx context.Context, art domain.Article) (int64, error)
	SyncStatus(ctx context.Context, article domain.Article) error
	List(ctx context.Context, id int64, offset int, limit int) ([]domain.Article, error)
	GetDetail(ctx context.Context, id int64) (domain.Article, error)
}

func NewArticleRepoImpl(dao dao.ArticleDao, cache cache.ArticleCache, l logger.Logger) ArticleRepo {
	return &ArticleRepoImpl{
		dao:   dao,
		cache: cache,
		l:     l,
	}
}

type ArticleRepoImpl struct {
	dao   dao.ArticleDao
	cache cache.ArticleCache
	l     logger.Logger
}

func (repo *ArticleRepoImpl) GetDetail(ctx context.Context, id int64) (domain.Article, error) {
	detail, err := repo.cache.GetDetail(ctx, repo.artKey(id))
	if err == nil {
		return detail,err
	}
	art, err := repo.dao.GetById(ctx, id)
	if err != nil {
		return domain.Article{},err
	}
	return repo.entityToDomain(art),err
}

func (repo *ArticleRepoImpl) List(ctx context.Context, id int64, offset int, limit int) ([]domain.Article, error) {
	//要在这里使用缓存了
	//	先查缓存
	if offset == 0 && limit <= 20 {
		firstPage, err := repo.cache.GetFirstPage(ctx, repo.firstPageKey(id))
		if err == nil {
			go func() {
				repo.cache.SetArtDetail(ctx, repo.artKey(firstPage[0].Id), firstPage[0])
			}()
			//提高性能再在这里缓存第一篇数据
			return firstPage, err
		}
		repo.l.Info("作者第一页缓存查询失败")
	}
	//	在查数据库
	page, err := repo.dao.OnePage(ctx, id, offset, limit)
	if err != nil || len(page) == 0 {
		return []domain.Article{}, err
	}
	articles := tools.Map[dao.Article, domain.Article](page, func(idx int, s dao.Article) domain.Article {
		return repo.entityToDomain(s)
	})
	//	回写缓存
	if offset == 0 {
		go func() {
			err = repo.cache.SetFirstPage(ctx, repo.firstPageKey(id), articles)
			repo.cache.SetArtDetail(ctx, repo.artKey(articles[0].Id), articles[0])
			if err != nil {
				repo.l.Error("回写创作者缓存失败", logger.String("err:", err.Error()))
			}
		}()

	}
	return articles, nil
}

func (repo *ArticleRepoImpl) SyncStatus(ctx context.Context, article domain.Article) error {
	return repo.dao.SyncStatus(ctx, repo.domainToEntity(article))
}

func (repo *ArticleRepoImpl) Sync(ctx context.Context, art domain.Article) (int64, error) {
	id, err := repo.dao.Sync(ctx, repo.domainToEntity(art))
	return id, err
}

func (repo *ArticleRepoImpl) Upset(ctx context.Context, article domain.Article) (int64, error) {
	art, err := repo.dao.Upset(ctx, repo.domainToEntity(article))
	return art.Id, err
}

func (repo *ArticleRepoImpl) entityToDomain(art dao.Article) domain.Article {
	return domain.Article{
		Id:      art.Id,
		Title:   art.Title,
		Content: art.Content,
		Ctime:   time.UnixMilli(art.Ctime),
		Utime:   time.UnixMilli(art.Utime),
		Status:  art.Status,
		Author: domain.Author{
			Id: art.AuthorId,
		},
	}
}
func (repo *ArticleRepoImpl) domainToEntity(art domain.Article) dao.Article {
	return dao.Article{
		Id:       art.Id,
		Title:    art.Title,
		Content:  art.Content,
		Ctime:    art.Ctime.UnixMilli(),
		Utime:    art.Utime.UnixMilli(),
		AuthorId: art.Author.Id,
		Status:   art.Status,
	}
}

// firstPageKey 创作者第一页的key
func (repo *ArticleRepoImpl) firstPageKey(uid int64) string {
	return fmt.Sprintf("article:firstPage:%v", uid)
}

// artKey 一个文章的Key
func (repo *ArticleRepoImpl) artKey(artId int64) string {
	return fmt.Sprintf("article:detail:%v", artId)
}
// artKey 一个文章的Key
func (repo *ArticleRepoImpl) artPubKey(artId int64) string {
	return fmt.Sprintf("article:PubDetail:%v", artId)
}

