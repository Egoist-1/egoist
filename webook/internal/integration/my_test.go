package integration

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"
	"testing"
)

type Test struct {
	suite.Suite
	server *gin.Engine
	db     *gorm.DB
	redis  *redis.Client
}

// 所有测试运行前调用这个
func (u *Test) SetupSuite() {
	fmt.Println("所有前")
}

// 所有测试运行后调用这个
func (u *Test) TearDownSuite() {
	fmt.Println("所有后")
}

// 套件中每个测试执行前都会调用这个 此处指的是一个Test函数执行前后
func (u *Test) SetupTest() {
	fmt.Println("前")
}

// 套件中每个测试执行后都会调用这个 此处指的是一个Test函数执行前后
func (u *Test) TearDownTest() {
	fmt.Println("后")
}

func (u *Test) SetupSubTest() {
	fmt.Println("自测试前")
}

func (u *Test) TearDownSubTest() {
	fmt.Println("自测试后")
}

func (u *Test) TestUser_signup() {
	t := u.T()
	testCases := []struct {
		name   string
		before func(t *testing.T)
		after  func(t *testing.T)
	}{
		{
			name: "1",
			before: func(t *testing.T) {
				fmt.Println("before 1")
			},
			after: func(t *testing.T) {
				fmt.Println("after 1")
			},
		},
		{
			name: "2",
			before: func(t *testing.T) {
				fmt.Println("before 2")
			},
			after: func(t *testing.T) {
				fmt.Println("after 2")
			},
		},
	}
	for _, tc := range testCases {
		u.Run(tc.name, func() {
			tc.before(t)
			tc.after(t)
		})

	}
}
func TestExample(t *testing.T) {
	suite.Run(t, new(Test))
}
