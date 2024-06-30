package article

import "time"

// 作者库
type Article struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"`
	Title    string `gorm:"type:varchar(4096)"`
	Content  string `gorm:"type:blob"`
	Ctime    int64
	Utime    int64 `gorm:"index:aut_ut,priority:2,BLOB"`
	AuthorId int64 `gorm:"index:aut_ut,priority:9"`
	Status   int
}

// 线上库
type ArticlePublish Article

// 文章 阅读 点赞 收藏 数量
type Interactive struct {
	Id int64
	// art id 和 业务表示
	Biz           string
	BizId         string
	ReadCnt       int64
	LikeCnt       int64
	CollectionCnt int64
	Ctime         time.Time
	Utime         time.Time
}

// User  点赞
type UserLike struct {
	Id    int64
	Biz   string
	BizId int64
	//那位用户
	Uid int64
	// 1-点赞 , 0-未点赞
	Status int8
	Utime  time.Time
	Ctime  time.Time
}

// User 收藏
type UserCollect struct {
	Id    int64
	Biz   string
	BizId int64
	//那位用户
	Uid int64
	//收藏在哪个收藏夹
	Cid   int64
	Utime time.Time
	Ctime time.Time
}

// 收藏夹
type Collection struct {
	Id    int64
	Name  string
	Uid   int64
	Biz   string
	Utime time.Time
	Ctime time.Time
}
