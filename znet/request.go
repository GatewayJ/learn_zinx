package znet

import "zinx/ziface"

type Request struct {
	conn ziface.IConnection
	data []byte
}


func (con Request) GetConnection () ziface.IConnection{
	return con.conn
}

func (con Request) GetData() []byte{
	return con.data
}