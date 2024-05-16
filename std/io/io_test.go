package io

import (
	"fmt"
	"strings"
	"testing"
)

//io Reader_Writer
func Test_Reader_Writer(t *testing.T)  {
	
}


//io ReaderAt_WriterAt 带有偏移量的io
func Test_ReaderAt_WriterAt(t *testing.T)  {
	//ReaderAy
	reader := strings.NewReader("123456")
	b := make([]byte,3)
	n,_  := reader.ReadAt(b, 2)
	fmt.Println(string(b))
	fmt.Println(n)
	//WriterAt
	
}

//io ReaderFrom_WriterTo
func Test_ReaderFrom_WriterFrom(t *testing.T)  {
	
}

//io Seeker 用于设置偏移量
func Test_Seeker(t *testing.T)  {
	
}


//io Closer 关闭数据流
func Test_Closer(t *testing.T)  {
	
}