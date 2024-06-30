package article

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
)

func NewArticleDaoImpl(db *gorm.DB) ArticleDao {
	return &ArticleDaoImpl{
		db: db,
	}
}

type ArticleDaoImpl struct {
	db *gorm.DB
}

func (dao *ArticleDaoImpl) GetById(ctx context.Context, id int64) (Article, error) {
	var art Article
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&art).Error
	return art,err
}

func (dao *ArticleDaoImpl) OnePage(ctx context.Context, uid int64, offset int, limit int) ([]Article, error) {
	article := make([]Article, limit)
	err := dao.db.WithContext(ctx).Model(Article{}).Where("author_id = ?", uid).
		Order("ctime desc").
		Offset(offset).Limit(limit).
		Find(&article).Error
	return article, err
}

func (dao *ArticleDaoImpl) SyncStatus(ctx context.Context, art Article) error {
	return dao.db.Transaction(func(tx *gorm.DB) error {
		res := tx.WithContext(ctx).Model(Article{}).
			Where("id = ? and author_id = ?", art.Id, art.AuthorId).
			Update("status", art.Status)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errors.New("500 尝试修改他人数据")
		}
		res = tx.WithContext(ctx).Model(ArticlePublish{}).
			Where("id = ? and author_id = ?", art.Id, art.AuthorId).
			Update("status", art.Status)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected != 1 {
			return errors.New("500 尝试修改他人数据")
		}
		return nil
	})
}

func (dao *ArticleDaoImpl) FindById(ctx context.Context, id int64) (Article, error) {
	var art Article
	err := dao.db.WithContext(ctx).Where("id = ?", id).First(&art).Error
	return art, err
}

func (dao *ArticleDaoImpl) Sync(ctx context.Context, art Article) (int64, error) {
	tx := dao.db.WithContext(ctx).Begin()
	defer tx.Rollback()
	txDao := NewArticleDaoImpl(tx)
	//更新制作库
	proArt, err := txDao.Upset(ctx, art)
	if err != nil {
		return 0, err
	}
	//获取制作库
	part, err := txDao.FindById(ctx, proArt.Id)
	if err != nil {
		return 0, err
	}
	//同步线上库
	artPub := ArticlePublish(part)

	err = tx.Clauses(clause.OnConflict{
		Columns: []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(map[string]interface{}{
			"title":   artPub.Title,
			"content": artPub.Content,
			"status":  artPub.Status,
			"utime":   artPub.Utime,
		}),
	}).Create(&artPub).Error
	if err != nil {
		return 0, err
	}
	tx.Commit()
	return part.Id, tx.Error
}

func (dao *ArticleDaoImpl) Upset(ctx context.Context, article Article) (Article, error) {
	var art Article
	if article.Id > 0 {
		err := dao.db.WithContext(ctx).Where("id=?", article.Id).Find(&art).Error
		if err != nil {
			return Article{}, err
		}
		if art.AuthorId != article.AuthorId {
			return Article{}, errors.New("系统错误")
		}
	}

	now := time.Now().UnixMilli()
	article.Utime = now
	article.Ctime = now
	err := dao.db.WithContext(ctx).Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns([]string{"title", "content", "utime", "status"}),
	}).Create(&article).Error
	return article, err
}
