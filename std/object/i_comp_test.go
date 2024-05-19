package object

import "fmt"

//可寻址 就是 可以获取一个变量的地址并且修改
//不可寻址的类型:
//字符串中的字节；
//常量
//map 对象中的元素（slice 对象中的元素是可寻址的，slice的底层是数组）；
//包级别的函数等

type Stu struct {
	Name string
}

type StuInt interface{}

func main() {
	var stu1, stu2 StuInt = &Stu{"Tom"}, &Stu{"Tom"}
	var stu3, stu4 StuInt = Stu{"Tom"}, Stu{"Tom"}
	fmt.Println(stu1 == stu2) // false
	fmt.Println(stu3 == stu4) // true
}

// stu1 和 stu2 对应的类型是 *Stu，值是 Stu 结构体的地址，两个地址不同，因此结果为 false。
// stu3 和 stu4 对应的类型是 Stu，值是 Stu 结构体，且各字段相等，因此结果为 true。