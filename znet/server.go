package znet

import (
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
		conn ,err :=listener.AcceptTCP()
		if err != nil{
			fmt.Print(err)
		}
		go func() {
			for {
				buf := make([]byte,512)
				cnt, err := conn.Read(buf)
				fmt.Printf("this is %s is\n",buf)
				if err != nil{
					fmt.Printf("read error is %s \n",err)
					return
				}
				if _,err :=conn.Write(buf[:cnt]) ;err != nil{
					fmt.Printf("write error is %s\n",err)
					return
				}

			}
		}()
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