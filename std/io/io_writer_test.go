package io

//Writer 从缓冲区读取资源并写入目标资源
import (
	"bytes"
	"fmt"
	"os"
	"testing"
)


func Test_WriterV1(t *testing.T)  {
	proverbs := []string{"awdawd","你好 golang","hello golang","hello world"}
	var writer bytes.Buffer
	for _,v  := range proverbs {
		n,err  := writer.Write([]byte(v))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(v) {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Print(n,"\t")
	}
	fmt.Println("\n",writer.String())
}

//实现writer
func Test_WriterV2(t *testing.T)  {
	
}