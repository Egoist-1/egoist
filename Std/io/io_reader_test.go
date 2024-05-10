package io

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

// io Reader接口
func Test_Reader(t *testing.T) {
	reader := strings.NewReader("hello golang1")
	b := make([]byte, 3)
	for {
		n, _ := reader.Read(b)
		if n < len(b) {
			fmt.Println(string(b[:n]))
			break
		}
		fmt.Println(string(b))
	}
}

// io 实现reader 功能:过滤掉非字母的字符
func Test_ReaderV2(t *testing.T) {
	str := "qwe123asd456q"
	alpha := NewAlphaReader(str)
	b := make([]byte,2)
	for{
		n, err := alpha.Read(b)
		if err != nil {
			fmt.Println(b[:n],n)
			break
		}
		fmt.Println(b,n)
	}
	
}

type alphaReader struct {
	src string //资源
	data []byte //数据
	cur int    //读取到的位置
}
func NewAlphaReader(src string) *alphaReader{
	return &alphaReader{
		src: src,
		data: []byte(src),
	}
}

func (a *alphaReader)alpha(b byte) bool{
    if (b >= 'a' && b <= 'z') ||  (b >= 'A' && b <= 'Z'){
		return true
	}
	return false
}

func (a *alphaReader)Read(b []byte) (n int,err error){
	if a.cur >= len(a.data) {
		return 0,io.EOF
	}
	readN := 0
	for i := 0; i < len(a.data); i++ {
		if a.cur >= len(a.data){
			return readN,io.EOF
		}
		ok := a.alpha(a.data[a.cur])
		if !ok {
			a.cur++
			continue
		}
		b[readN] = a.data[a.cur]
		a.cur++
		readN++
		if readN == len(b){
			return readN,nil
		}
	}
	return readN,nil
}

// 组合多个Reader 目的是重用和屏蔽下层实现的复杂度
func Test_ReaderV3(t *testing.T)  {
	
}