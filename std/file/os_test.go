package file

import (
	"fmt"
	"os"
	"testing"
)

func Test_File(t *testing.T)  {
	file,err := os.Open("pahtV1_test.go")
	b2 := []byte{'1','2'}
	n,err := file.Write(b2)
	fmt.Println(err)
	fmt.Println(n)
	file.Close()
}