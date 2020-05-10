package main

import (
	"fmt"
	"zinx/ziface"
	"zinx/znet"
)

//PingRouter  ping 路由测试
type PingRouter struct {
	znet.BaseRoute
}

// PreHandle 请求之前的钩子
func (br *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTcpConnect().Write([]byte("before ping ..."))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

// Handle 请求钩子
func (br *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTcpConnect().Write([]byte(" ping ..."))
	if err != nil {
		fmt.Println("call back  ping error")
	}
}

// PostHandle 请求之后的钩子
func (br *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTcpConnect().Write([]byte("post ping ..."))
	if err != nil {
		fmt.Println("call back post ping error")
	}
}

func main() {
	server := znet.NewServer("zinxv0.3")
	server.AddRouter(&PingRouter{})
	server.Server()
}
