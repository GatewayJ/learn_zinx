package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/ziface"
)

type Server struct {
	Name string
	IPVersion string
	IP string
	Port int
}

func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	if _,err := conn.Write(data[:cnt]); err != nil{
		return errors.New("call back error")
	}
	return nil
}

func (s * Server) Start(){
	go func() {
		addr ,err :=net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
		if err != nil{
			fmt.Print(err)
		}
		listener ,err:= net.ListenTCP(s.IPVersion,addr)
		if err != nil{
			fmt.Print(err)
		}
		var cid uint32
		cid = 0
		for {
			conn,err := listener.AcceptTCP()
			if err != nil{
				fmt.Printf("Accept err %s",err)
				continue
			}
			dealConn := NewConnection(conn,cid,CallBackToClient)
			cid ++
			go dealConn.start()

		}
	}()

}

func (s * Server) Stop(){

}

func (s * Server) Server(){
	s.Start()
	select {

	}
}


func NewServer(name string) ziface.IServer {
	s := &Server{
		Name:name,
		IPVersion:"tcp4",
		IP:"0.0.0.0",
		Port:8999,
	}
	return s
}