package test

import (
	"testing"
)

// 测试覆盖率 -cover

/*
//会中断的
Fail     测试失败,继续执行  继续执行之后的代码
FailNow  测试失败,测试中断

SkipNow  跳过测试,测试中断  不执行后续代码

Fatal	==	log + FailNow
Fatalf	==	logf+ FailNow

//不会中断
Log
Logf

skip	== 	log + skipNow
skipf   ==  logf + skipNow

Error 	== 	log + fail
Erroorf	== 	logf+ fail
*/

// 表驱动测试 测试覆盖率 测试覆盖率 -cover
func Test_Table_Driven(t *testing.T) {
	fibTests := []struct {
		in       int //input
		expected int //expected result
	}{
		{1, 2},
		{2, 2},
		{3, 2},
		{4, 2},
		{5, 2},
		{6, 2},
		{7, 13},
	}
	for _, tt := range fibTests {
		actual := Fib(tt.in)
		if actual != tt.expected {
			t.Errorf("%v 预期: %v 实际 %v", tt.in, actual, tt.expected)
		} else {
			t.Log(tt.in)
		}
	}
}


