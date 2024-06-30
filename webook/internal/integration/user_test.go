package integration

import (
	"7day/webook/internal/integration/startup"
	"7day/webook/internal/repository/dao/user"
	"7day/webook/internal/web/jwt"
	"7day/webook/ioc"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type UserTest struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
	redis  redis.Cmdable
}

// 所有测试运行前调用这个
func (u *UserTest) SetupSuite() {
	app := startup.InitWebServer()
	u.server = app.Web
	//提前设置好Token
	u.server.Use(func(ctx *gin.Context) {
		ctx.Set("claims", &jwt.UserClaims{
			Id: 123,
		})
		ctx.Next()
	})
	u.db = startup.InitTestDB()
	u.redis = ioc.InitRedis()
}

// 所有测试运行后调用这个
func (u *UserTest) TearDownSuite() {
	err := u.db.Exec("TRUNCATE TABLE `users`").Error
	assert.NoError(u.T(), err)
}

// 套件中每个测试执行前都会调用这个 此处指的是一个Test函数执行前后
func (u *UserTest) SetupTest() {
}

// 套件中每个测试执行后都会调用这个 此处指的是一个Test函数执行前后
func (u *UserTest) TearDownTest() {
}

func (u *UserTest) TestUser_signup() {
	t := u.T()
	testCases := []struct {
		name       string
		before     func(t *testing.T)
		after      func(t *testing.T)
		reqBody    string
		wantCode   int
		wantResult Result[any]
	}{
		{
			name: "注册成功",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {
				err := u.db.Exec("TRUNCATE TABLE `users`").Error
				assert.NoError(u.T(), err)
			},
			reqBody: `
{
	"email": "test@test.com",
	"password": "123",
	"confirm_Password": "123"
}`,
			wantCode: http.StatusOK,
			wantResult: Result[any]{
				Code: 400,
				Msg:  "注册成功",
				Data: nil,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			req, err := http.NewRequest(http.MethodPost, "/users/signup", bytes.NewBuffer([]byte(tc.reqBody)))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			u.server.ServeHTTP(resp, req)
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var webRes Result[any]
			//将响应里的body 解析到webRes里 用于 equal
			err = json.NewDecoder(resp.Body).Decode(&webRes)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantResult, webRes)
			tc.after(t)
		})

	}
}

func (u *UserTest) TestUser_login() {
	t := u.T()
	testCases := []struct {
		name     string
		reqBody  string
		before   func(t *testing.T)
		after    func(t *testing.T)
		wantCode int
		wantBody Result[any]
	}{
		{
			name: "登录成功",
			reqBody: `
{
	"email": "test@test.com",
	"password": "123"
}`,
			before: func(t *testing.T) {
				password, err := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.DefaultCost)
				assert.NoError(t, err)
				email := "test@test.com"

				err = u.db.Model(dao.User{}).Create(&dao.User{
					Email: sql.NullString{
						String: "test@test.com",
						Valid:  email != "",
					},
					Password: string(password),
				}).Error
				assert.NoError(t, err)
			},
			after: func(t *testing.T) {
				err := u.db.Exec("TRUNCATE TABLE `users`").Error
				assert.NoError(u.T(), err)
			},
			wantCode: 200,
			wantBody: Result[any]{
				Code: 2,
				Msg:  "登录成功",
				Data: nil,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			req, err := http.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer([]byte(tc.reqBody)))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			u.server.ServeHTTP(resp, req)
			fmt.Println(resp.Body.String())
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var res Result[any]
			err = json.NewDecoder(resp.Body).Decode(&res)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantBody, res)
			tc.after(t)
		})
	}
}

