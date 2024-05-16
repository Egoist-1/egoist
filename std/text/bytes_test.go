package strings

import (
	"bytes"
	"fmt"
	"testing"
)


func Test_Bytest(t *testing.T)  {
	a := bytes.NewBufferString("123")
	a.Write([]byte("456"))
	fmt.Println(a.String())
}