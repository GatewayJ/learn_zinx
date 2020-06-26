package ziface

//IServer 服务接口
type IServer interface {
	Start()
	Stop()
	Server()
	AddRouter(router IRouter)
}
