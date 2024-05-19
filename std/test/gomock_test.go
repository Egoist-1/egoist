package test

import (
	test "7day/std/test/mock"
	"errors"
	"testing"

	"go.uber.org/mock/gomock"
)

//mockgen -source=源文件 -destination=目标文件 -package=

func Test_Mock(t *testing.T) {

}

/*
mock 在这里模拟了:GetFromDB内部DB.get返回不同的值来测试不同的结果
ctrl.Finish() 断言 DB.Get()被是否被调用，如果没有被调用，后续的 mock 就失去了意义；


mock.expect.get().return()
参数
Eq(value) 表示与 value 等价的值。
Any() 可以用来表示任意的入参。
Not(value) 用来表示非 value 以外的值。
Nil() 表示 None 值

返回值
Return 返回确定的值
Do Mock 方法被调用时，要执行的操作吗，忽略返回值。
DoAndReturn 可以动态地控制返回值


调用次数
Times() 断言 Mock 方法被调用的次数。
MaxTimes() 最大次数。
MinTimes() 最小次数。
AnyTimes() 任意次数（包括 0 次）。

 调用顺序(InOrder)

*/ 

func Test_DB_test(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() //断言 db.get 方法是否被调用
	m := test.NewMockDB(ctrl)

	m.EXPECT().Get(gomock.Eq("Tom")).Return(100, errors.New("not exist"))

	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1 but got", v)
	}
}
