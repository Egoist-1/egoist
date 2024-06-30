package dao

import (
	"7day/webook/internal/repository/dao/article"
	dao "7day/webook/internal/repository/dao/user"
	"gorm.io/gorm"
)

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(
		dao.User{},
		article.Article{},
		article.ArticlePublish{},
		article.Interactive{},
		article.UserLike{},
		article.UserCollect{},
		article.Collection{},
	)
}
