package codec

import (
	"io"
)

//Header
type Header struct {
	ServiceMethod string //服务名 和 方法名
	Seq           string //请求的序号
	Error         string //错误信息
}

// Codec 对消息体进行编解码的interface,抽象出接口是为了实现不用的Codec
type Codec interface {
	io.Closer
	ReadHeader(*Header)error
	ReadBody(interface{})error
	Write(*Header,interface{})error
}


type NewCodecFunc func (io.ReadWriteCloser) Codec

// Type 编解码类型
type Type string

var (
	 GobType Type = "application/gob"
	 JsonType Type = "application/json"
)

// NewCodecFuncMap 类型与函数的映射
var NewCodecFuncMap map[Type]NewCodecFunc


func init()  {
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}