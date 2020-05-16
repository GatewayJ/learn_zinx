package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
	Conn     *net.TCPConn
	ConnId   uint32
	isClosed bool
	HandApi  ziface.HandelFunc
	ExitChan chan bool
	Router   ziface.IRoute
}

func NewConnection(conn *net.TCPConn, connID uint32, router ziface.IRoute) *Connection {
	c := &Connection{
		Conn:     conn,
		ConnId:   connID,
		isClosed: false,
		Router:   router,
		ExitChan: make(chan bool, 1),
	}
	return c
}

func (c *Connection) StartReader() {

	defer fmt.Sprintf("ConnId: %s,RemoteAddr: %s", c.ConnId, c.RemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err", err)
			continue
		}
		req := Request{
			conn: c,
			data: buf,
		}
		go func(Request ziface.IRequest) {
			c.Router.PreHandle(Request)
			c.Router.Handle(Request)
			c.Router.PostHandle(Request)
		}(&req)

	}
}

func (c *Connection) Start() {
	c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("connection stop")
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	c.Conn.Close()
	close(c.ExitChan)
}

func (c *Connection) GetTcpConnect() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send(data []byte) error {
	return nil
}
