package web

import (
	"time"
)

type artReq struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type reqList struct {
	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}
type ArticleVO struct {
	Id       int64     `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Abstract string    `json:"abstract"`
	Ctime    time.Time `json:"ctime"`
	Utime    time.Time `json:"utime"`
	Status   int       `json:"status"`
	Author   string    `json:"author"`
	//阅读 点赞 收藏
	ReadCnt      int64 `json:"read_cnt"`
	LikeCnt      int64 `json:"like_cnt"`
	CollectedCnt int64 `json:"collected_cnt"`
	//本人是否收藏,点赞
	Collected bool `json:"collected"`
	Liked     bool `json:"liked"`
}