func (u *UserTest) TestUser_loginSMS() {
	t := u.T()
	testCases := []struct {
		name       string
		before     func(t *testing.T)
		after      func(t *testing.T)
		reqBody    string
		wantCode   int
		wantResult Result[any]
	}{
		{
			name: "发送成功",
			before: func(t *testing.T) {

			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				result, err := u.redis.GetDel(ctx, "user:123").Result()
				u.redis.GetDel(ctx, "user:123:cnt")
				cancel()
				assert.NoError(t, err)
				assert.True(t, len(result) == 6)

			},
			reqBody: `
{
	"phone": "123"
}`,
			wantCode: 200,
			wantResult: Result[any]{
				Code: 2,
				Msg:  "发送成功",
				Data: nil,
			},
		},
		//{
		//	name:       "输入太频繁",
		//	before:     nil,
		//	after:      nil,
		//	reqBody:    "",
		//	wantCode:   0,
		//	wantResult: Result[any]{},
		//},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			req, err := http.NewRequest(http.MethodPost, "/users/login_sms/code/send", bytes.NewBuffer([]byte(tc.reqBody)))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			u.server.ServeHTTP(resp, req)
			fmt.Println(resp.Body.String())
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var res Result[any]
			err = json.NewDecoder(resp.Body).Decode(&res)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantResult, res)
			tc.after(t)
		})

	}
}
func (u *UserTest) TestUser_verify() {
	t := u.T()
	testCases := []struct {
		name       string
		before     func(t *testing.T)
		after      func(t *testing.T)
		reqBody    string
		wantCode   int
		wantResult Result[any]
	}{
		{
			name: "登录成功",
			before: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				_, err := u.redis.Set(ctx, "user:12345678901", "123456", time.Minute*10).Result()
				assert.NoError(t, err)
				_, err = u.redis.Set(ctx, "user:12345678901:cnt", 3, time.Minute*10).Result()
				assert.NoError(t, err)
				cancel()
			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				_, err := u.redis.GetDel(ctx, "user:12345678901").Result()
				assert.NoError(t, err)
				_, err = u.redis.GetDel(ctx, "user:12345678901:cnt").Result()
				assert.NoError(t, err)
				cancel()
			},
			reqBody: `
{
	"phone": "12345678901",
	"code": "123456"
}`,
			wantCode: 200,
			wantResult: Result[any]{
				Code: 2,
				Msg:  "登录成功",
				Data: nil,
			},
		},
		{
			name: "输入错误,重新输入",
			before: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				_, err := u.redis.Set(ctx, "user:12345678901", "123456", time.Minute*10).Result()
				assert.NoError(t, err)
				_, err = u.redis.Set(ctx, "user:12345678901:cnt", 3, time.Minute*10).Result()
				assert.NoError(t, err)
				cancel()
			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				_, err := u.redis.GetDel(ctx, "user:12345678901").Result()
				assert.NoError(t, err)
				_, err = u.redis.GetDel(ctx, "user:12345678901:cnt").Result()
				assert.NoError(t, err)
				cancel()
			},
			reqBody: `
{
	"phone": "12345678901",
	"code": "123457"
}`,
			wantCode: 200,
			wantResult: Result[any]{
				Code: 4,
				Msg:  "输入错误,请重新输入",
				Data: nil,
			},
		},
		{
			name: "频繁输入,请重新获取验证码",
			before: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				_, err := u.redis.Set(ctx, "user:12345678901", "123456", time.Minute*10).Result()
				assert.NoError(t, err)
				_, err = u.redis.Set(ctx, "user:12345678901:cnt", 0, time.Minute*10).Result()
				assert.NoError(t, err)
				cancel()
			},
			after: func(t *testing.T) {
				ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
				_, err := u.redis.GetDel(ctx, "user:12345678901").Result()
				assert.NoError(t, err)
				_, err = u.redis.GetDel(ctx, "user:12345678901:cnt").Result()
				assert.NoError(t, err)
				cancel()
			},
			reqBody: `
{
	"phone": "12345678901",
	"code": "123457"
}`,
			wantCode: 200,
			wantResult: Result[any]{
				Code: 4,
				Msg:  "频繁输入,请重新获取验证码",
				Data: nil,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc.before(t)
			req, err := http.NewRequest(http.MethodPost, "/users/login_sms", bytes.NewBuffer([]byte(tc.reqBody)))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()
			u.server.ServeHTTP(resp, req)
			assert.Equal(t, tc.wantCode, resp.Code)
			if resp.Code != 200 {
				return
			}
			var res Result[any]
			err = json.NewDecoder(resp.Body).Decode(&res)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantResult, res)
			tc.after(t)
		})

	}
}

func TestUser(t *testing.T) {
	suite.Run(t, new(UserTest))
}
