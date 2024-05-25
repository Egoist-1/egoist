package rpc

import (
	"7day/7day/rpc/codec"
	"encoding/json"
	"io"
	"log"
	"net"
	"reflect"
	"sync"
)

/*
客户端与服务端的通信需要协商一些内容，
例如 HTTP 报文，分为 header 和 body 2 部分，body 的格式和长度通过
 header 中的 Content-Type 和 Content-Length 指定，
 服务端通过解析 header 就能够知道如何从 body 中读取需要的信息。
 对于 RPC 协议来说，这部分协商是需要自主设计的。为了提升性能，
 一般在报文的最开始会规划固定的字节，来协商相关的信息。
 比如第1个字节用来表示序列化方式，第2个字节表示压缩方式，
第3-6字节表示 header 的长度，7-10 字节表示 body 的长度。
*/

const MagicNumber = 0x3bef5c

// 协商的编解码方式
type Option struct {
	MagicNumber int
	CodecType   codec.Type
}

var DefaultOption = &Option{
	MagicNumber: MagicNumber,
	CodecType:   codec.GobType,
}

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

var DefaultServer = NewServer()

// 等待连接
func (server *Server) Accept(lis net.Listener) {
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accpet error", err)
			return
		}
		go server.ServeConn(conn)
	}
}

func (server *Server) ServeConn(conn io.ReadWriteCloser) {
	defer func() { _ = conn.Close() }()
	var opt Option
	if err := json.NewDecoder(conn).Decode(&opt); err != nil {
		log.Println("rpc server:options error", err)
		return
	}
	if opt.MagicNumber != MagicNumber {
		log.Println("rpc server: invald magic number %x", opt, MagicNumber)
		return
	}
	f := codec.NewCodeFuncMap[opt.CodecType]
	if f == nil {
		log.Println("rpc server: invalid codec type %s", opt.CodecType)
		return
	}
	server.serveCodec(f(conn))
}

// 发生错误时,占位符
var invalidRequest = struct{}{}

func (server *Server) serveCodec(cc codec.Codec) {
	sending := new(sync.Mutex)
	wg := new(sync.WaitGroup)
	for {
		req, err := server.readRequest(cc)
		if err != nil {
			if req == nil {
				break
			}
			req.h.Error = err.Error()
			server.sendResponse(cc, req.h, invalidRequest, sending)
			continue
		}
		wg.Add(1)
		go server.handleRequest(cc, req, sending, wg)
	}
	wg.Wait()
	_ = cc.Close()
}

type request struct {
	h            *codec.Header
	argv, replyv reflect.Value
}

func (server *Server) readRequestHeader(cc codec.Codec) (*codec.Header, error) {
	var h codec.Header
	if err := cc.ReadHeader(&h); err != nil {
		if err != io.EOF && err != io.ErrUnexpectedEOF {
			log.Println("rpc server: read header error:", err)
		}
		return nil, err
	}
	return &h, nil
}

func(Server *Server)readrequest(cc codec.Codec)(*request,error){

}
