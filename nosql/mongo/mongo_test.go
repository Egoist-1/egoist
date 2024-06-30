package mongo

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/event"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
	"time"
)

func TestMongo(t *testing.T) {
	// 控制初始化超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	monitor := &event.CommandMonitor{
		//每个命令(查询执行之前)
		Started: func(ctx context.Context, startedEvent *event.CommandStartedEvent) {
			fmt.Println(startedEvent.Command)
		},
		Succeeded: func(ctx context.Context, succeededEvent *event.CommandSucceededEvent) {

		},
		Failed: func(ctx context.Context, failedEvent *event.CommandFailedEvent) {

		},
	}
	opts := options.Client().ApplyURI("mongodb://root:root@localhost:27017").
		SetMonitor(monitor)
	client, err := mongo.Connect(ctx, opts)
	assert.NoError(t, err)
	mdb := client.Database("webook")
	col := mdb.Collection("articles")
	defer func() {
		col.DeleteMany(ctx, bson.D{})
	}()
	for i := range 10 {
		_, err = col.InsertOne(ctx, Article{
			Id:      int64(i),
			Title:   fmt.Sprintf("我的标题%d", i),
			Content: "我的内容",
		})
	}

	var art Article
	// 查询 FindOne
	filter := bson.M{"id": 1}
	err = col.FindOne(ctx, filter).Decode(&art)
	assert.NoError(t, err)
	fmt.Printf("结果1: %v\n", art)

	art = Article{}
	filter2 := bson.M{"id": 456}
	err = col.FindOne(ctx, filter2).Decode(&art)
	if err == mongo.ErrNoDocuments {
		fmt.Println("没有发现文档")
	}
	fmt.Printf("结果1: %v\n", art)

	//更新
	sets := bson.M{"$set": bson.M{"title": "新的标题"}}
	updateRes, err := col.UpdateMany(ctx, filter, sets)
	assert.NoError(t, err)
	fmt.Println("affected", updateRes.ModifiedCount)
	fmt.Println("affected", updateRes.UpsertedCount)
	fmt.Println("affected", updateRes.UpsertedID)
	fmt.Println("affected", updateRes.MatchedCount)

	//or and in查询
	or := bson.A{bson.M{"id": 1}, bson.M{"id": 2}}
	sets = bson.M{"$or": or}
	orRes, err := col.Find(ctx, sets)
	assert.NoError(t, err)
	var ars []Article
	err = orRes.All(ctx, &ars)
	assert.NoError(t, err)
	fmt.Println(ars)

	//in := bson.D{bson.E{"id", bson.M{"$in": []any{123, 456}}}}

	//只返回指定字段
	//inRes, err = col.Find(ctx, in, options.Find().SetProjection(bson.M{
	//	"id":    1,
	//	"title": 1,
	//}))
}

type Article struct {
	Id       int64  `bson:"id,omitempty"`
	Title    string `bson:"title,omitempty"`
	Content  string `bson:"content,omitempty"`
	AuthorId int64  `bson:"author_id,omitempty"`
	Status   uint8  `bson:"status,omitempty"`
	Ctime    int64  `bson:"ctime,omitempty"`
	Utime    int64  `bson:"utime,omitempty"`
}
