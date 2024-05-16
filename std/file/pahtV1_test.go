package file

import (
	"fmt"
	"path/filepath"
	"testing"
)

func Test_Path(t *testing.T)  {
	//去除最后一个元素
	// str := filepath.Dir("/es/docker/")
	//返回最后一个元素 如果为  "" = . |  / = / | docker/ =  docker
	str := filepath.Base("docker/")
	fmt.Println(str)
	//val2 相对于 val1的路径
	fmt.Println(filepath.Rel("/home/polaris/studygolang", "/data/studygolang"))
	// ../../../data/studygolang 
	//返回目录
	dir,file:= filepath.Split("/docker/es/ik/ss.yml")
	fmt.Println(dir,file)
}