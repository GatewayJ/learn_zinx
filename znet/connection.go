package znet

import (
	"fmt"
	"net"
	"zinx/ziface"
)

type Connection struct {
 	Conn * net.TCPConn
 	ConnId uint32
 	isClosed bool
 	HandApi ziface.HandelFunc
 	ExitChan chan bool
 }

 func NewConnection (conn *net.TCPConn,connID uint32, callbackApi ziface.HandelFunc)  *Connection{
 	c := &Connection{
		Conn:     conn,
		ConnId:   connID,
		isClosed: false,
		HandApi:  callbackApi,
		ExitChan: make(chan bool,1),
	}
	return c
 }



 func (c * Connection) StartReader(){

 	defer fmt.Sprintf("ConnId: %s,RemoteAddr: %s",c.ConnId,c.RemoteAddr().String())
 	defer c.stop()

 	for {
 		buf := make([]byte,512)
 		cnt ,err := c.Conn.Read(buf)
 		if err!= nil {
			fmt.Println("recv buf err", err)
			continue
		}
		if err := c.HandApi(c.Conn,buf,cnt);err != nil{
			fmt.Printf("ConnId: %d,Hander  Err: %s",c.ConnId,err)
			break
		}
	}
 }




 func (c * Connection) start(){
	c.StartReader()
 }

 func (c * Connection)stop(){
		fmt.Println("connection stop")
		if c.isClosed == true{
			return
		}
		c.isClosed = true
		c.Conn.Close()
		close(c.ExitChan)
 }

 func (c * Connection) GetTcpConnect() * net.TCPConn{
		return c.Conn
 }

 func (c * Connection) GetConnId() uint32{
		return c.ConnId
 }

 func (c * Connection) RemoteAddr() net.Addr{
		return c.Conn.RemoteAddr()
 }

 func (c * Connection) Send(data []byte) error{
		return nil
 }









