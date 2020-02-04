package ziface

import "net"

type con uint32

type IConnection interface {
	start()
	stop()
	GetTcpConnect() * net.TCPConn
	GetConnId() uint32
	RemoteAddr() net.Addr
	Send(data []byte)  error
}


type HandelFunc func(*net.TCPConn,[]byte, int)  error