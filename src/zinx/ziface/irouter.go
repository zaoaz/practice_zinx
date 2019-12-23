package ziface

//定义路由抽象接口
//路由中的数据使用IRequest
type IRouter interface {
	//处理connection 之前的hook
	PreHandle(request IRequest)
	//处理conn业务主方法hook
	Handle(request IRequest)
	//处理conn业务之后hook
	PostHandle(request IRequest)
}
