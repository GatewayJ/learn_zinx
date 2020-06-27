package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"zinx/ziface"
)

// GlobalObj 全局参数
type GlobalObj struct {
	// server端配置
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	Version        string
	MaxConn        int
	MaxPackageSize uint32
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	dir, _ := os.Getwd()
	fmt.Println("当前路径：", dir)
	data, err := ioutil.ReadFile("conf\\zinx.json")
	if err != nil {
		print(err)
		panic(err)
	}
	err = json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}
}

func init() {
	GlobalObject = &GlobalObj{
		Name:           "ZinxServer",
		Version:        "v0.4",
		Host:           "0.0.0.0",
		MaxConn:        1000,
		MaxPackageSize: 4096,
	}
	GlobalObject.Reload()
}
