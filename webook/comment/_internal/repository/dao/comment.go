package dao

type CommentSubject struct {
	Id         int64
	ObjId      int64
	ObjType    int8
	MemberId   int64 //作者用户id
	Count      int32 //评论总数
	RootCount  int32 //根评论总数
	AllCount   int32 //评论+回复总数
	State      int8  //
	Attrs      int32 //
	CreateTime int64
	UpdateTime int64
}
type CommentIndex struct {
	Id         int64
	ObjId      int64
	ObjType    int8
	MemberId   int64 //发白哦这用户id
	Root       int64 //根评论id 不为0是回复评论
	Parent     int64 //parent父评论id 为0是root评论
	Floor      int32 //评论楼层
	Count      int32 //总数
	RootCount  int32 //根评论总数
	Like       int32 //点赞
	Hate       int32 //点踩
	State      int8
	Attrs      int32
	CreateTime int64
	UpdateTime int64
}

type CommentContent struct {
	CommentId   int64  //主键
	AtMemberIds string //对象id
	Ip          int64  //对象类型
	Platform    int8   //发表者用户id
	Device      string //根评论id,不为0是回复评论
	Message     string //评论内容
	Meta        string //评论元数据:背景,字体
	CreateTime  int64  //
	UpdateTime  int64
}
