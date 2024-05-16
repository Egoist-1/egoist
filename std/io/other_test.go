package io

import (
	"fmt"
	"os"
	"testing"
)

// os File Write
func Test_Write(t *testing.T)  {
	proverbs := []string{
	    "nihao golang",
		"你好 golang",
		"hello world",
		"qwe123",
	}
	file,err  := os.Create("test.txt")
	if err != nil {
		os.Exit(1)
	}
	for _,v  := range proverbs {
		file.Write([]byte(v))
	}

}

// os File Read
func Test_Read(t *testing.T)  {
	file,err  := os.Open("./test.txt")
	if err != nil {
		os.Exit(1)
	}
	b := make([]byte,1024)
	n,_  := file.Read(b)
	fmt.Println(string(b))
	fmt.Println(n)
}

//标准输入输出
func Test_Std(t *testing.T)  {
	proverbs := []string{
        "Channels orchestrate mutexes serialize\n",
        "Cgo is not Go\n",
        "Errors are values\n",
        "Don't panic\n",
    }
    for _, p := range proverbs {
        // 因为 os.Stdout 也实现了 io.Writer
        n, err := os.Stdout.Write([]byte(p))
        if err != nil {
            fmt.Println(err)
            os.Exit(1)
        }
        if n != len(p) {
            fmt.Println("failed to write data")
            os.Exit(1)
        }
	}
}

// io.Copy()
func Test_Copy(t *testing.T)  {

}

// io.WriteString()
func Test_WriteString(t *testing.T)  {

}