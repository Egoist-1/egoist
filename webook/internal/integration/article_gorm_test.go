package integration

import (
	"7day/webook/internal/domain"
	"7day/webook/internal/integration/startup"
	"7day/webook/internal/repository/dao/article"
	"7day/webook/internal/web"
	"7day/webook/ioc"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type ArticleGormTestSuite struct {
	suite.Suite
	db       *gorm.DB
	server   *gin.Engine
	redis    redis.Cmdable
	artId    int64
	AuthorId int64
}

// 所有测试前
func (u *ArticleGormTestSuite) SetupSuite() {
	app := startup.InitWebServer()
	u.server = app.Web
	u.db = startup.InitTestDB()
	u.redis = ioc.InitRedis()
	u.artId = 10
	u.AuthorId = 123
}

// 自测试后
func (a *ArticleGormTestSuite) TearDownSubTest() {
	err := a.db.Exec("TRUNCATE TABLE `articles`").Error
	assert.NoError(a.T(), err)
	err = a.db.Exec("TRUNCATE TABLE `article_publishes`").Error
	assert.NoError(a.T(), err)
	err = a.redis.FlushAll(context.Background()).Err()
	assert.NoError(a.T(), err)
}

func (a *ArticleGormTestSuite) Test_Article_edit() {
	t := a.T()
	testCases := []struct {
		name     string
		before   func(t *testing.T)
		after    func(t *testing.T)
		reqBody  Article
		wantCode int
		wantBody Result[int64]
	}{
		{
			name: "创建新的贴子",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {

			},
			reqBody: Article{
				Title:   "新的标题",
				Content: "新的贴子",
			},
			wantCode: 200,
			wantBody: Result[int64]{
				Code: 2,
				Msg:  "成功",
				Data: 1,
			},
		},
		{
			name: "更新贴子",
			before: func(t *testing.T) {
				err := a.db.Exec("insert into articles (id,title,content,author_id) values (1,'新的贴子','新的标题',123)").Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {

			},
			reqBody: Article{
				Id:      1,
				Title:   "更新的贴子",
				Content: "更新的标题",
			},
			wantCode: 200,
			wantBody: Result[int64]{
				Code: 2,
				Msg:  "成功",
				Data: 1,
			},
		},
		{
			name: "更新别人的贴子",
			before: func(t *testing.T) {
				err := a.db.Exec("insert into articles (id,title,content,author_id) values (1,'新的贴子','新的标题',777)").Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {

			},
			reqBody: Article{
				Id:      1,
				Title:   "更新的标题",
				Content: "更新的贴子",
			},
			wantCode: 200,
			wantBody: Result[int64]{
				Code: 5,
				Msg:  "系统错误",
				//Data: 0,
			},
		},
	}
	for _, tc := range testCases {
		a.Run(tc.name, func() {
			tc.before(t)
			reqbody, err := json.Marshal(tc.reqBody)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/article/edit", bytes.NewBuffer(reqbody))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			tokenString, err := SetToken()
			assert.NoError(t, err)
			req.Header.Set("Authorization", tokenString)
			resp := httptest.NewRecorder()
			a.server.ServeHTTP(resp, req)
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var art Result[int64]
			err = json.NewDecoder(resp.Body).Decode(&art)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, art)
			tc.after(t)
		})
	}
}

func (a *ArticleGormTestSuite) Test_Article_publish() {

	t := a.T()
	testCases := []struct {
		name     string
		before   func(t *testing.T)
		after    func(t *testing.T)
		reqBody  Article
		wantCode int
		wantBody Result[int64]
	}{
		{
			name: "新建并发表",
			before: func(t *testing.T) {
				var art article.ArticlePublish
				affected := a.db.Where("id = ?", 1).First(&art).RowsAffected
				assert.Equal(t, int64(0), affected)
			},
			after: func(t *testing.T) {
				//prod
				var art article.Article
				err := a.db.Where("id = ?", 1).First(&art).Error
				assert.NoError(t, err)
				assert.Equal(t, "新的标题", art.Title)
				assert.Equal(t, "新的内容", art.Content)
				assert.True(t, art.Utime > 0)
				assert.True(t, art.Ctime > 0)
				assert.True(t, art.Status == 1)
				//pub
				var artPub article.ArticlePublish
				err = a.db.Where("id = ?", 1).First(&artPub).Error
				assert.NoError(t, err)
				//prod pub
				assert.Equal(t, art, article.Article(artPub))
			},
			reqBody: Article{
				Title:   "新的标题",
				Content: "新的内容",
			},
			wantCode: 200,
			wantBody: Result[int64]{
				Code: 2,
				Msg:  "发布成功",
				Data: 1,
			},
		},
		{
			name: "更新草稿 发布线上",
			before: func(t *testing.T) {
				art := article.Article{
					Id:       10,
					Title:    "新的标题",
					Content:  "新的内容",
					Ctime:    555,
					Utime:    444,
					AuthorId: 123,
					Status:   0,
				}
				err := a.db.Create(&art).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//prod
				var art article.Article
				err := a.db.Where("id = ?", 10).First(&art).Error
				assert.NoError(t, err)
				assert.Equal(t, "更新的标题", art.Title)
				assert.Equal(t, "更新的内容", art.Content)
				assert.True(t, art.Utime > 555)
				assert.True(t, art.Ctime > 0)
				assert.True(t, art.Status == 1)
				//pub
				var artPub article.ArticlePublish
				err = a.db.Where("id = ?", 10).First(&artPub).Error
				assert.NoError(t, err)
				//prod pub
				assert.Equal(t, art, article.Article(artPub))
			},
			reqBody: Article{
				Id:      10,
				Title:   "更新的标题",
				Content: "更新的内容",
			},
			wantCode: 200,
			wantBody: Result[int64]{
				Code: 2,
				Msg:  "发布成功",
				Data: 10,
			},
		},
		{
			name: "更新草稿 更新线上",
			before: func(t *testing.T) {
				art := article.Article{
					Id:       10,
					Title:    "新的标题",
					Content:  "新的内容",
					Ctime:    555,
					Utime:    444,
					AuthorId: 123,
					Status:   0,
				}
				err := a.db.Create(&art).Error
				assert.NoError(t, err)
				artPub := article.ArticlePublish{
					Id:       10,
					Title:    "新的标题",
					Content:  "新的内容",
					Ctime:    555,
					Utime:    444,
					AuthorId: 123,
					Status:   0,
				}
				err = a.db.Create(&artPub).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//prod
				var art article.Article
				err := a.db.Where("id = ?", 10).First(&art).Error
				assert.NoError(t, err)
				assert.Equal(t, "更新的标题", art.Title)
				assert.Equal(t, "更新的内容", art.Content)
				assert.True(t, art.Utime > 555)
				assert.True(t, art.Ctime > 0)
				assert.True(t, art.Status == 1)
				//pub
				var artPub article.ArticlePublish
				err = a.db.Where("id = ?", 10).First(&artPub).Error
				assert.NoError(t, err)
				//prod pub
				assert.Equal(t, art, article.Article(artPub))
			},
			reqBody: Article{
				Id:      10,
				Title:   "更新的标题",
				Content: "更新的内容",
			},
			wantCode: 200,
			wantBody: Result[int64]{
				Code: 2,
				Msg:  "发布成功",
				Data: 10,
			},
		},
		{
			name: "更新别人的数据",
			before: func(t *testing.T) {
				art := article.Article{
					Id:       10,
					Title:    "新的标题",
					Content:  "新的内容",
					Ctime:    555,
					Utime:    444,
					AuthorId: 1,
					Status:   0,
				}
				err := a.db.Create(&art).Error
				assert.NoError(t, err)
				artPub := article.ArticlePublish{
					Id:       10,
					Title:    "新的标题",
					Content:  "新的内容",
					Ctime:    555,
					Utime:    444,
					AuthorId: 1,
					Status:   0,
				}
				err = a.db.Create(&artPub).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//prod
				var art article.Article
				err := a.db.Where("id = ?", 10).First(&art).Error
				assert.NoError(t, err)
				assert.Equal(t, "新的标题", art.Title)
				assert.Equal(t, "新的内容", art.Content)
				assert.True(t, art.Utime == 444)
				assert.True(t, art.Ctime == 555)
				assert.True(t, art.Status == 0)
				//pub
				var artPub article.ArticlePublish
				err = a.db.Where("id = ?", 10).First(&artPub).Error
				assert.NoError(t, err)
				//prod pub
				assert.Equal(t, art, article.Article(artPub))
			},
			reqBody: Article{
				Id:      10,
				Title:   "更新的标题",
				Content: "更新的内容",
			},
			wantCode: 200,
			wantBody: Result[int64]{
				Code: 5,
				Msg:  "系统错误",
			},
		},
	}
	for _, tc := range testCases {
		a.Run(tc.name, func() {
			tc.before(t)
			reqBody, err := json.Marshal(tc.reqBody)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/article/publish", bytes.NewBuffer([]byte(reqBody)))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			tokenString, err := SetToken()
			assert.NoError(t, err)
			req.Header.Set("Authorization", tokenString)
			//
			resp := httptest.NewRecorder()
			a.server.ServeHTTP(resp, req)
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var art Result[int64]
			err = json.NewDecoder(resp.Body).Decode(&art)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, art)
			tc.after(t)
		})
	}
}

func (a *ArticleGormTestSuite) Test_Article_withdraw() {
	t := a.T()
	testCases := []struct {
		name     string
		before   func(t *testing.T)
		after    func(t *testing.T)
		reqBody  Article
		wantCode int
		wantBody Result[int64]
	}{
		{
			name: "修改状态",
			before: func(t *testing.T) {
				art := article.Article{
					Id:       a.artId,
					Title:    "新的标题",
					Content:  "新的内容",
					Ctime:    555,
					Utime:    444,
					AuthorId: 123,
					Status:   1,
				}
				err := a.db.Create(&art).Error
				assert.NoError(t, err)
				artPub := article.ArticlePublish{
					Id:       a.artId,
					Title:    "新的标题",
					Content:  "新的内容",
					Ctime:    555,
					Utime:    444,
					AuthorId: 123,
					Status:   1,
				}
				err = a.db.Create(&artPub).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//prod
				var art article.Article
				err := a.db.Where("id = ?", a.artId).First(&art).Error
				assert.NoError(t, err)
				assert.Equal(t, "新的标题", art.Title)
				assert.Equal(t, "新的内容", art.Content)
				assert.True(t, art.Status == 0)
				//pub
				var artPub article.ArticlePublish
				err = a.db.Where("id = ?", a.artId).First(&artPub).Error
				assert.NoError(t, err)
				//prod pub
				assert.Equal(t, art, article.Article(artPub))
			},
			reqBody: Article{
				Id: a.artId,
			},
			wantCode: 200,
			wantBody: Result[int64]{
				Code: 2,
				Msg:  "修改成功",
			},
		},
		{
			name: "修改别人的数据",
			before: func(t *testing.T) {
				art := article.Article{
					Id:       a.artId,
					Title:    "新的标题",
					Content:  "新的内容",
					Ctime:    555,
					Utime:    444,
					AuthorId: 1,
					Status:   1,
				}
				err := a.db.Create(&art).Error
				assert.NoError(t, err)
				artPub := article.ArticlePublish{
					Id:       a.artId,
					Title:    "新的标题",
					Content:  "新的内容",
					Ctime:    555,
					Utime:    444,
					AuthorId: 1,
					Status:   1,
				}
				err = a.db.Create(&artPub).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				//prod
				var art article.Article
				err := a.db.Where("id = ?", a.artId).First(&art).Error
				assert.NoError(t, err)
				assert.Equal(t, "新的标题", art.Title)
				assert.Equal(t, "新的内容", art.Content)
				assert.True(t, art.Status == 1)
				//pub
				var artPub article.ArticlePublish
				err = a.db.Where("id = ?", a.artId).First(&artPub).Error
				assert.NoError(t, err)
				//prod pub
				assert.Equal(t, art, article.Article(artPub))
			},
			reqBody: Article{
				Id: a.artId,
			},
			wantCode: 200,
			wantBody: Result[int64]{
				Code: 5,
				Msg:  "系统错误",
			},
		},
	}
	for _, tc := range testCases {
		a.Run(tc.name, func() {
			tc.before(t)
			reqBody, err := json.Marshal(tc.reqBody)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/article/withdraw", bytes.NewBuffer([]byte(reqBody)))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			tokenString, err := SetToken()
			assert.NoError(t, err)
			req.Header.Set("Authorization", tokenString)
			//
			resp := httptest.NewRecorder()
			a.server.ServeHTTP(resp, req)
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var art Result[int64]
			err = json.NewDecoder(resp.Body).Decode(&art)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, art)
			tc.after(t)
		})
	}
}

