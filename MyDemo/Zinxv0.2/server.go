package main

import (
	"zinx/znet"
)

func main() {
	server := znet.NewServer("zinxv0.1")
	server.Server()
}
