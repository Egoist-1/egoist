package geerpc

import (
	"7day/7day/geerpc/codec"
	"sync"
)

// Call 承载一次RPC调用所需的信息
type Call struct {
	Seq           uint64 //序列号
	ServiceMethod string //format "<service>.<method>"
	Args          interface{}
	Reply         interface{}
	Error         error
	Done          chan *Call //呼叫完成后
}

// done 调用结束时,会调用call.done通知调用方
func (call *Call) done() {
	call.Done <- call
}

type Client struct {
	cc      codec.Codec
	opt     *Option
	sending sync.Mutex //保证请求的有序发送,即防止出现多个请求报文混淆
	header  codec.Header
	mu      sync.Mutex
	seq     uint64 //seq 用于给发送的请求编号，每个请求拥有唯一编号。
	pending map[uint64]*Call
	// closing 和 shutdown 任意一个值置为 true，则表示 Client 处于不可用的状态，但有些许的差别，
	//closing 是用户主动关闭的，即调用 Close 方法，而 shutdown 置为 true 一般是有错误发生。
	closing  bool
	shutdown bool
}
