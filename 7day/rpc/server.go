package rpc

import (
	"7day/7day/rpc/codec"
	"io"
	"log"
	"net"
)

const MagicNumber = 0x3bef5c

//Option 需要协商的编解码
/*
 在一次连接中,Option固定在报文的最开始,header和body可以有多个
 即:
	| Option | Header1 | Body1 | Header2 | Body2 | ...
 */
type Option struct{
	MagicNumber int
	CodeType codec.Type
}

// Server 代表一个 RPC 服务
type Server struct {}

func NewServer() *Server {
	return &Server{}
}
// DefaultServer 默认server实例,为了用户使用方便
var DefaultServer = NewServer()


// Accept 接受连接和监听服务请求
func (s *Server)Accept(lis net.Listener) {
	net.Listen("tcp",":123")
	//等待socker连接
	for{
		conn,err := lis.Accept()
		if err != nil {
			log.Println("rpc server: accept errpr:",err)
			return
		}
		go s.ServerConn(conn)
		}	
}

func Accept(lis net.Listener){DefaultServer.Accept(lis)}

func (s *Server)ServerConn(conn io.ReadWriteCloser) {
	defer func ()  {_ = conn.Close()}()
	
}