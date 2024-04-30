package io

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

func TestNewReader(t *testing.T) {
	str := "12345678901234567890"
	reader := strings.NewReader(str)
	bytes := make([]byte, 7)
	for {
		num, err := reader.Read(bytes)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%v--%v\n", num, string(bytes))
	}
}

// 自己实现reader
func TestReader(t *testing.T) {
	str := "q1werasd123zxc2q"
	reader := NewAlphaReader(str)
	b := make([]byte, 10)
	for {
		n, err := reader.Read(b)
		if err == io.EOF {
			fmt.Print(string(b[:n]),"\n")
			break
		}
		fmt.Print(string(b))
	}
}

// 过滤掉非字母的字符
type alphaReader struct {
	str string
	cur int //读取到的位置
}

func NewAlphaReader(str string) *alphaReader {
	return &alphaReader{
		str: str,
	}
}

// 过滤函数
func (a *alphaReader) alpha(b byte) byte {
	if (b >= 'A' && b <= 'Z') || (b >= 'a' && b <= 'z') {
		return b
	}
	return 0
}

// Read 参数p 时缓冲区大小,指定每次读取的大小,
// 返回值 n 返回的字节数,err
func (a *alphaReader) Read(p []byte) (n int, err error) {
	if a.cur > len(a.str) {
		return 0 ,io.EOF
	}

	temp := make([]byte,0,len(p))
	for i:=0;i < len(p);i++{  
		if a.cur >= len(a.str) {
			return i,io.EOF
		}
		b := a.alpha(a.str[a.cur])
		a.cur++
		if b != 0 {
			temp = append(temp, b)
		}
	}	
	copy(p,temp)
	return len(p),nil
}
