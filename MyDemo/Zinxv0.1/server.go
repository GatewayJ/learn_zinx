package main

import (
	"fmt"
	"zinx/znet"
)

func main(){
	server := znet.NewServer("zinxv0.1")
	fmt.Println(server)
	server.Server()
}
