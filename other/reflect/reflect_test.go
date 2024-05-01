package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

// 动态调用函数
func TestReflect_method(t *testing.T) {
	var r = ref{}
	reflect.ValueOf(r).MethodByName("Refmethod")
}
type ref struct{}
func (t ref) Refmethod() {
	fmt.Println("动态调用函数")
}



// struck tag
func TestReflect_tag(t *testing.T) {
	U := User{
		Name: "张三",
		Age: 22,
	}
	for i := 0; i < reflect.TypeOf(U).NumField(); i ++{
		if tag,ok := reflect.TypeOf(U).Field(i).Tag.Lookup("json"); ok {
			fmt.Println(tag)
		}
		fmt.Println(reflect.TypeOf(U).Field(i).Tag.Get("test"))
	}
}
type User struct {
	Name string `json:"name" test:"name"`
	Age  int    `json:"age" test:"age"`
}





// 类型转换和赋值  通过T标签设置 newT的值
func TestReflect_Kind_trasform(t *testing.T)  {
	 t1 := T{
		A: 12,
		B: "李四",
	}
	tt := reflect.TypeOf(t1)
	tv := reflect.ValueOf(t1)
	nt := &newT{} 
	newtt :=  reflect.ValueOf(nt)
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		newTtag := field.Tag.Get("newT")
		tValue := tv.Field(i)
		newtt.Elem().FieldByName(newTtag).Set(tValue)
	}
	fmt.Println(nt)
}

type T struct {
	A int  `newT:"AA"`
	B string `newT:"BB"`
}

type newT struct {
	AA int
	BB string
}
