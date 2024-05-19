package sync

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"bytes"
)

//一句话总结：保存和复用临时对象，减少内存分配，降低 GC 压力

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Geektutu", Age: 25})


//只需要实现new函数,对象池中没有对象时,将会调用new函数创建
var studentPool = sync.Pool{
	New: func() any {
		return new(Student)
	},
}

func TestPool(t *testing.T){

	//get & put
	stu,ok  := studentPool.Get().(*Student)
	if !ok {
		t.Fatal("stu get error")
	}
	fmt.Println(stu.Name)
	json.Unmarshal(buf,stu)
	studentPool.Put(stu)//返回对象池
	fmt.Println(stu.Name)
}


func BenchmarkUnmarshal(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		stu := &Student{}
		json.Unmarshal(buf,stu)
	}
}

func BenchmarkUnmarshalWithPool(b *testing.B)  {
	for i := 0; i < b.N; i++ {
		stu := studentPool.Get().(*Student)
		json.Unmarshal(buf,stu)
		studentPool.Put(stu)
	}
}

//bytes.Buffer
var bufferPool = sync.Pool{
	New: func() interface{} {
		return &bytes.Buffer{}
	},
}

var data = make([]byte, 10000)

func BenchmarkBufferWithPool(b *testing.B) {
	for n := 0; n < b.N; n++ {
		buf := bufferPool.Get().(*bytes.Buffer)
		buf.Write(data)
		buf.Reset()
		bufferPool.Put(buf)
	}
}

func BenchmarkBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		var buf bytes.Buffer
		buf.Write(data)
	}
}