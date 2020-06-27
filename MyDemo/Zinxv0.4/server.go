package main

import (
	"fmt"

	"zinx/znet"

	"zinx/ziface"
)

//PingRouter  ping 路由测试
type PingRouter struct {
	znet.BaseRoute
}

// PreHandle 请求之前的钩子
func (br *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle...")
	_, err := request.GetConnection().GetTcpConnect().Write([]byte("before ping ...\n"))
	if err != nil {
		fmt.Println("call back before ping error")
	}
}

// Handle 请求钩子
func (br *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handle...")
	_, err := request.GetConnection().GetTcpConnect().Write([]byte("ping ...ping ...ping ...\n"))
	if err != nil {
		fmt.Println("call back  ping error")
	}
}

// PostHandle 请求之后的钩子
func (br *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle...")
	_, err := request.GetConnection().GetTcpConnect().Write([]byte("post ping ...\n"))
	if err != nil {
		fmt.Println("call back post ping error")
	}
}

// 主函数
func main() {
	server := znet.NewServer("zinxv0.4")
	server.AddRouter(&PingRouter{})
	server.Server()
}
