package dao

// CommentSubject 一个内容就有一个主题表
type CommentSubject struct {
	Id         int64
	BizId      int64 //Biz_id 和 biz 组成唯一键
	Biz        string
	MemberId   int64 //作者用户id
	Count      int32 //评论总数 //根评论的楼层 删除评论不会-1
	RootCount  int32 //根评论总数 一级楼层的个数
	AllCount   int32 //评论+回复总数  当前可见的评论
	CreateTime int64
	UpdateTime int64
	//State      int8  //
	//Attrs      int32 //
}

type CommentIndex struct {
	Id         int64
	BizId      int64
	Biz        string
	MemberId   int64 //发表评论用户id
	Root       int64 //根评论id 不为0是回复评论
	Parent     int64 //parent父评论id 为0是root评论
	Floor      int32 //评论楼层 楼层号
	Count      int32 //总数
	RootCount  int32 //根评论总数
	Like       int32 //点赞
	Hate       int32 //点踩
	CreateTime int64
	UpdateTime int64
	//Attrs      int32
	//State      int8
}

// tcpflow
type CommentContent struct {
	CommentId   int64  //主键
	AtMemberIds string //对象id
	Ip          int64  //对象类型
	Platform    int8   //ping tai
	Device      string //根评论id,不为0是回复评论
	Message     string //评论内容
	Meta        string //评论元数据:背景,字体
	CreateTime  int64  //
	UpdateTime  int64
}
