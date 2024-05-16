package io

import (
	"fmt"
	"testing"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}
//定制print
func (p *Person)String()string{
	return "123"
}
// 定制格式化(可以实现自定义输出格式(占位符))
func (p *Person)Format(f fmt.State,c rune)  {
	f.Write([]byte("123531"))
}
func Test_Fmt(t *testing.T) {
	p := &Person{"zhangsan",28,0}
	fmt.Println(p)
	fmt.Printf("%s\n",p)
}
