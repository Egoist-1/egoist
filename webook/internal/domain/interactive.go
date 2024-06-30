package domain

type InterActive struct {
	//一个文章的id
	Biz int64
	//一个业务名称,如 movie article
	BizId      string
	ReadCnt    int64
	LikeCnt    int64
	CollectCnt int64
	//用户是否收藏和点赞
	Collected bool
	Liked     bool
}

// 收藏夹
type Collection struct {
	Name string
	Uid  string
	Resource []Resource
}
type Resource struct {
	Biz   string
	BizId int64
}
