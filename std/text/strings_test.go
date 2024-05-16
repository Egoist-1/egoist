package strings

import (
	"fmt"
	"strings"
	"testing"
	"unicode"
) 
func Test_stringsV1(t *testing.T)  {
	fmt.Println(strings.ContainsAny("failu re", "z & i"))//t
	//当 sep 为空时，Count 的返回值是：utf8.RuneCountInString(s) + 1
	fmt.Println(strings.Count("1231", ""))
	//字符串分割, s
	fmt.Printf("Fields are: %q\n", strings.Fields("  foo bar  baz   "))
	fmt.Printf("Fields are: %q\n", strings.FieldsFunc("  foo bar  baz   ",func(r rune) bool {
		if r == 'f' {
			return true
		}
		return  false}))
	//split	sep == "" 表示分割所有
	fmt.Printf("%q\n", strings.Split("foo,bar,baz", ","))
	fmt.Printf("%q\n", strings.SplitAfter("foo,bar,baz", ","))
	//n < 0 == split ,n=0 ==[] ,n 控制返回的元素个数
	fmt.Printf("%q\n", strings.SplitN("foo,bar,baz", ",", -1))
	//index lastindex
	//
	fmt.Println(strings.Repeat("ha",5))
	//替换
	mapping := func(r rune) rune {
		switch {
		case r >= 'A' && r <= 'Z': // 大写字母转小写
			return r + 32
		case r >= 'a' && r <= 'z': // 小写字母不处理
			return r
		case unicode.Is(unicode.Han, r): // 汉字换行
			return '\n'
		}
		return -1 // 过滤所有非字母、汉字的字符
	}
	fmt.Println(strings.Map(mapping, "Hello你#￥%……\n（'World\n,好Hello^(&(*界gopher..."))
	// 如果 n < 0，则不限制替换次数，即全部替换
	//func Replace(s, old, new string, n int) string
	//大小写转换 tolower toupper
	//首字母大写
	//func Title(s string) string
	//全部大写
	//func ToTitle(s string) string
	//func ToTitleSpecial(c unicode.SpecialCase, s string) string
}