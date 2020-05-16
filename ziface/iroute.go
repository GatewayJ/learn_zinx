package ziface

type IRoute interface {
	PreHandle(request IRequest)
	Handle(request IRequest)
	PostHandle(request IRequest)
}
