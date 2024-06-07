package codec

import (
	"bufio"
	"encoding/gob"
	"io"
	"log"
)

// var  _ Codec = (*GobCodec)(nil)

type GobCodec struct {
	conn io.ReadWriteCloser //连接实例
	buf  *bufio.Writer      //防止阻塞创建带缓存的writer
	dec  *gob.Decoder       //解码
	enc  *gob.Encoder       //编码
}

func NewGobCodec(conn io.ReadWriteCloser) Codec{
	buf := bufio.NewWriter(conn)
	return &GobCodec{
		conn: conn,
		buf:  buf,
		dec:  gob.NewDecoder(conn),
		enc:  gob.NewEncoder(buf),
	}
}

func(c *GobCodec)ReadHeader(h *Header)error{
	return c.dec.Decode(h)
}

func(c *GobCodec)ReadBody(body interface{})error{
	return c.dec.Decode(body)
}

func (g *GobCodec)Write(h *Header,body interface{}) (err error) {
	defer func ()  {
		_ = g.buf.Flush()
		if err != nil {
			_ = g.Close()
		}
	}()
	if err := g.enc.Encode(h); err != nil {
		log.Println("rpc codec:gob error encoding header:",err)
		return err
	}
	if err := g.enc.Encode(body);err!=nil {
		log.Println("rpc codec:gob error encoding body:",err)
		return err
	}
	return nil
}

func (g *GobCodec) Close() error {
	return g.conn.Close()
}
