package ziface

type IServer interface {
	start()
	stop()
	server()
}