func (a *ArticleGormTestSuite) Test_Article_list() {
	t := a.T()
	var testCasest = []struct {
		name     string
		before   func(t *testing.T)
		after    func(t *testing.T)
		reqBody  listReq
		wantCode int
		wantBody Result[[]web.ArticleVO]
	}{
		{
			name: "获取一页文章,没走缓存",
			//插入id 1-10 的文章
			before: func(t *testing.T) {
				var arts []article.Article
				for i := 1; i <= 10; i++ {
					arts = append(arts, article.Article{
						Id:       int64(i),
						Title:    fmt.Sprintf("标题%v", i),
						Content:  fmt.Sprintf("内容%v", i),
						Ctime:    555 + int64(i),
						Utime:    444 + int64(i),
						AuthorId: a.AuthorId,
						Status:   1,
					})
				}
				err := a.db.Create(&arts).Error
				assert.NoError(t, err)
				//-------
				var arts2 []article.Article
				for i := 55; i <= 65; i++ {
					arts2 = append(arts2, article.Article{
						Id:       int64(i),
						Title:    fmt.Sprintf("标题%v", i),
						Content:  fmt.Sprintf("内容%v", i),
						Ctime:    555 + int64(i),
						Utime:    444 + int64(i),
						AuthorId: 234,
						Status:   1,
					})
				}
				err = a.db.Create(&arts2).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				bytes, err := a.redis.Get(context.TODO(), "article:firstPage:123").Bytes()
				assert.NoError(t, err)
				var redisArt []domain.Article
				err = json.Unmarshal(bytes, &redisArt)
				assert.NoError(t, err)
				article := make([]article.Article, 3)
				err = a.db.WithContext(context.Background()).Model(Article{}).Where("author_id = ?", 123).
					Order("ctime desc").
					Offset(0).Limit(3).
					Find(&article).Error
				assert.NoError(t, err)
				for i := range article {
					assert.Equal(t, article[i].Id, redisArt[i].Id)
				}

			},
			reqBody: listReq{
				Offset: 0,
				Limit:  3,
			},
			wantCode: 200,
			wantBody: Result[[]web.ArticleVO]{
				Code: 2,
				Msg:  "ok",
				Data: []web.ArticleVO{web.ArticleVO{
					Id:       10,
					Title:    "标题10",
					Abstract: "内容10",
					Ctime:    time.UnixMilli(555 + 10),
					Utime:    time.UnixMilli(444 + 10),
					Status:   1,
				}, web.ArticleVO{
					Id:       9,
					Title:    "标题9",
					Abstract: "内容9",
					Ctime:    time.UnixMilli(555 + 9),
					Utime:    time.UnixMilli(444 + 9),
					Status:   1,
				}, web.ArticleVO{
					Id:       8,
					Title:    "标题8",
					Abstract: "内容8",
					Ctime:    time.UnixMilli(555 + 8),
					Utime:    time.UnixMilli(444 + 8),
					Status:   1,
				}},
			},
		},
		{
			name: "获取一页文章,走缓存",
			//插入id 1-10 的文章
			before: func(t *testing.T) {
				var arts []domain.Article
				for i := 10; i >= 8; i-- {
					arts = append(arts, domain.Article{
						Id:      int64(i),
						Title:   fmt.Sprintf("标题%v", i),
						Content: fmt.Sprintf("内容%v", i),
						Ctime:   time.UnixMilli(555 + int64(i)),
						Utime:   time.UnixMilli(444 + int64(i)),
						Status:  1,
						Author:  domain.Author{Id: a.AuthorId},
					})
				}
				marshal, err := json.Marshal(arts)
				assert.NoError(t, err)
				err = a.redis.Set(context.TODO(), "article:firstPage:123", marshal, 0).Err()
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
			},
			reqBody: listReq{
				Offset: 0,
				Limit:  3,
			},
			wantCode: 200,
			wantBody: Result[[]web.ArticleVO]{
				Code: 2,
				Msg:  "ok",
				Data: []web.ArticleVO{web.ArticleVO{
					Id:       10,
					Title:    "标题10",
					Abstract: "内容10",
					Ctime:    time.UnixMilli(555 + 10),
					Utime:    time.UnixMilli(444 + 10),
					Status:   1,
				}, web.ArticleVO{
					Id:       9,
					Title:    "标题9",
					Abstract: "内容9",
					Ctime:    time.UnixMilli(555 + 9),
					Utime:    time.UnixMilli(444 + 9),
					Status:   1,
				}, web.ArticleVO{
					Id:       8,
					Title:    "标题8",
					Abstract: "内容8",
					Ctime:    time.UnixMilli(555 + 8),
					Utime:    time.UnixMilli(444 + 8),
					Status:   1,
				}},
			},
		},
		{
			name: "获取最后一页",
			//插入id 1-10 的文章
			before: func(t *testing.T) {
				var arts []article.Article
				for i := 1; i <= 10; i++ {
					arts = append(arts, article.Article{
						Id:       int64(i),
						Title:    fmt.Sprintf("标题%v", i),
						Content:  fmt.Sprintf("内容%v", i),
						Ctime:    555 + int64(i),
						Utime:    444 + int64(i),
						AuthorId: a.AuthorId,
						Status:   1,
					})
				}
				err := a.db.Create(&arts).Error
				assert.NoError(t, err)
				//-------
				var arts2 []article.Article
				for i := 55; i <= 65; i++ {
					arts2 = append(arts2, article.Article{
						Id:       int64(i),
						Title:    fmt.Sprintf("标题%v", i),
						Content:  fmt.Sprintf("内容%v", i),
						Ctime:    555 + int64(i),
						Utime:    444 + int64(i),
						AuthorId: 234,
						Status:   1,
					})
				}
				err = a.db.Create(&arts2).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {

			},
			reqBody: listReq{
				Offset: 7,
				Limit:  3,
			},
			wantCode: 200,
			wantBody: Result[[]web.ArticleVO]{
				Code: 2,
				Msg:  "ok",
				Data: []web.ArticleVO{web.ArticleVO{
					Id:       3,
					Title:    "标题3",
					Abstract: "内容3",
					Ctime:    time.UnixMilli(555 + 3),
					Utime:    time.UnixMilli(444 + 3),
					Status:   1,
				}, web.ArticleVO{
					Id:       2,
					Title:    "标题2",
					Abstract: "内容2",
					Ctime:    time.UnixMilli(555 + 2),
					Utime:    time.UnixMilli(444 + 2),
					Status:   1,
				}, web.ArticleVO{
					Id:       1,
					Title:    "标题1",
					Abstract: "内容1",
					Ctime:    time.UnixMilli(555 + 1),
					Utime:    time.UnixMilli(444 + 1),
					Status:   1,
				}},
			},
		},
		{
			name: "作者没有文章",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {

			},
			reqBody: listReq{
				Offset: 0,
				Limit:  2,
			},
			wantCode: 200,
			wantBody: Result[[]web.ArticleVO]{
				Code: 2,
				Msg:  "ok",
				Data: []web.ArticleVO{},
			},
		},
	}
	for _, tc := range testCasest {
		a.Run(tc.name, func() {
			tc.before(t)
			reqBody, err := json.Marshal(tc.reqBody)
			assert.NoError(t, err)
			req, err := http.NewRequest(http.MethodPost, "/article/list", bytes.NewBuffer((reqBody)))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			tokenString, err := SetToken()
			assert.NoError(t, err)
			req.Header.Set("Authorization", tokenString)
			resp := httptest.NewRecorder()
			a.server.ServeHTTP(resp, req)
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var resBody Result[[]web.ArticleVO]
			err = json.NewDecoder(resp.Body).Decode(&resBody)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, resBody)
			tc.after(t)
		})

	}
}

func TestArticleGorm(t *testing.T) {
	suite.Run(t, new(ArticleGormTestSuite))
}

type listReq struct {
	Offset int
	Limit  int
}

type Article struct {
	Id      int64
	Title   string
	Content string
}
