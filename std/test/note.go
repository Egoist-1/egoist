package test



/*
	-cover   	测试覆盖率
	-v			显示每个用例的结果
*/

/*
	"github.com/stretchr/testify/assert"  减少了大量的错误处理
	"github.com/stretchr/testify/require"  require.NoErro require.Error
						testify/suite		
*/


/*
	mock
	mockgen -source=源文件 -destination=目标文件 -package
	

	
*/



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




/*
	t.Helper()  帮助函数 如:多个函数调用该函数 就可以提示具体在那个函数处调用失败
	func createMulTestCase(t *testing.T, c *calcCase) {
	t.Helper()   加入的报错信息 行数: 50 不加:42
	if ans := Mul(c.A, c.B); ans != c.Expected {
	}

}

func TestMul(t *testing.T) {
	createMulTestCase(t, &calcCase{2, 3, 6})
	createMulTestCase(t, &calcCase{2, -3, -6})
	createMulTestCase(t, &calcCase{2, 0, 1}) // wrong case
}
*/

// setup 和 teardown
/*

*/