package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

/*
reflect 的 typeOf 和  typeValue 分别返回 类型 和 值
可以使用 reflect.indirect()
*/

// 动态调用函数
func TestReflect_method(t *testing.T) {
	var r = ref{}
	reflect.ValueOf(r).MethodByName("RefmethodV1").Call(nil)
	a := reflect.ValueOf(1)
	b := reflect.ValueOf("123")
	in := []reflect.Value{a, b}
	reflect.ValueOf(r).MethodByName("RefmethodV2").Call(in)
}

type ref struct{}

func (t ref) RefmethodV1() {
	fmt.Println("动态调用函数")
}
func (t ref) RefmethodV2(a int, b string) {
	fmt.Println("hello"+b, a)
}

// struck tag
func TestReflect_tag(t *testing.T) {
	U := User{
		Name: "张三",
		Age:  22,
	}
	for i := 0; i < reflect.TypeOf(U).NumField(); i++ {
		if tag, ok := reflect.TypeOf(U).Field(i).Tag.Lookup("json"); ok {
			fmt.Println(tag)
		}
		fmt.Println(reflect.TypeOf(U).Field(i).Tag.Get("gorm"))
	}
}

type User struct {
	Name string `json:"name" gorm:"name"`
	Age  int    `json:"age" gorm:"age"`
}

// 类型转换和赋值  通过T标签设置 newT的值
func TestReflect_Kind_trasform(t *testing.T) {
	t1 := T{
		A: 12,
		B: "李四",
	}
	tt := reflect.TypeOf(t1)
	tv := reflect.ValueOf(t1)
	nt := &newT{}
	newtt := reflect.ValueOf(nt)
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		newTtag := field.Tag.Get("newT")
		tValue := tv.Field(i)
		newtt.Elem().FieldByName(newTtag).Set(tValue)
	}
	fmt.Println(nt)
}

type T struct {
	A int    `newT:"AA"`
	B string `newT:"BB"`
}

type newT struct {
	AA int
	BB string
}
