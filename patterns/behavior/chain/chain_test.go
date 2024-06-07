package chain

import (
	"log"
	"testing"
)

type Filter func() error

type FilterChain func(next Filter) Filter

// 组装类
type MyServer struct {
	root Filter
}

// 调用组装的链
func (m *MyServer) Server() error {
	return m.root()
}

// 将链子组装
func NewMyServer(flts ...FilterChain) *MyServer {
	var root Filter = func() error {
		log.Println("这是最后一个 filter")
		return nil
	}
	for i := len(flts) - 1; i >= 0; i-- {
		root = flts[i](root)
	}
	return &MyServer{
		root: root,
	}
}
func TestChain(t *testing.T) {
	var firs FilterChain = func(next Filter) Filter {
		return func() error {
			log.Println("第一个执行前")
			err := next()
			log.Println("第一个执行后")
			return err
		}
	}
	var second FilterChain = func(next Filter) Filter {
		return func() error {
			log.Println("第二个执行前")
			err := next()
			log.Println("第二个执行后")
			return err
		}
	}
	server := NewMyServer(firs, second)
	server.Server()
}
